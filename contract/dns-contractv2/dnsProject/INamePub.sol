//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;



interface INamePub{
    function getNameInfo(bytes32 hash_) external view returns(uint256,uint256,address,bool);
}