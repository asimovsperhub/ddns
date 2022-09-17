//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

library LibDnsPayment{
    struct DnsPaymentItem {
        address ERC20Addr;
        uint8 decimals;
        string name;
        string symbol;
        bool disabled;
    }

    struct DnsPaymentList {
        address[] arrErc20Addr;
        mapping(address=>DnsPaymentItem) items;
    }


    event evAddPayment(address erc20Addr,uint8 decimal,string name,string symbol);
    event evUpdatePayment(address erc20Addr,string name,string symbol);
    event evUpdateDecimals(address erc20Addr,uint8 decimal);
    event evRemovePayment(address erc20Addr);
    event evDisablePayment(address erc20Addr);
    event evEnablePayment(address erc20Addr);


    function insert(DnsPaymentList storage self, address erc20Addr_, uint8 decimal_,string calldata name_, string calldata symbol_) public {
        require(decimal_>0,"decimal not correct");
        require(self.items[erc20Addr_].decimals == 0,"duplicate payment type");
        self.arrErc20Addr.push(erc20Addr_);
        self.items[erc20Addr_] = DnsPaymentItem(erc20Addr_,decimal_,name_,symbol_,false);
        emit evAddPayment(erc20Addr_,decimal_,name_,symbol_);
    }

    function updateName(DnsPaymentList storage self, address erc20Addr_,string calldata name_,string calldata symbol_) public {
        require(self.items[erc20Addr_].decimals>0,"no payment in dns payment list");
        self.items[erc20Addr_].name = name_;
        self.items[erc20Addr_].symbol = symbol_;
        emit evUpdatePayment(erc20Addr_,name_,symbol_);
    }
    function updateDecimals(DnsPaymentList storage self, address erc20Addr_,uint8 decimal) public{
        require(self.items[erc20Addr_].decimals>0,"no payment in dns payment list");
        self.items[erc20Addr_].decimals = decimal;
        emit evUpdateDecimals(erc20Addr_,decimal);
    }


    function remove(DnsPaymentList storage self,address erc20Addr_) public {
        require(self.items[erc20Addr_].decimals>0,"no payment in dns payment list");

        for (uint i=0;i<self.arrErc20Addr.length;i++){
            if (erc20Addr_ == self.arrErc20Addr[i]){
                self.arrErc20Addr[i] = self.arrErc20Addr[self.arrErc20Addr.length-1];
                self.arrErc20Addr.pop();
                break;
            }
        }

        delete self.items[erc20Addr_];
        emit evRemovePayment(erc20Addr_);
    }

    function disable(DnsPaymentList storage self, address erc20Addr_) public {
        require(self.items[erc20Addr_].decimals>0,"no payment in dns payment list");

        self.items[erc20Addr_].disabled = true;

        emit evDisablePayment(erc20Addr_);
    }

    function enable(DnsPaymentList storage self, address erc20Addr_) public{
        require(self.items[erc20Addr_].decimals>0,"no payment in dns payment list");

        self.items[erc20Addr_].disabled = false;

        emit evEnablePayment(erc20Addr_);
    }

    function getPaymentCount(DnsPaymentList storage self) external view returns (uint256){
        return self.arrErc20Addr.length;
    }

    function getPayment(DnsPaymentList storage self, uint256 index_) external view returns (string memory,string memory,address,uint8,bool){
        require(index_ < self.arrErc20Addr.length,"index out of range");
        DnsPaymentItem memory item = self.items[self.arrErc20Addr[index_]];
        return (item.name,item.symbol,item.ERC20Addr,item.decimals,item.disabled);
    }

    function getPaymentByAddr(DnsPaymentList storage self,address erc20Addr_) external view returns(string memory,string memory,uint8,bool){
        require(self.items[erc20Addr_].decimals>0,"erc20 address not exists");
        DnsPaymentItem memory item = self.items[erc20Addr_];
        return (item.name,item.symbol,item.decimals,item.disabled);
    }

}