//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "../lib/LibDnsArray.sol";


interface Icost{
    function getTopLevelNamePrice(address erc20Addr_,uint80 lastPriceId_, uint8 nameLength_,uint8 year_)
            external view returns (uint256,bool);
    function getSecondLevelNamePrice(bytes32 fatherHash_,address erc20Addr_,uint80 lastPriceId_, uint8 nameLength_, uint8 year_)
            external  view returns(uint256,uint256,bool);
    function priceIdIsValid(address erc20Addr_,uint80 lastPriceId) external view returns(bool);
}