//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./LibDnsPrice.sol";
import "./IDnsPrice.sol";
import "../owners/IDnsOwnerPub.sol";
import "../payment/IDnsPayment.sol";
import "../daoowner/DaoOwner.sol";


contract DnsPrice is IDnsPrice,DaoOwner{
    using LibDnsPrice for LibDnsPrice.NamePrice;
    mapping(bytes32=>mapping(address=>LibDnsPrice.NamePrice)) public namePrice;
    address public paymentC;

    constructor(address owner_,address paymentC_) DaoOwner(owner_){
        //0x0000000000000000000000000000000000000000000000000000000000000000,0x0000000000000000000000000000000000000000,  200000000000000000
//        namePrice[bytes32(0)][address(0)].defaultPrice = 2e17;
//        //0x0000000000000000000000000000000000000000000000000000000000000000,0x0000000000000000000000000000000000000000,3,2000000000000000000
//        //0x0000000000000000000000000000000000000000000000000000000000000000,0x0000000000000000000000000000000000000000,4,1000000000000000000
//        //0x0000000000000000000000000000000000000000000000000000000000000000,0x0000000000000000000000000000000000000000,5,500000000000000000
//        //0x0000000000000000000000000000000000000000000000000000000000000000,0x0000000000000000000000000000000000000000,6,200000000000000000
//
//        namePrice[bytes32(0)][address(0)].insert(3,3e18);
//        namePrice[bytes32(0)][address(0)].insert(4,2e18);
//        namePrice[bytes32(0)][address(0)].insert(5,1e18);
//        namePrice[bytes32(0)][address(0)].insert(7,2e17);
//        //0x0000000000000000000000000000000000000000000000000000000000000000,0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD,200000000
//        namePrice[bytes32(0)][0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].defaultPrice = 2e8; //200$
//        //0x0000000000000000000000000000000000000000000000000000000000000000,0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD,3,3000000000
//        //0x0000000000000000000000000000000000000000000000000000000000000000,0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD,4,2000000000
//        //0x0000000000000000000000000000000000000000000000000000000000000000,0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD,6,1000000000
//        namePrice[bytes32(0)][0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].insert(3,3e9);
//        namePrice[bytes32(0)][0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].insert(4,2e9);
//        namePrice[bytes32(0)][0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].insert(5,1e9);
//        namePrice[bytes32(0)][0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].insert(7,2e8);

        _priceInitialTop(bytes32(0));

        paymentC = paymentC_;
    }

    event EvAddPrice(bytes32 nameHash,address erc20Addr,uint256 defaultPrice);
    event EvAddPriceLen(bytes32 nameHash,address erc20Addr,uint256 len, uint256 price);
    event EvDelPriceLen(bytes32 nameHash,address erc20Addr,uint256 len);
    event EvRemovePrice(bytes32 nameHash,address erc20Addr);
    event EvAddPriceLenArray(bytes32 nameHash,address erc20Addr,uint256 defaultPrice, uint256 len,uint256 price);


    function _priceInitialTop(bytes32 hash_) internal {
        namePrice[hash_][address(0)].defaultPrice = 2e17;
        //0x0000000000000000000000000000000000000000000000000000000000000000,0x0000000000000000000000000000000000000000,3,2000000000000000000
        //0x0000000000000000000000000000000000000000000000000000000000000000,0x0000000000000000000000000000000000000000,4,1000000000000000000
        //0x0000000000000000000000000000000000000000000000000000000000000000,0x0000000000000000000000000000000000000000,5,500000000000000000
        //0x0000000000000000000000000000000000000000000000000000000000000000,0x0000000000000000000000000000000000000000,6,200000000000000000

        namePrice[hash_][address(0)].insert(3,3e18);
        namePrice[hash_][address(0)].insert(4,2e18);
        namePrice[hash_][address(0)].insert(5,1e18);
        namePrice[hash_][address(0)].insert(7,2e17);
        //0x0000000000000000000000000000000000000000000000000000000000000000,0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD,200000000
        namePrice[hash_][0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].defaultPrice = 2e8; //200$
        //0x0000000000000000000000000000000000000000000000000000000000000000,0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD,3,3000000000
        //0x0000000000000000000000000000000000000000000000000000000000000000,0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD,4,2000000000
        //0x0000000000000000000000000000000000000000000000000000000000000000,0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD,6,1000000000
        namePrice[hash_][0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].insert(3,3e9);
        namePrice[hash_][0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].insert(4,2e9);
        namePrice[hash_][0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].insert(5,1e9);
        namePrice[hash_][0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].insert(7,2e8);
    }

    function _priceInitialSub(bytes32 hash_) internal{
        //0.02
        namePrice[hash_][address(0)].defaultPrice = 2e16;
        //0.07
        namePrice[hash_][address(0)].insert(3,7e16);
        //0.04
        namePrice[hash_][address(0)].insert(4,4e16);
        //0.03
        namePrice[hash_][address(0)].insert(5,3e16);
        //0.02
        namePrice[hash_][address(0)].insert(7,2e16);

        namePrice[hash_][0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].defaultPrice = 2e7; //20$
        //0x0000000000000000000000000000000000000000000000000000000000000000,0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD,3,3000000000
        //0x0000000000000000000000000000000000000000000000000000000000000000,0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD,4,2000000000
        //0x0000000000000000000000000000000000000000000000000000000000000000,0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD,6,1000000000
        namePrice[hash_][0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].insert(3,64e7); //640$
        namePrice[hash_][0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].insert(4,2e8);  //200$
        namePrice[hash_][0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].insert(5,1e8);  //100$
        namePrice[hash_][0x3B00Ef435fA4FcFF5C209a37d1f3dcff37c705aD].insert(7,2e7);  //20$
    }

    function PriceInit(bytes32 hash_) external onlyNameOwner(hash_){
        _priceInitialSub(hash_);
    }

    function AddPrice(bytes32 hash_,address erc20Addr_,uint256 defaultPrice_) external onlyNameOwner(hash_){
        require(namePrice[hash_][erc20Addr_].defaultPrice == 0,"price existed");
        //        (,,,bool isEnable) = IDnsPayment(paymentC).getPaymentByAddr(erc20Addr_);
        //        require(!isEnable,"not a support erc20 payment");
        //        namePrice[hash_][erc20Addr_].defaultPrice = defaultPrice_;
        _addPrice(hash_,erc20Addr_,defaultPrice_);
        emit EvAddPrice(hash_,erc20Addr_,defaultPrice_);
    }

    function _addPrice(bytes32 hash_,address erc20Addr_,uint256 defaultPrice_) internal {
        (,,,bool isEnable) = IDnsPayment(paymentC).getPaymentByAddr(erc20Addr_);
        require(!isEnable,"not a support erc20 payment");
        namePrice[hash_][erc20Addr_].defaultPrice = defaultPrice_;
    }

    function AddPriceLen(bytes32 hash_,address erc20Addr_,uint256 len_, uint256 price_) external onlyNameOwner(hash_){
        require(namePrice[hash_][erc20Addr_].defaultPrice > 0,"price not exitsts");
        _addPriceLen(hash_,erc20Addr_,len_,price_);
        emit EvAddPriceLen(hash_,erc20Addr_,len_,price_);
    }

    function _addPriceLen(bytes32 hash_,address erc20Addr_, uint256 len_, uint256 price_) internal{
        namePrice[hash_][erc20Addr_].insert(len_,price_);
    }

    function AddPriceLenArray(bytes32 hash_,address erc20Addr_, uint256 defaultPrice_,
        uint256[] calldata lenArr_, uint256[] calldata priceArr_) external onlyNameOwner(hash_){
        _addPrice(hash_,erc20Addr_,defaultPrice_);
        for(uint256 i=0;i<lenArr_.length && i<priceArr_.length; i++){
            _addPriceLen(hash_,erc20Addr_,lenArr_[i],priceArr_[i]);
            emit EvAddPriceLenArray(hash_,erc20Addr_,defaultPrice_,lenArr_[i],priceArr_[i]);
        }

    }

    function DelPriceLen(bytes32 hash_,address erc20Addr_, uint256 len_) external onlyNameOwner(hash_){
        require(namePrice[hash_][erc20Addr_].defaultPrice > 0,"price not exitsts");
        namePrice[hash_][erc20Addr_].remove(len_);
        emit EvDelPriceLen(hash_,erc20Addr_,len_);
    }

    function RemovePrice(bytes32 hash_,address erc20Addr_) external onlyNameOwner(hash_){
        require(namePrice[hash_][erc20Addr_].defaultPrice > 0,"price not exitsts");
        delete namePrice[hash_][erc20Addr_];
        emit EvRemovePrice(hash_,erc20Addr_);
    }

    function Price(bytes32 hash_,address erc20Addr_,uint256 len_) external override view returns(uint256) {
        require(namePrice[hash_][erc20Addr_].defaultPrice > 0,"price not exitsts");
        (,,,bool isEnable) = IDnsPayment(paymentC).getPaymentByAddr(erc20Addr_);
        require(!isEnable,"not a support erc20 payment");
        return namePrice[hash_][erc20Addr_].price(len_);
    }

    function PriceGet(bytes32 hash_, address erc20Addr_,uint256[8] calldata lenArr_) external view returns(uint256[8] memory){
        require(namePrice[hash_][erc20Addr_].defaultPrice > 0,"price not exitsts");
        (,,,bool isEnable) = IDnsPayment(paymentC).getPaymentByAddr(erc20Addr_);
        require(!isEnable,"not a support erc20 payment");
        uint256[8] memory r;
        for(uint256 i=0;i<lenArr_.length;i++){
            if (lenArr_[i] == 0){
                break;
            }
            r[i] = namePrice[hash_][erc20Addr_].lenPrice[lenArr_[i]].price;
        }

        return r;
    }

    function lenIsExists(bytes32 hash_,address erc20Addr_,uint256 len_) internal view returns(bool){
        for(uint i=0;i<namePrice[hash_][erc20Addr_].arrLens.length;i++){
            if(namePrice[hash_][erc20Addr_].arrLens[i] == len_){
                return true;
            }
        }
        return false;
    }
    function Erc20IsSupport(bytes32 hash_,address erc20Addr_) external view returns(bool){
        if (namePrice[hash_][erc20Addr_].defaultPrice > 0){
            return true;
        }
        return false;
    }
    function PriceGetOne(bytes32 hash_,address erc20Addr_,uint256 len_) external view returns(bool,uint256) {
        require(namePrice[hash_][erc20Addr_].defaultPrice > 0,"price not exitsts");
        (,,,bool isEnable) = IDnsPayment(paymentC).getPaymentByAddr(erc20Addr_);
        require(!isEnable,"not a support erc20 payment");
        if(lenIsExists(hash_,erc20Addr_,len_)){
            return (true,namePrice[hash_][erc20Addr_].lenPrice[len_].price);
        }else{
            return (false,0);
        }

    }

}