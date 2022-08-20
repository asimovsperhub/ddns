//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "../lib/LibDnsToolKit.sol";

library LibDnsConf{

    struct ConfTypeStatus{
        bool status;
        uint256 idx;
    }

    struct NameConfMap{
        bytes[] confTypes;
        mapping(bytes32=>ConfTypeStatus) isTypeExists;
        mapping(bytes32=>mapping(uint256=>bytes)) confItem;
    }

    event EvAddType(bytes typeName);
    event EvAddConfMap(bytes32 nameHash, bytes name, bytes value);
    event EvUpdateConfMap(bytes32 nameHash, bytes name, bytes value);
    event EvRemoveConfMap(bytes32 nameHash,bytes name);

    function addType(NameConfMap storage self,bytes memory tName_) public {
        bytes32  hash = keccak256(abi.encodePacked(tName_));
        require(self.isTypeExists[hash].status==false,"type is exists");
        self.confTypes.push(tName_);
        self.isTypeExists[hash] = ConfTypeStatus(true,self.confTypes.length-1);

        emit EvAddType(tName_);
    }

    function addMap(NameConfMap storage self, bytes32 nameHash_,bytes memory tName_, bytes memory v_) public{
        bytes32  hash = keccak256(abi.encodePacked(tName_));
        require(self.isTypeExists[hash].status == true,"type not found");
        self.confItem[nameHash_][self.isTypeExists[hash].idx] = v_;

        emit EvAddConfMap(nameHash_,tName_,v_);
    }

    function updateMap(NameConfMap storage self, bytes32 nameHash_,bytes memory tName_, bytes memory v_) public{
        bytes32  hash = keccak256(abi.encodePacked(tName_));
        require(self.isTypeExists[hash].status == true,"type not found");

        self.confItem[nameHash_][self.isTypeExists[hash].idx] = v_;
        emit EvUpdateConfMap(nameHash_,tName_,v_);
    }

    function removeMap(NameConfMap storage self,bytes32 nameHash_,bytes memory tName_) public{
        bytes32  hash = keccak256(abi.encodePacked(tName_));
        require(self.isTypeExists[hash].status == true,"type not found");
        require(self.confItem[nameHash_][self.isTypeExists[hash].idx].length >0, "no config");
        delete self.confItem[nameHash_][self.isTypeExists[hash].idx];

        emit EvRemoveConfMap(nameHash_,tName_);
    }

    function getMap(NameConfMap storage self,bytes32 nameHash_, bytes memory tName_) public view returns (bool, bytes memory){
        bytes32  hash = keccak256(abi.encodePacked(tName_));
        if(self.isTypeExists[hash].status == true && self.confItem[nameHash_][self.isTypeExists[hash].idx].length >0){
            return (true,self.confItem[nameHash_][self.isTypeExists[hash].idx]);
        }

        return (false,new bytes(0));
    }
    //bytes2Bytes32
    function getAllTypes(NameConfMap storage self) public view returns(bytes memory){

        bytes memory r = new bytes(self.confTypes.length*32);

        for(uint256 i=0;i<self.confTypes.length;i++){
            r = abi.encodePacked(r,LibDnsToolKit.bytes2Bytes32(self.confTypes[i]));
        }
        return r;
    }

}