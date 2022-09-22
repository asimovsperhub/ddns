//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./LibDnsAccountant.sol";
import "./IDnsAccountant.sol";
import "../../ERC20/IERC20.sol";
import "../multisig/MultiSig.sol";
import "../IDnsTopLevelName.sol";
import "../../utils/Owner.sol";

contract DnsAccountant is IDnsAccountant,MultiSig{
    using LibDnsAccountant for LibDnsAccountant.NameAccountant;
    LibDnsAccountant.NameAccountant private nameAccountant;
    IDnsTopLevelName public dnsTop;


    //1. revenue record in every erc721 contract address
    //2. all cash asset transfer to this
    //3. all revenue are recorded in this contract include dnsDao revenue
    //4. only erc721 contract owner can withdraw cash

    event EvDeposit(address operator, address contractAddr,address erc20Addr, uint256 amount);
    event EvWithdrawCash(address operator,address erc20Addr,uint256 amount,address out, address contractAddr);

    constructor(address dnsTop_) MultiSig(){
        dnsTop = IDnsTopLevelName(dnsTop_);
    }

    function setRelation(address dnsTop_) external{
        require(msg.sender == owned(address(dnsTop)).owner());
        dnsTop = IDnsTopLevelName(dnsTop_);
    }


    function deposit(address contractAddr_,address erc20Addr_, uint256 amount_) external override{
        require(msg.sender == address(dnsTop) || msg.sender == dnsTop.getOperator(),"not allowed");
        nameAccountant.deposit(contractAddr_,erc20Addr_,amount_);
        emit EvDeposit(msg.sender,contractAddr_,erc20Addr_,amount_);
    }

    function withdrawCash(address erc20Addr_,uint256 amount_,address payable out_, address contractAddr_)
    external
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

    function get(address account_,address erc20Addr_) external view returns(uint256){
        return nameAccountant.get(account_,erc20Addr_);
    }

    receive()  payable external {}
    fallback() payable external{}

}