//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "../DnsSubName.sol";

library LibDnsName{
    function NewDnsSubName(
        bytes32 fatherHash_,
        address taxC_,
        address priceC_,
        address accountC_,
        address nameOwner_,
        address ownerC_) public returns(address){
        return address(new DnsSubName(fatherHash_,taxC_,priceC_,accountC_,nameOwner_,ownerC_));
    }
}



