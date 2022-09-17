//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


interface IDnsPrice{
    function Price(bytes32 hash_,address erc20Addr_,uint256 len_) external view returns(uint256);
}