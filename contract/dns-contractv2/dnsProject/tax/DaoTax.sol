//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "./LibDaoTax.sol";
import "./IDaoTax.sol";
import "../payment/IDnsPayment.sol";
import "../owners/IDnsOwnerPub.sol";
import "../daoowner/DaoOwner.sol";

contract DaoTax is IDaoTax,DaoOwner{
    using LibDaoTax for LibDaoTax.NameTax;
    mapping(address=>LibDaoTax.NameTax) nameTaxStore;
    string public name;
    address public paymentC;

    event EvAddTax(address erc20Addr, uint8 taxType, uint256 defaultTax, uint256 percent, uint256 percentBase);
    event EvUpdateTax(address erc20Addr, uint8 taxType, uint256 defaultTax, uint256 percent, uint256 percentBase);
    event EvAddTaxLen(address erc20Addr,uint256 len,uint256 tax,uint256 percent);
    event EvDelTaxLen(address erc20Addr,uint256 len);

    constructor(string memory name_, address paymentC_,address ownerC_) DaoOwner(ownerC_){
         nameTaxStore[address(0)].percentBase = 1e8;
         nameTaxStore[address(0)].taxType = 0;
         nameTaxStore[address(0)].defaultTax = 1e16;
         nameTaxStore[address(0)].taxPercent = 1e6;

        nameTaxStore[address(0)].insert(3,5e16,1e6);
        nameTaxStore[address(0)].insert(4,3e16,1e6);
        nameTaxStore[address(0)].insert(6,1e16,1e6);

        //usdt: 0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD,0,2000000,1000000,100000000
        nameTaxStore[0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].percentBase = 1e8;
        nameTaxStore[0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].taxType = 0;
        nameTaxStore[0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].defaultTax = 2e6;
        nameTaxStore[0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].taxPercent = 1e6;

        nameTaxStore[0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].insert(3,5e6,1e6);
        nameTaxStore[0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].insert(4,3e6,1e6);
        nameTaxStore[0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].insert(6,2e6,1e6);

        paymentC = paymentC_;
        name = name_;
    }

    //add eth AddTax(address(0),0,1e18,1e6,1e18);
    function AddTax(address erc20Addr_, uint8 taxType_, uint256 defaultTax_, uint256 percent_,uint256 percentBase_) external override onlyDaoOwner{
        require(nameTaxStore[erc20Addr_].defaultTax == 0,"erc20 existed");
        (,,,bool isEnable) = IDnsPayment(paymentC).getPaymentByAddr(erc20Addr_);
        require(!isEnable,"not a support erc20 payment");
        nameTaxStore[erc20Addr_].percentBase = percentBase_;
        nameTaxStore[erc20Addr_].taxType = taxType_;
        nameTaxStore[erc20Addr_].defaultTax = defaultTax_;
        nameTaxStore[erc20Addr_].taxPercent = percent_;
        emit EvAddTax(erc20Addr_,taxType_,defaultTax_,percent_,percentBase_);
    }

    function UpdateTax(address erc20Addr_,uint8 taxType_, uint256 defaultTax_, uint256 percent_,uint256 percentBase_) external override onlyDaoOwner{
        require(nameTaxStore[erc20Addr_].defaultTax > 0,"erc20 not exists");
        nameTaxStore[erc20Addr_].percentBase = percentBase_;
        nameTaxStore[erc20Addr_].taxType = taxType_;
        nameTaxStore[erc20Addr_].defaultTax = defaultTax_;
        nameTaxStore[erc20Addr_].taxPercent = percent_;
        emit EvUpdateTax(erc20Addr_,taxType_,defaultTax_,percent_,percentBase_);
    }

    function AddTaxLen(address erc20Addr_,uint256 len_,uint256 tax_,uint256 percent_) external override onlyDaoOwner {
        require(nameTaxStore[erc20Addr_].defaultTax > 0,"erc20 not exists");
        nameTaxStore[erc20Addr_].insert(len_,tax_,percent_);
        emit EvAddTaxLen(erc20Addr_,len_,tax_,percent_);
    }

    function DelTaxLen(address erc20Addr_,uint256 len_) external override onlyDaoOwner {
        require(nameTaxStore[erc20Addr_].defaultTax > 0,"erc20 not exists");
        nameTaxStore[erc20Addr_].remove(len_);
        emit EvDelTaxLen(erc20Addr_,len_);
    }

    function tax(address erc20Addr_,uint256 len_) external override view returns(uint8,uint256,uint256,uint256){
        require(nameTaxStore[erc20Addr_].defaultTax > 0,"erc20 not exists");
        (,,,bool isEnable) = IDnsPayment(paymentC).getPaymentByAddr(erc20Addr_);
        require(!isEnable,"not a support erc20 payment");
        return nameTaxStore[erc20Addr_].tax(len_);
    }

}