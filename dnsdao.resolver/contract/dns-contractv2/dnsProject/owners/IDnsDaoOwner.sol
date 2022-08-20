//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./IDnsOwner.sol";

interface IDnsDaoOwner is IDnsOwner{
    function addAllowance(address user_) external;
}


