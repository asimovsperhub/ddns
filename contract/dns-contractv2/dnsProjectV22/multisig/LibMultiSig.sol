//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../lib/LibDnsArray.sol";


library LibMultiSig{
    using LibBytes32Array for bytes32[];
    using LibAddressArray for address[];

    struct safeSig{
        uint256 maxSig;
        uint256 sigCount;
        mapping(address=>bool) signerSet;
        address[] signerList;
    }

    struct NameMultiSigner{
        bool work;
        bool lock;
        uint256 SignerCount;
        mapping(address=>bool) signerSet;
        address[] nameSignerList;
        mapping(bytes32=>safeSig) task;
        bytes32[] taskList;
    }

    function openSafeSig(NameMultiSigner storage self) public{
        self.work = true;
    }

    function isMultiSigOpen(NameMultiSigner storage self) public view returns(bool){
        return self.work;
    }

    function getSignerCount(NameMultiSigner storage self) public view returns (uint256){
        return self.SignerCount;
    }

    function closeSafeSig(NameMultiSigner storage self) public{
        self.work = false;
    }

    function addSigner(NameMultiSigner storage self,address signer_) public{
        require(!self.signerSet[signer_],"user is in set");
        self.SignerCount ++;
        self.signerSet[signer_] = true;
        self.nameSignerList.add(signer_);
    }

    function removeSigner(NameMultiSigner storage self, address signer_) public{
        require(self.signerSet[signer_],"user is not in set");
        self.SignerCount --;
        delete self.signerSet[signer_];
        self.nameSignerList.del(signer_);
    }


    function newSafeSig(NameMultiSigner storage self,address signer_,bytes32 hash_,uint256 count_) public {
        require(self.SignerCount >= count_,"sig count not correct");
        require(self.signerSet[signer_],"signer not valid");
        require(self.task[hash_].maxSig == 0,"task is existed");
        require(!self.lock,"only one task");

        self.task[hash_].maxSig = count_;
        self.task[hash_].sigCount ++;
        self.task[hash_].signerSet[signer_] = true;
        self.task[hash_].signerList.add(signer_);
        self.lock = true;

        self.taskList.add(hash_);
    }

    function incSafeSig(NameMultiSigner storage self,address signer_,bytes32 hash_) public {
        require(self.task[hash_].maxSig>0,"safe sig is not exists");
        require(self.signerSet[signer_],"signer not valid");
        require(!self.task[hash_].signerSet[signer_],"sig is existed");

        self.task[hash_].sigCount ++;
        self.task[hash_].signerSet[signer_] = true;
        self.task[hash_].signerList.add(signer_);
    }

    function getSafeSig(NameMultiSigner storage self,bytes32 hash_,address candidate_) public view returns(bool){
        require(self.task[hash_].maxSig>0,"safe sig is not exists");
        require(self.task[hash_].signerSet[candidate_] ,"not a candidate");

        if(self.task[hash_].sigCount == self.task[hash_].maxSig){
            return true;
        }

        return false;
    }

    function removeSafeSig(NameMultiSigner storage self,bytes32 hash_,address candidate_) public{
        require(self.task[hash_].maxSig>0 && self.lock,"safe sig is not exists");
        require(self.signerSet[candidate_],"not a candidate, can't delete it");

        self.lock = false;

        delete self.task[hash_];
        self.taskList.del(hash_);
    }

    function isExisted(NameMultiSigner storage self,bytes32 hash_) public view returns(bool){
        if(self.task[hash_].maxSig>0){
            return true;
        }
        return false;
    }

}