//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "./IDnsErc721Owner.sol";


contract Erc721Owner is IDnsErc721Owner {
    address payable public erc721Owner;
    address payable public  owner;

    constructor(address owner_,address erc721Owner_){
        erc721Owner = payable(erc721Owner_);
        owner = payable(owner_);
    }

    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }

    function transferOwnership(address payable newOwner) external override{
        require(msg.sender == erc721Owner,"not allowed");
        owner = newOwner;
    }

    function transferErc721Owner(address payable erc721Owner_) external {
        require(msg.sender == owner,"not the owner");
        erc721Owner = erc721Owner_;
    }

}