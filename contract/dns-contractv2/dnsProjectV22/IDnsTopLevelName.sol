//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


interface IDnsTopLevelName {
    function getErc721Contract(bytes32 fatherHash) external view returns(address);
    function getOperator() external view returns(address);
//    function getAllowed(bytes32 fatherHash_) external override view returns(address);
}