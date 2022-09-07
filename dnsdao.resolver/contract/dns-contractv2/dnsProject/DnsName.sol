//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../ERC721/ERC721.sol";
import "../utils/Owner.sol";
import "./lib/LibDnsToolKit.sol";
import "./tax/IDnsTax.sol";
import "./price/IDnsPrice.sol";
import "./accountant/IDnsAccountant.sol";
import "../ERC20/IERC20.sol";
import "./lib/LibDnsName.sol";
import "./lib/LibDnsOwner.sol";
import "./lib/LibDnsSig.sol";
import "./owners/IDnsDaoOwner.sol";
import "./owners/IDnsErc721Owner.sol";

//import "./INamePub.sol";

contract DnsName is ERC721,owned{
    using Strings for uint256;
    uint256 public gNameId;
    address public taxC;
    address public accountantC;
    bytes private baseUri;
    address public ownerC;
    address public priceC;
    address public sigUser;
    mapping(uint32=>bool) public passCardUsed;

    struct NameItem{
        bytes entireName;
        bool openToReg;
        address erc721Addr;
        uint256 expireTime;
        uint256 tokenId;
    }
    mapping(bytes32=>NameItem) public nameStore;
    mapping(uint256=>bytes32) public id2Hash;

    event EvMintDnsName(address erc20Addr, string entireName,uint8 year);
    event EvChargeDnsName(address erc20Addr, bytes32 nameHash, uint8 year, bool isTransferOwner);
    event EvOpenToReg(bool open,bytes32 hash);
    event EvNewSubName(bytes32 hash, uint256 tokenId, address erc721Addr);

    constructor(string memory name_,
        string memory symbol_
    ) ERC721(name_,symbol_){
        gNameId = 1;
        ownerC = LibDnsOwner.NewDnsOwner(msg.sender);
    }


    function setBaseUri(string memory baseUri_) external onlyOwner{
        baseUri = bytes(baseUri_);
    }

    function transferOwnership(address payable newOwner_) external override onlyOwner{
        owner = newOwner_;
        IDnsDaoOwner(ownerC).UpdateOwner(bytes32(0),newOwner_);
    }

    function setRegistered(bool open_,bytes32 hash_) external {
        require(nameStore[hash_].tokenId>0,"name was not registered");
        require(super.ownerOf(nameStore[hash_].tokenId) == msg.sender,"not a valid user");
        require(nameStore[hash_].openToReg != open_,"nothing to do..");
        if(open_ && nameStore[hash_].erc721Addr == address(0)){
            nameStore[hash_].erc721Addr = LibDnsName.NewDnsSubName(
                hash_,
                taxC,
                priceC,
                accountantC,
                msg.sender,ownerC);
            emit EvNewSubName(hash_,nameStore[hash_].tokenId,nameStore[hash_].erc721Addr);
            IDnsDaoOwner(ownerC).addAllowance(address(nameStore[hash_].erc721Addr));
            IDnsDaoOwner(ownerC).UpdateOwner(hash_,msg.sender,nameStore[hash_].erc721Addr);
        }
        nameStore[hash_].openToReg = open_;
        emit EvOpenToReg(open_,hash_);
    }

    function setContract(address taxC_,address accountantC_, address priceC_,address sigUser_,bool open_) external onlyOwner{
        taxC = taxC_;
        accountantC = accountantC_;
        priceC = priceC_;
        sigUser = sigUser_;
        IDnsDaoOwner(ownerC).openLevel2Reg(open_);
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool){
        return super.supportsInterface(interfaceId);
    }

    function _baseURI() internal view override returns (string memory){
        return string(baseUri);
    }

    function tokenURI(uint256 tokenId) public view override returns (string memory){
        super._requireMinted(tokenId);
        string memory baseURI = _baseURI();
        return bytes(baseURI).length > 0 ? string(abi.encodePacked(baseURI, tokenId.toString())) : "";
    }

    function approve(address to, uint256 tokenId) public  override{
        require(nameStore[id2Hash[tokenId]].expireTime>block.timestamp,"not a valid token");
        return super.approve(to,tokenId);
    }

    function getApproved(uint256 tokenId) public view override returns (address){
        require(nameStore[id2Hash[tokenId]].expireTime>block.timestamp,"not a valid token");
        return super.getApproved(tokenId);
    }

    function setApprovalForAll(address operator_, bool approved) public override {
        return super.setApprovalForAll(operator_,approved);
    }

    function isApprovedForAll(address owner, address operator_) public view override returns (bool){
        return super.isApprovedForAll(owner,operator_);
    }

    function _dnsTransfer(address from, address to, uint256 tokenId) internal {
        require(nameStore[id2Hash[tokenId]].expireTime>block.timestamp,"not a valid token");
        require(from != to,"same address error");
        IDnsDaoOwner(ownerC).UpdateOwner(id2Hash[tokenId],to);
        if(nameStore[id2Hash[tokenId]].erc721Addr != address(0)){
            IDnsErc721Owner(nameStore[id2Hash[tokenId]].erc721Addr).transferOwnership(payable(to));
        }
    }

    function transferFrom(address from, address to, uint256 tokenId) public override{
        _dnsTransfer(from,to,tokenId);
        return super.transferFrom(from,to,tokenId);
    }

    function safeTransferFrom(address from, address to, uint256 tokenId) public override {
        safeTransferFrom(from, to, tokenId, "");
    }

    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory data) public override {
        _dnsTransfer(from,to,tokenId);
        super.safeTransferFrom(from,to,tokenId,data);
    }

    function _extendExpire(bytes32 nameHash_,uint8 year_,bool transfer_) internal {
        if (nameStore[nameHash_].expireTime > block.timestamp){
            nameStore[nameHash_].expireTime += year_ * 365 days;
        }else{
            nameStore[nameHash_].expireTime = block.timestamp + year_ * 365 days;
            if(ownerOf(nameStore[nameHash_].tokenId) != msg.sender && transfer_){
                _transfer(ownerOf(nameStore[nameHash_].tokenId),msg.sender,nameStore[nameHash_].tokenId);
                if(nameStore[nameHash_].erc721Addr != address(0)){
                    IDnsErc721Owner(nameStore[nameHash_].erc721Addr).transferOwnership(payable(msg.sender));
                }
                IDnsDaoOwner(ownerC).UpdateOwner(nameHash_,msg.sender);
            }
        }
    }

    function _chargeTopName(address erc20Addr_, bytes32 nameHash_,uint8 year_) internal{
        uint256 cost = year_*IDnsPrice(priceC).Price(bytes32(0),erc20Addr_,nameStore[nameHash_].entireName.length);
        require(cost>0,"payment not support");
        if (erc20Addr_ == address(0)){
            require( msg.value>=cost,"payout is not enough");
            IDnsAccountant(accountantC).deposit(erc20Addr_,msg.value);
            payable(accountantC).transfer(msg.value);
        }else{
            require(IERC20(erc20Addr_).balanceOf(msg.sender)>=cost &&
                IERC20(erc20Addr_).allowance(msg.sender,address(this))>= cost,"payout is not enough");
            IDnsAccountant(accountantC).deposit(erc20Addr_,cost);
            IERC20(erc20Addr_).transferFrom(msg.sender,address(accountantC),cost);
        }

    }

    function ChargeName(address erc20Addr_, bytes32 nameHash_,uint8 year_, bool transfer_) external payable{
        require(nameStore[nameHash_].tokenId > 0,"name was registered");
        _chargeTopName(erc20Addr_,nameHash_,year_);
        _extendExpire(nameHash_,year_,transfer_);
        emit EvChargeDnsName(erc20Addr_,nameHash_,year_,transfer_);
    }

    function _mintName(bytes memory entireName_,uint8 year_) internal {
        id2Hash[gNameId] = LibDnsToolKit.entireNameHash(string(entireName_));
        nameStore[id2Hash[gNameId]].expireTime = block.timestamp + year_ * 365 days;
        nameStore[id2Hash[gNameId]].entireName = entireName_;
        nameStore[id2Hash[gNameId]].tokenId =gNameId;
        _mint(msg.sender,gNameId);
        IDnsDaoOwner(ownerC).UpdateOwner(id2Hash[gNameId],msg.sender,gNameId,entireName_);
        gNameId++;
    }

    // function _mintTopName(address erc20Addr_,string memory entireName_,uint8 year_) internal{
    //     require(LibDnsToolKit.verifyRoot(bytes(entireName_)),"not a root Name");
    //     require(nameStore[LibDnsToolKit.entireNameHash(entireName_)].tokenId == 0,"name was registered");
    //     uint256 cost = uint256(year_)*IDnsPrice(priceC).Price(bytes32(0),erc20Addr_,bytes(entireName_).length);
    //     require(cost>0,"payment not support");
    //     if (erc20Addr_ == address(0)){
    //         require(msg.value>=cost,"payout is not enough");
    //         IDnsAccountant(accountantC).deposit(erc20Addr_,msg.value);
    //         payable(accountantC).transfer(msg.value);

    //     }else{
    //         require(IERC20(erc20Addr_).balanceOf(msg.sender)>=cost &&
    //             IERC20(erc20Addr_).allowance(msg.sender,address(this))>= cost,"payout is not enough");
    //         IDnsAccountant(accountantC).deposit(erc20Addr_,cost);
    //         IERC20(erc20Addr_).transferFrom(msg.sender,address(accountantC),cost);
    //     }
    // }

    function MintNameBySig(string memory entireName_,
        uint8 year_, address erc20Addr_,
        uint256 price_, uint32 passId_,bytes memory sig) external payable{
        require(nameStore[LibDnsToolKit.entireNameHash(entireName_)].tokenId == 0,"name was registered");
        // require(timeStamp_+10 > block.timestamp,"sig is expired");
        //require sig is right
        require(!passCardUsed[passId_],"pass card is used");
        require(LibDnsSignature.SigUserAddr(
            keccak256(abi.encodePacked(entireName_,year_,erc20Addr_,price_,msg.sender,passId_)),
            sig) == sigUser,"sig not correct");

        if (erc20Addr_ == address(0)){
            require(msg.value>=price_,"payout is not enough");
            IDnsAccountant(accountantC).deposit(erc20Addr_,msg.value);
            payable(accountantC).transfer(msg.value);

        }else{
            require(IERC20(erc20Addr_).balanceOf(msg.sender)>=price_ &&
                IERC20(erc20Addr_).allowance(msg.sender,address(this))>= price_,"payout is not enough");
            IDnsAccountant(accountantC).deposit(erc20Addr_,price_);
            IERC20(erc20Addr_).transferFrom(msg.sender,address(accountantC),price_);
        }
        _mintName(bytes(entireName_),year_);
        emit EvMintDnsName(erc20Addr_,entireName_,year_);
    }

    // function MintName(address erc20Addr_, string memory entireName_,uint8 year_) external payable{
    //     _mintTopName(erc20Addr_,entireName_,year_);
    //     _mintName(bytes(entireName_),year_);
    //     emit EvMintDnsName(erc20Addr_,entireName_,year_);
    // }
}