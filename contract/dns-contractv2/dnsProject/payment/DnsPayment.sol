//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./IDnsPayment.sol";
import "./LibDnsPayment.sol";
import "../../utils/Owner.sol";
import "../owners/IDnsOwnerPub.sol";
import "../daoowner/DaoOwner.sol";


contract DnsPayment is IDnsPayment,DaoOwner{
    using LibDnsPayment for LibDnsPayment.DnsPaymentList;
    LibDnsPayment.DnsPaymentList private paymentList;
    string public dnsPaymentName;

    constructor(string memory name_,address ownerC_) DaoOwner(ownerC_){
        dnsPaymentName = name_;
        paymentList.insert(address(0),18,"ethereum","eth");
        paymentList.insert(0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD,6,"Tether USD","USDT");
    }

    function getPaymentCount() external override view returns (uint256){
        return paymentList.getPaymentCount();
    }
    function getPayment(uint256 index_) external override view returns (string memory,string memory,address,uint8,bool){
        return paymentList.getPayment(index_);
    }
    function getPaymentByAddr(address erc20Addr_) external override view returns(string memory,string memory,uint8,bool){
        return paymentList.getPaymentByAddr(erc20Addr_);
    }

    function addPayment(address erc20Addr_, uint8 decimal_,string calldata name_, string calldata symbol_) external onlyDaoOwner{
        return paymentList.insert(erc20Addr_,decimal_,name_,symbol_);
    }

    function updateName( address erc20Addr_,string calldata name_,string calldata symbol_) external onlyDaoOwner {
        return paymentList.updateName(erc20Addr_,name_,symbol_);
    }
    function updateDecimals(address erc20Addr_,uint8 decimal) external onlyDaoOwner{
        return paymentList.updateDecimals(erc20Addr_,decimal);
    }

    function disable(address erc20Addr_) external onlyDaoOwner {
        return paymentList.disable(erc20Addr_);
    }

    function enable(address erc20Addr_) external onlyDaoOwner{
        return paymentList.enable(erc20Addr_);
    }

}