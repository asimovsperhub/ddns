//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


interface IDnsOwnerPub{
    function getOwnerByHash(bytes32 nameHash_) external view returns(address);
    function getOwnerByContract(address contractAddr_) external view returns (address);
    function operatorIsAllowance(address contractAddr_) external view returns(bool);
    function getEntireName(bytes32 nameHash_) external view returns(string memory);
    function getUDIDSymbol(bytes32 nameHash_) external view returns(string memory);
    function isOpened() external returns (bool);
}
