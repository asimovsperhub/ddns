//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./LibDnsConf.sol";
import "../owners/IDnsOwnerPub.sol";
import "../daoowner/DaoOwner.sol";

contract DnsConf is DaoOwner{
    using LibDnsConf for LibDnsConf.NameConfMap;
    LibDnsConf.NameConfMap private nameConfStore;

    event EvAddType(bytes tName);
    event EvAddMap(bytes32 nameHash, bytes tName, bytes val);
    event EvUpdateMap(bytes32 nameHash, bytes tName, bytes val);
    event RemoveMap(bytes32 nameHash, bytes tName);

    constructor(address owner_) DaoOwner(owner_){
        nameConfStore.addType("ipv4");
        nameConfStore.addType("ipv6");
        nameConfStore.addType("eth");
        nameConfStore.addType("btc");
        nameConfStore.addType("sol");
        nameConfStore.addType("etc");
        nameConfStore.addType("trx");
        nameConfStore.addType("eos");
        nameConfStore.addType("bnb");
        nameConfStore.addType("fil");
        nameConfStore.addType("resolver");
        nameConfStore.addType("cname");
    }

    function addType(bytes memory tName_) external onlyDaoOwner{
        require(tName_.length < 32,"type name length too long");
        nameConfStore.addType(tName_);
        emit EvAddType(tName_);
    }

    function addMap( bytes32 nameHash_,bytes memory tName_, bytes memory v_) external onlyNameOwner(nameHash_){
        nameConfStore.addMap(nameHash_,tName_,v_);
        emit EvAddMap(nameHash_,tName_,v_);
    }
    function updateMap( bytes32 nameHash_,bytes memory tName_, bytes memory v_) external onlyNameOwner(nameHash_){
        nameConfStore.updateMap(nameHash_,tName_,v_);
        emit EvUpdateMap(nameHash_,tName_,v_);
    }

    function removeMap(bytes32 nameHash_,bytes memory tName_) external onlyNameOwner(nameHash_){
        nameConfStore.removeMap(nameHash_,tName_);
        emit RemoveMap(nameHash_,tName_);
    }

    function getMap(bytes32 nameHash_,bytes memory tName_) external view returns(bool,bytes memory){
        return nameConfStore.getMap(nameHash_,tName_);
    }

    function getAllTypes() external view returns(bytes memory){
        return nameConfStore.getAllTypes();
    }

}