//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


interface IDnsAccountant{
    // function setAllowance(address cAddr_) external;
    // function setAccountantOwner(address contractAddr_,address owner_) external;
    function deposit(address erc20Addr_, uint256 amount_) external;
    function deposit(address contractAddr_,address erc20Addr_, uint256 amount_) external;
    function withdrawCash(address erc20Addr_,uint256 amount_,address payable out_, address contractAddr_) external;
    function get(address account_,address erc20Addr_) external returns(uint256);
}