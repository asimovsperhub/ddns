//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


interface IMultiSig {
    function getSafeSig(address contractAddr_,bytes32 hash_,address candidate_) external view returns(bool);
}