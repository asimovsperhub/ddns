
//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "../dnserc721/dnsNameErc721.sol";

library LibDnsNameErc721{
    function NewDnsNameErc721(
        string memory name_,
        string memory symbol_,
        address owner_,
        bytes32 hash_) public returns(address){
        return address(new DnsNameErc721(name_,symbol_,owner_,hash_));
    }
}


