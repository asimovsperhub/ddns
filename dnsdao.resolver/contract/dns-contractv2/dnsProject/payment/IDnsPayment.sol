//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


interface IDnsPayment{
    function getPaymentCount() external view returns (uint256);
    function getPayment(uint256 index_) external view returns (string memory,string memory,address,uint8,bool);
    function getPaymentByAddr(address erc20Addr_) external view returns(string memory,string memory,uint8,bool);
}



