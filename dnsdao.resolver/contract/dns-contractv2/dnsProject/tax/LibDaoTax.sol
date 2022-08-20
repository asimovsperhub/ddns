//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


library LibDaoTax{
    struct NameItem{
        uint256 idx;
        uint256 tax;
        uint256 taxPercent;
    }
    struct NameTax{
        uint256 percentBase;
        uint8 taxType;  //0,1,2
        uint256 defaultTax;
        uint256 taxPercent;
        uint256[] arrLens;
        mapping(uint256=>NameItem) lenTax;
    }

    // function updateTaxType(NameTax storage self, uint8 taxType_) public{
    //     self.taxType = taxType_;
    // }

    // function updateDefaultTax(NameTax storage self,uint256 tax_,uint256 percent_) public{
    //     self.defaultTax = tax_;
    //     self.taxPercent = percent_;
    // }

    // function updatePercentBase(NameTax storage self,uint256 base_) public{
    //     self.percentBase = base_;
    // }
    function update(NameTax storage self, uint8 taxType_,uint256 tax_,uint256 percent_,uint256 percentBase_) public{
        self.taxType = taxType_;
        self.defaultTax = tax_;
        self.taxPercent = percent_;
        self.percentBase = percentBase_;
    }

    function insert(NameTax storage self,uint256 len_,uint256 tax_,uint256 percent_) public {
        if(self.lenTax[len_].tax>0){
            self.lenTax[len_].tax = tax_;
            self.lenTax[len_].taxPercent = percent_;
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
                self.lenTax[self.arrLens[uint256(i)]].idx = uint256(i);
            }
            self.arrLens[uint256(idx)] = len_;
        }else{
            idx = int256(self.arrLens.length-1);
        }

        self.lenTax[len_] = NameItem(uint256(idx),tax_,percent_);

    }

    function remove(NameTax storage self,uint256 len_) public {
        require(self.lenTax[len_].tax>0,"not exists");
        uint256 idx = self.lenTax[len_].idx;


        for(uint256 i=idx;i<self.arrLens.length-1;i++){
            self.arrLens[i] = self.arrLens[i+1];
            self.lenTax[self.arrLens[i]].idx = i;
        }
        self.arrLens.pop();

        delete self.lenTax[len_];
    }

    function tax(NameTax storage self,uint256 len_) public view returns(uint8,uint256,uint256,uint256){
        if(self.lenTax[len_].tax>0){
            return (self.taxType,self.lenTax[len_].tax,self.lenTax[len_].taxPercent,self.percentBase);
        }
        if(self.arrLens.length==0){
            return (self.taxType,self.defaultTax,self.taxPercent,self.percentBase);
        }
        for(uint256 i=0;i<self.arrLens.length;i++){
            if(self.arrLens[i]>len_){
                if (i==0){
                    return (self.taxType,self.lenTax[self.arrLens[i]].tax,self.lenTax[self.arrLens[i]].taxPercent,self.percentBase);
                }else{
                    return (self.taxType,self.lenTax[self.arrLens[i-1]].tax,self.lenTax[self.arrLens[i-1]].taxPercent,self.percentBase);
                }

            }
        }
        return (self.taxType,self.defaultTax,self.taxPercent,self.percentBase);
    }

}