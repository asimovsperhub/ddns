//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IDnsTax{
    function tax(address erc20Addr_,uint256 len_) external view returns(uint8,uint256,uint256,uint256);
}