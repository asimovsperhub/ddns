//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./IMultiSig.sol";
import "./LibMultiSig.sol";
import "../lib/LibDnsArray.sol";
import "../dnserc721/DnsErc721Owner.sol";


contract MultiSig is IMultiSig{
    using LibMultiSig for LibMultiSig.NameMultiSigner;
    using LibAddressArray for address[];
    mapping(address=>LibMultiSig.NameMultiSigner) public multiSignerStore;
    mapping(address=>address[]) public contractList;
    uint256 public defaultCount;

    constructor() {
        defaultCount = 3;
    }

    modifier sigMatched(address contractAddr_,bytes32 hash_){
        require(contractAddr_!=address(0),"contract address not valid");
        if(multiSignerStore[contractAddr_].work == false){
            require(msg.sender== Erc721Owner(contractAddr_).owner(),"not owner");
            _;
        }else{
            require(multiSignerStore[contractAddr_].signerSet[msg.sender],"signer not valid");
            if(!multiSignerStore[contractAddr_].isExisted(hash_) ){
                require(!multiSignerStore[contractAddr_].lock,"lock not free");
                multiSignerStore[contractAddr_].newSafeSig(msg.sender,hash_,defaultCount);
            }else{
                multiSignerStore[contractAddr_].incSafeSig(msg.sender,hash_);
                if(multiSignerStore[contractAddr_].getSafeSig(hash_,msg.sender)){
                    _;
                    multiSignerStore[contractAddr_].removeSafeSig(hash_,msg.sender);
                }
            }
        }
    }

    modifier onlyMultiSigOwner(address contractAddr_){
        require(multiSignerStore[contractAddr_].work == false, "multi signature is opened");
        require(msg.sender==Erc721Owner(contractAddr_).owner(),"not owner");
        _;
    }

    function addSigner(address contractAddr_,address signer_) public
    sigMatched(contractAddr_,keccak256(abi.encodePacked(contractAddr_,signer_,this.addSigner.selector))){
        multiSignerStore[contractAddr_].addSigner(signer_);
        contractList[signer_].add(contractAddr_);
    }

    function delSigner(address contractAddr_,address signer_) public
    sigMatched(contractAddr_,keccak256(abi.encodePacked(contractAddr_,signer_,this.delSigner.selector))){
        multiSignerStore[contractAddr_].removeSigner(signer_);
        contractList[signer_].del(contractAddr_);
    }

    function openMultiSig(address contractAddr_) external
    onlyMultiSigOwner(contractAddr_){
        require(multiSignerStore[contractAddr_].getSignerCount()>=defaultCount,"singer user must moren defaultCount");
        require(!(multiSignerStore[contractAddr_].isMultiSigOpen()),"multi sig is opened");
        return multiSignerStore[contractAddr_].openSafeSig();
    }

    function closeMultiSig(address contractAddr_) external
    sigMatched(contractAddr_,keccak256(abi.encodePacked(contractAddr_,this.closeMultiSig.selector))){
        if(multiSignerStore[contractAddr_].isMultiSigOpen()){
            multiSignerStore[contractAddr_].closeSafeSig();
        }
    }

    function removeSafeSig(address contractAddr_,bytes32 hash_) external{
        return multiSignerStore[contractAddr_].removeSafeSig(hash_,msg.sender);
    }

    function getSafeSig(address contractAddr_, bytes32 hash_,address candidate_) external  override view returns(bool){
        require(multiSignerStore[contractAddr_].isMultiSigOpen(),"multi sig is opened");

        return multiSignerStore[contractAddr_].getSafeSig(hash_,candidate_);
    }

    function getSignerSetAddress(address contractAddr_) external view returns(bytes memory){
        bytes memory ret;
        for(uint i=0;i<multiSignerStore[contractAddr_].nameSignerList.length;i++){
            ret = abi.encodePacked(ret,multiSignerStore[contractAddr_].nameSignerList[i]);
        }
        return ret;
    }

    function _getTaskSignerSetAddress(address contractAddr_,bytes32 hash_) internal view returns(bytes memory){
        bytes memory ret;
        for(uint i=0;i<multiSignerStore[contractAddr_].task[hash_].signerList.length;i++){
            ret = abi.encodePacked(ret,multiSignerStore[contractAddr_].task[hash_].signerList[i]);
        }
        return ret;
    }

    function getTaskSignerSetAddress(address contractAddr_,bytes32 hash_) external view returns(bytes memory){
        return _getTaskSignerSetAddress(contractAddr_,hash_);
    }

    function getAllTask(address contractAddr_) external view returns(bytes memory){
        bytes memory ret;
        for(uint i=0;i<multiSignerStore[contractAddr_].taskList.length;i++){
            ret = abi.encodePacked(ret,multiSignerStore[contractAddr_].taskList[i]);
        }
        return ret;
    }

    function getTaskInfo(address contractAddr_,bytes32 hash_) external view returns(uint256 max,uint256 cnt, bytes memory){
        return (multiSignerStore[contractAddr_].task[hash_].maxSig,
        multiSignerStore[contractAddr_].task[hash_].sigCount,
        _getTaskSignerSetAddress(contractAddr_,hash_));
    }

    function getContractList(address signer_,uint256 idx_) external view returns(uint256,uint256, address[8] memory){
        address[8] memory r;
        uint256 cnt = 0;
        if(idx_>=contractList[signer_].length){
            return (contractList[signer_].length,0,r);
        }

        for(uint256 i=idx_;i<contractList[signer_].length&&i<(idx_+8);i++){
            r[i] = contractList[signer_][i];
            cnt ++;
        }
        return (contractList[signer_].length,cnt,r);
    }

}