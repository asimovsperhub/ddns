//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../owners/IDnsOwnerPub.sol";

contract DaoOwner {
    address public owner;

    constructor(address  owner_){
        owner = owner_;
    }

    modifier onlyDaoOwner{
        require(IDnsOwnerPub(owner).getOwnerByHash(bytes32(0)) == msg.sender,"not name owner");
        _;
    }

    modifier operatorIsAllowed{
        require(IDnsOwnerPub(owner).operatorIsAllowance(msg.sender) ,"not a valid address");
        _;
    }

    modifier onlyContractOwner(address contractAddr_){
        require(IDnsOwnerPub(owner).getOwnerByContract(contractAddr_) == msg.sender,"not owner");
        _;
    }

    modifier onlyNameOwner(bytes32 hash_){
        require(IDnsOwnerPub(owner).getOwnerByHash(hash_) == msg.sender,"not name owner");
        _;
    }

    function transferDaoOwner(address newOwner_) external onlyDaoOwner{
        owner = newOwner_;
    }
}