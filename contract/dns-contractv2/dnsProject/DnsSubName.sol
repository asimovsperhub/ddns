//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./owners/DnsErc721Owner.sol";
import "../ERC721/ERC721.sol";
import "./lib/LibDnsToolKit.sol";
import "./tax/IDnsTax.sol";
import "./price/IDnsPrice.sol";
import "./accountant/IDnsAccountant.sol";
import "../ERC20/IERC20.sol";
import "./owners/IDnsDaoOwner.sol";
import "./owners/IDnsOwnerPub.sol";
import "./lib/LibDnsSig.sol";
//import "./INamePub.sol";

contract DnsSubName is ERC721,Erc721Owner{
    using Strings for uint256;
    bytes32 public fatherHash;
    uint256 public gNameId;
    address public taxC;
    address public priceC;
    address public accountantC;
    bytes private baseUri;
    address public fatherAddr;
    address public ownerC;
    uint256 public taxPrice = 5000000000000000000;
    mapping(uint32=>bool) passCardUsed;
    address public sigUser;

    struct NameItem{
        bytes entireName;
        //        bytes32 hash;
        mapping(bytes32=>bytes) subName;
        uint256 expireTime;
        uint256 tokenId;
    }

    mapping(bytes32=>NameItem) public nameStore;
    mapping(uint256=>bytes32) public id2Hash;

    event EvMintSubName(address erc20Addr, string entireName,uint8 year);
    event EvChargeSubName(address erc20Addr, bytes32 nameHash,uint8 year,bool transfer);
    event EvAddSubName(bytes32 hash, string entireName);
    event EvDelSubName(bytes32 hash, string entireName);

    constructor(
        bytes32 fatherHash_,
        address taxC_,
        address priceC_,
        address accountC_,
        address nameOwner_,
        address ownerC_)
    ERC721("","")
    Erc721Owner(nameOwner_){
        gNameId = 1;
        taxC = taxC_;
        priceC = priceC_;
        accountantC = accountC_;
        fatherHash = fatherHash_;
        fatherAddr = msg.sender;
        ownerC = ownerC_;
    }

    //    function getNameInfo(bytes32 hash_) external view override returns(uint256,uint256,address,bool){
    //        return (nameStore[hash_].expireTime,nameStore[hash_].tokenId,address(0),false);
    //    }

    function setBaseUri(string memory baseUri_,address sigUser_) external onlyOwner{
        baseUri = bytes(baseUri_);
        sigUser = sigUser_;
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool){
        return super.supportsInterface(interfaceId);
    }

    //    function ownerOf(uint256 tokenId) public view override returns (address){
    //        return super.ownerOf(tokenId);
    //    }
    //
    //    function balanceOf(address owner) public view override returns (uint256){
    //        return super.balanceOf(owner);
    //    }
    function name() public view override returns (string memory){
        return IDnsOwnerPub(ownerC).getEntireName(fatherHash);
    }

    function symbol() public view override returns (string memory){
        return IDnsOwnerPub(ownerC).getUDIDSymbol(fatherHash);
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
        return super.getApproved(tokenId);
    }

    function setApprovalForAll(address operator, bool approved) public override {
        return super.setApprovalForAll(operator,approved);
    }

    function isApprovedForAll(address owner, address operator) public view override returns (bool){
        return super.isApprovedForAll(owner,operator);
    }

    function _dnsTransfer(address from, address to, uint256 tokenId) internal {
        require(nameStore[id2Hash[tokenId]].expireTime>block.timestamp,"not a valid token");
        require(from != to,"same address error");
        IDnsDaoOwner(ownerC).UpdateOwner(id2Hash[tokenId],to);
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


    function _extendExpire(bytes32 nameHash_,uint8 year_, bool transfer_) internal {
        if (nameStore[nameHash_].expireTime > block.timestamp){
            nameStore[nameHash_].expireTime += year_ * 365 days;
        }else{
            nameStore[nameHash_].expireTime = block.timestamp + year_ * 365 days;
            if(ownerOf(nameStore[nameHash_].tokenId) != msg.sender && transfer_){
                _transfer(ownerOf(nameStore[nameHash_].tokenId),msg.sender,nameStore[nameHash_].tokenId);
                IDnsDaoOwner(ownerC).UpdateOwner(nameHash_,msg.sender);
            }
        }
    }

    function _chargeSubName(address erc20Addr_, bytes32 nameHash_,uint8 year_) internal{
        bytes memory entireName = nameStore[nameHash_].entireName;
        (uint256 taxAmount, uint256 leftAmount,uint256 max) = _calcCost(erc20Addr_,entireName,year_);

        IDnsAccountant(accountantC).deposit(fatherAddr,erc20Addr_,taxAmount);
        if(leftAmount>0){
            IDnsAccountant(accountantC).deposit(erc20Addr_,leftAmount);
        }
        if(erc20Addr_ != address (0)){
            IERC20(erc20Addr_).transferFrom(msg.sender,address(accountantC),max);
        }else{
            payable(accountantC).transfer(max);
        }
    }

    function ChargeName(address erc20Addr_, bytes32 nameHash_,uint8 year_,bool transfer_) external payable{
        require(nameStore[nameHash_].tokenId > 0,"name was not registered");
        _chargeSubName(erc20Addr_,nameHash_,year_);
        _extendExpire(nameHash_,year_,transfer_);
        emit EvChargeSubName(erc20Addr_, nameHash_,year_,transfer_);
    }

    function _mintName(bytes memory entireName_,uint8 year_) internal {
        bytes32 hash = LibDnsToolKit.entireNameHash(string(entireName_));
        nameStore[hash].expireTime = block.timestamp + year_ * 365 days;
        nameStore[hash].entireName = entireName_;
        nameStore[hash].tokenId =gNameId;
        _mint(msg.sender,gNameId);
        id2Hash[gNameId] = hash;
        IDnsDaoOwner(ownerC).UpdateOwner(hash,msg.sender,gNameId,entireName_);
        gNameId++;
    }

    function _calcCost(address erc20Addr_,bytes memory entireName_,uint8 year_) internal returns(uint256,uint256,uint256) {
        (,uint256 defaultTax,,) =
        IDnsTax(taxC).tax(erc20Addr_,LibDnsToolKit.getSubName(bytes(entireName_)).length);
        uint256 price = IDnsPrice(priceC).Price(fatherHash,erc20Addr_,LibDnsToolKit.getSubName(bytes(entireName_)).length)*year_;
        defaultTax = defaultTax*year_;
        uint256 max = LibDnsToolKit.max(price,defaultTax);
        require(max>0,"payment not support");
        if(erc20Addr_ == address(0)){
            require(msg.value>=max,"payout is not enough");
            max = msg.value;
        }else{
            require(IERC20(erc20Addr_).balanceOf(msg.sender)>=max &&
                IERC20(erc20Addr_).allowance(msg.sender,address(this))>= max,"payout is not enough");
        }

        return (defaultTax,max-defaultTax,max);
    }

    function _mintSubName(address erc20Addr_,string memory entireName_,uint8 year_) internal{
        require(LibDnsToolKit.fatherNameHash(bytes(entireName_))==fatherHash,"not a correct sub Name");
        require(nameStore[LibDnsToolKit.entireNameHash(entireName_)].tokenId == 0,"name was registed");

        (uint256 taxAmount, uint256 leftAmount,uint256 max) = _calcCost(erc20Addr_,bytes(entireName_),year_);

        IDnsAccountant(accountantC).deposit(fatherAddr,erc20Addr_,taxAmount);
        if(leftAmount>0){
            IDnsAccountant(accountantC).deposit(erc20Addr_,leftAmount);
        }
        if(erc20Addr_ != address (0)){
            IERC20(erc20Addr_).transferFrom(msg.sender,address(accountantC),max);
        }else{
            payable(accountantC).transfer(max);
        }
    }

    function MintName(address erc20Addr_, string memory entireName_,uint8 year_) external payable{
        require(IDnsOwnerPub(ownerC).isOpened(),"not opened to public");
        _mintSubName(erc20Addr_,entireName_,year_);
        _mintName(bytes(entireName_),year_);
        emit EvMintSubName(erc20Addr_, entireName_,year_);
    }
    function MintNameBySig(string memory entireName_,
        uint8 year_, address erc20Addr_,
        uint256 price_, uint32 passId_,bytes memory sig) external payable{
        // require(price_>=taxPrice,"tax not enough");
        require(LibDnsSignature.SigUserAddr(
            keccak256(abi.encodePacked(entireName_,year_,erc20Addr_,price_,msg.sender,passId_)),
            sig) == sigUser,"sig not correct");
        require(!passCardUsed[passId_],"pass is used");
        passCardUsed[passId_] = true;
        uint256 leftAmount = 0;
        uint256 taxAmount = 0;
        if(price_>taxPrice){
            leftAmount = price_-taxPrice;
            taxAmount = taxPrice;
        }else{
            taxAmount = price_;
        }
        if (taxAmount > 0){
            IDnsAccountant(accountantC).deposit(fatherAddr,erc20Addr_,taxAmount);
        }
        if(leftAmount>0){
            IDnsAccountant(accountantC).deposit(erc20Addr_,leftAmount);
        }
        if(price_>0){
            if(erc20Addr_ != address (0)){
                IERC20(erc20Addr_).transferFrom(msg.sender,address(accountantC),price_);
            }else{
                payable(accountantC).transfer(price_);
            }
        }
        _mintName(bytes(entireName_),year_);
        emit EvMintSubName(erc20Addr_, entireName_,year_);
    }

    // function AddSubName(bytes32 hash_, string memory entireName_) external{
    //     require(nameStore[hash_].tokenId > 0,"name was not registered");
    //     require(super.ownerOf(nameStore[hash_].tokenId)==msg.sender,"user not valid");
    //     require(LibDnsToolKit.getLeve2FatherNameHash(bytes(entireName_))==hash_,"not a valid sub name");

    //     bytes32 entireHash = LibDnsToolKit.entireNameHash(entireName_);

    //     nameStore[hash_].subName[entireHash] = bytes(entireName_);

    //     emit EvAddSubName(hash_, entireName_);
    // }

    // function DelSubName(bytes32 hash_,string memory entireName_) external{
    //     require(nameStore[hash_].tokenId > 0,"name was not registered");
    //     require(super.ownerOf(nameStore[hash_].tokenId)==msg.sender,"user not valid");
    //     require(LibDnsToolKit.getLeve2FatherNameHash(bytes(entireName_))==hash_,"not a valid sub name");
    //     bytes32 entireHash = LibDnsToolKit.entireNameHash(entireName_);

    //     delete nameStore[hash_].subName[entireHash];

    //     emit EvDelSubName(hash_, entireName_);
    // }

}