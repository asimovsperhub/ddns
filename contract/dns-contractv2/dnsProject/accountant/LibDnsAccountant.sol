//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;



library LibDnsAccountant{
    struct NameAccountant {
        mapping(address=>mapping(address=>uint256)) accountant;
    }

    event EvDeposit(address account,address erc20Addr,uint256 amount);
    event EvWithdraw(address account,address erc20Addr,uint256 amount);

    function deposit(NameAccountant storage self, address account_,address erc20Addr_, uint256 amount_) public {
        self.accountant[account_][erc20Addr_] += amount_;
        emit EvDeposit(account_,erc20Addr_,amount_);
    }

    function withdraw(NameAccountant storage self,address account_,address erc20Addr_, uint256 amount_) public{
        require(self.accountant[account_][erc20Addr_] >= amount_,"amount not enough");
        self.accountant[account_][erc20Addr_] -= amount_;
        emit EvWithdraw(account_,erc20Addr_,amount_);
    }

    function get(NameAccountant storage self,address account_,address erc20Addr_) public view returns(uint256){
        return self.accountant[account_][erc20Addr_];
    }

}




