//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../../ERC721/ERC721.sol";
import "./DnsErc721Owner.sol";
import "./IDnsNameErc721.sol";
import "../IDnsTopLevelName.sol";
import "../../ERC721/utils/Strings.sol";

contract DnsNameErc721 is ERC721,Erc721Owner,IDnsNameErc721{
    using Strings for uint256;
    uint256 public gNameId;
    bytes32 private ownerHash;
    bytes public baseUri;
    address private sigUser;
//    address private operator;

    struct IDInfo{
        uint32 expireTime;
        bytes32 nameHash;
    }

    mapping(uint256=>IDInfo) public nameIdInfo;

    constructor(string memory name_, string memory symbol_,address owner_,bytes32 hash_)
            ERC721(name_,symbol_) Erc721Owner(owner_,msg.sender){
        gNameId = 0;
        ownerHash = hash_;
//        operator = operator_;
    }

    function setBaseUri(string memory baseUri_) external onlyOwner{
        baseUri = bytes(baseUri_);
    }

    function setSigUser(address sigUser_) external onlyOwner{
        sigUser = sigUser_;
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool){
        return super.supportsInterface(interfaceId);
    }

    function _baseURI() internal view override returns (string memory){
        //if baseUri.length == 0 then read from owner contract
        return string(baseUri);
    }

    function tokenURI(uint256 tokenId) public view override returns (string memory){
        super._requireMinted(tokenId);
        string memory baseURI = _baseURI();
        return bytes(baseURI).length > 0 ? string(abi.encodePacked(baseURI, tokenId.toString())) : "";
    }

    function _dnsTransfer(address to, uint256 tokenId) internal{
        require(nameIdInfo[tokenId].expireTime > block.timestamp,"name expired");
        if(ownerHash == bytes32(0)){
            if(IDnsTopLevelName(address(erc721Owner)).getErc721Contract(nameIdInfo[tokenId].nameHash) != address(0)){
                Erc721Owner(IDnsTopLevelName(address(erc721Owner)).getErc721Contract(nameIdInfo[tokenId].nameHash)).transferOwnership(payable(to));
            }
        }
    }

    function approve(address to, uint256 tokenId) public  override{
        return super.approve(to,tokenId);
    }

    function getApproved(uint256 tokenId) public view override returns (address){
        return super.getApproved(tokenId);
    }

    function setApprovalForAll(address operator_, bool approved) public override {
        return super.setApprovalForAll(operator_,approved);
    }

    function isApprovedForAll(address owner, address operator_) public view override returns (bool){
        return super.isApprovedForAll(owner,operator_);
    }

    function transferFrom(address from, address to, uint256 tokenId) public override{
        _dnsTransfer(to,tokenId);
        return super.transferFrom(from,to,tokenId);
    }

    function safeTransferFrom(address from, address to, uint256 tokenId) public override {
        safeTransferFrom(from, to, tokenId, "");
    }

    function safeTransferFrom(address from, address to, uint256 tokenId, bytes memory data) public override {
        _dnsTransfer(to,tokenId);
        super.safeTransferFrom(from,to,tokenId,data);
    }

    function MintId(bytes32 hash_,uint32 expireTime_,address idOwner_) external override returns(uint256) {
        require((ownerHash == bytes32(0) && erc721Owner == msg.sender) ||
            (ownerHash != bytes32(0) && IDnsTopLevelName(erc721Owner).getOperator() == msg.sender), "not allowed");
        gNameId ++;
        _mint(idOwner_,gNameId);
        nameIdInfo[gNameId] = IDInfo(expireTime_,hash_);
        return gNameId;
    }

    function DnsTransfer(address to, uint256 tokenId) external override{
        require((ownerHash == bytes32(0) && erc721Owner == msg.sender) ||
            (ownerHash != bytes32(0) && IDnsTopLevelName(erc721Owner).getOperator() == msg.sender), "not allowed");
        _transfer(ownerOf(tokenId),to,tokenId);
        _dnsTransfer(to,tokenId);
    }

    function DnsExtendExpire(uint256 tokenId_,uint32 expireTime_) external override {
        require((ownerHash == bytes32(0) && erc721Owner == msg.sender) ||
            (ownerHash != bytes32(0) && IDnsTopLevelName(erc721Owner).getOperator() == msg.sender), "not allowed");
        nameIdInfo[tokenId_].expireTime = expireTime_;
    }

    function SigUserAddr() external override view returns(address){
        return sigUser;
    }

}
