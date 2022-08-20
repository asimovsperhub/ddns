//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

library LibDnsPrice{
    struct NameItem{
        uint256 idx;
        uint256 price;
    }
    struct NamePrice{
        uint256 defaultPrice;
        uint256[] arrLens;
        mapping(uint256=>NameItem) lenPrice;
    }

    function init(NamePrice storage self,uint256 price_) public{
        self.defaultPrice = price_;
    }

    function isInit(NamePrice storage self) public view returns(bool){
        if (self.defaultPrice > 0){
            return true;
        }
        return false;
    }

    function insert(NamePrice storage self,uint256 len_,uint256 price_) public {
        if(self.lenPrice[len_].price>0){
            self.lenPrice[len_].price = price_;
            return;
        }
        self.arrLens.push(len_);
        int256 idx = -1;
        for(int256 i=0;i<int256(self.arrLens.length);i++){
            if (self.arrLens[uint256(i)] > len_){
                idx = i;
                break;
            }
        }
        if(idx>=0){
            for(int256 i=int256(self.arrLens.length-1);i>idx;i--){
                self.arrLens[uint256(i)] = self.arrLens[uint256(i-1)];
                self.lenPrice[self.arrLens[uint256(i)]].idx = uint256(i);
            }
            self.arrLens[uint256(idx)] = len_;
        }else{
            idx = int256(self.arrLens.length -1);
        }

        self.lenPrice[len_] = NameItem(uint256(idx),price_);

    }

    function remove(NamePrice storage self,uint256 len_) public {
        require(self.lenPrice[len_].price>0,"not exists");
        uint256 idx = self.lenPrice[len_].idx;


        for(uint256 i=idx;i<self.arrLens.length-1;i++){
            self.arrLens[i] = self.arrLens[i+1];
            self.lenPrice[self.arrLens[i]].idx = i;
        }
        self.arrLens.pop();

        delete self.lenPrice[len_];
    }


    function price(NamePrice storage self,uint256 len_) public view returns(uint256){
        if(self.lenPrice[len_].price>0){
            return self.lenPrice[len_].price;
        }
        if(self.arrLens.length==0){
            return self.defaultPrice;
        }
        for(uint256 i=0;i<self.arrLens.length;i++){
            if(self.arrLens[i]>len_){
                if (i==0){
                    return self.lenPrice[self.arrLens[i]].price;
                }else{
                    return self.lenPrice[self.arrLens[i-1]].price;
                }

            }
        }
        return self.defaultPrice;
    }

    function getDefaultPrice(NamePrice storage self) public view returns(uint256){
        return self.defaultPrice;
    }

}