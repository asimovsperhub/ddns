//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "../owners/DnsOwners.sol";

library LibDnsOwner{

    function NewDnsOwner(address owner_) public returns(address){
        return address(new DnsOwners(owner_));
    }

}