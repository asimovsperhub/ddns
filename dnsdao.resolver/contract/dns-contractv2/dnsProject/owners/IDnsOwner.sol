//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IDnsOwner{
//    function UpdateOwner(bytes32 nameHash_, address nameOwner_,uint256 tokenId_, bytes memory entireName_, address contractAddr_) external;
    function UpdateOwner(bytes32 nameHash_, address nameOwner_,uint256 tokenId_, bytes memory entireName_) external;

    function UpdateOwner(bytes32 nameHash_, address nameOwner_, address contractAddr_) external;
    function UpdateOwner(bytes32 nameHash_, address nameOwner_) external;
}