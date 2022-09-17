//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./LibDnsAccountant.sol";
import "./IDnsAccountant.sol";
import "../../ERC20/IERC20.sol";
import "../owners/IDnsOwnerPub.sol";
import "../daoowner/DaoOwner.sol";
import "../multisig/MultiSig.sol";

contract DnsAccountant is IDnsAccountant,MultiSig{
    using LibDnsAccountant for LibDnsAccountant.NameAccountant;
    LibDnsAccountant.NameAccountant private nameAccountant;

    //1. 帐记录在合约上
    //2. 所有的cash保存在this上
    //3. 合约可以记录自己的帐，也可以记录到国库的帐
    //4. 只有合约所有人才可以提cash

    event EvDeposit(address operator, address contractAddr,address erc20Addr, uint256 amount);
    event EvWithdrawCash(address operator,address erc20Addr,uint256 amount,address out, address contractAddr);

    constructor(address ownerC_) MultiSig(ownerC_){}

    function deposit(address erc20Addr_, uint256 amount_) external override operatorIsAllowed{
        nameAccountant.deposit(msg.sender,erc20Addr_,amount_);
        emit EvDeposit(msg.sender,msg.sender,erc20Addr_,amount_);
    }

    function deposit(address contractAddr_,address erc20Addr_, uint256 amount_) external override operatorIsAllowed{
        nameAccountant.deposit(contractAddr_,erc20Addr_,amount_);
        emit EvDeposit(msg.sender,contractAddr_,erc20Addr_,amount_);
    }

    function withdrawCash(address erc20Addr_,uint256 amount_,address payable out_, address contractAddr_)
    external
    override
//    onlyContractOwner(contractAddr_)
    sigMatched(contractAddr_,keccak256(abi.encodePacked(erc20Addr_,amount_,out_,contractAddr_,this.withdrawCash.selector))){
        require(nameAccountant.get(contractAddr_,erc20Addr_)>=amount_,"amount not enough");
        if(erc20Addr_ == address(0)){
            require(address(this).balance >= amount_,"amount not enough");
            out_.transfer(amount_);
        }else{
            require(IERC20(erc20Addr_).balanceOf(address(this))>=amount_,"not engough");
            IERC20(erc20Addr_).transfer(out_,amount_);
        }
        nameAccountant.withdraw(contractAddr_,erc20Addr_,amount_);
        emit EvWithdrawCash(msg.sender,erc20Addr_,amount_,out_,contractAddr_);
    }

    function get(address account_,address erc20Addr_) external override view returns(uint256){
        return nameAccountant.get(account_,erc20Addr_);
    }

    receive()  payable external {}
    fallback() payable external{}

}