//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IDnsErc721Owner{
    function transferOwnership(address payable newOwner) external;
}