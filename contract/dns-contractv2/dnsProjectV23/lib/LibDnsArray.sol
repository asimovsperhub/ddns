//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;



library LibBytes32Array{
    function add(bytes32[] storage self, bytes32 hash_) external returns(uint256){
        self.push(hash_);
        return self.length-1;
    }

    function addNoDup(bytes32[] storage self, bytes32 hash_) external returns(bool,uint256) {
        for(uint256 i=0;i<self.length;i++){
            if(self[i] == hash_){
                return (false,0);
            }
        }
        self.push(hash_);
        return (true,self.length-1);
    }

    function del(bytes32[] storage self, bytes32 hash_) external returns(bool){
        for(uint256 i=0;i<self.length;i++){
            if(self[i] == hash_){
                self[i] = self[self.length-1];
                self.pop();
                return true;
            }
        }
        return false;
    }
}

library LibAddressArray{
    function add(address[] storage self, address addr_) external returns(uint256){
        self.push(addr_);
        return self.length-1;
    }

    function addNoDup(address[] storage self, address addr_) external returns(bool,uint256) {
        for(uint256 i=0;i<self.length;i++){
            if(self[i] == addr_){
                return (false,0);
            }
        }
        self.push(addr_);
        return (true,self.length-1);
    }

    function del(address[] storage self, address addr_) external returns(bool){
        for(uint256 i=0;i<self.length;i++){
            if(self[i] == addr_){
                self[i] = self[self.length-1];
                self.pop();
                return true;
            }
        }
        return false;
    }
}