//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


interface IDnsAccountant{
    function deposit(address contractAddr_,address erc20Addr_, uint256 amount_) external;
}