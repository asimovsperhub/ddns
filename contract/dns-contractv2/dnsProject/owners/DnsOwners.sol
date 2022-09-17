//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./IDnsOwnerPub.sol";
import "./IDnsDaoOwner.sol";
import "../../ERC721/utils/Strings.sol";
import "../../utils/Owner.sol";

contract DnsOwners is IDnsOwnerPub,IDnsDaoOwner,owned{

    using Strings for uint256;
    bool public open2Reg;

    struct OwnerItem {
        address dnsOwner;
        uint256 tokenId;
        bytes entireName;
    }

    mapping(bytes32=>OwnerItem) public dnsOwners;

    mapping(address=>bytes32) public contractAddrs;

    mapping(address=>bool) public allowance;
    // mapping(uint32=>bool) public passCardUsed;
    // mapping(address=>address) public SigUser;

    event EvUpdateOwner(bytes32 nameHash, address nameOwner, address contractAddr);
    event EvAddAllowance(address user);

    constructor(address owner_){
        allowance[msg.sender] = true;
        dnsOwners[bytes32(0)] = OwnerItem(owner_,0,new bytes(0));
        contractAddrs[msg.sender] = bytes32(0);
    }

    modifier allowed{
        require(allowance[msg.sender],"not allowed");
        _;
    }



    function getEntireName(bytes32 nameHash_) external view override returns(string memory){
        return string(dnsOwners[nameHash_].entireName);
    }
    function getUDIDSymbol(bytes32 nameHash_) external view override returns(string memory){
        return string(abi.encodePacked("UDID#",dnsOwners[nameHash_].tokenId.toString()));
    }

    function UpdateOwner(bytes32 nameHash_, address nameOwner_, address contractAddr_) external override allowed{
        dnsOwners[nameHash_].dnsOwner = nameOwner_;
        contractAddrs[contractAddr_] = nameHash_;
        emit EvUpdateOwner(nameHash_,nameOwner_,contractAddr_);
    }

    function UpdateOwner(bytes32 nameHash_, address nameOwner_,uint256 tokenId_, bytes memory entireName_) external override allowed{
        dnsOwners[nameHash_] = OwnerItem(nameOwner_,tokenId_,entireName_);
        emit EvUpdateOwner(nameHash_,nameOwner_,address(0));
    }

    function UpdateOwner(bytes32 nameHash_, address nameOwner_) external override allowed{
        dnsOwners[nameHash_].dnsOwner = nameOwner_;
        emit EvUpdateOwner(nameHash_,nameOwner_,address(0));
    }

    function addAllowance(address user_) external override onlyOwner{
        allowance[user_] = true;
        emit EvAddAllowance(user_);
    }

    function openLevel2Reg(bool o) external override onlyOwner{
        open2Reg = o;
    }

    function isOpened() external view override returns (bool){
        return open2Reg;
    }

    function getOwnerByHash(bytes32 nameHash_) external view override returns(address){
        return dnsOwners[nameHash_].dnsOwner;
    }

    function getOwnerByContract(address contractAddr_) external view override returns (address){
        return dnsOwners[contractAddrs[contractAddr_]].dnsOwner;
    }

    function operatorIsAllowance(address contractAddr_) external view override returns(bool){
        return allowance[contractAddr_];
    }

}