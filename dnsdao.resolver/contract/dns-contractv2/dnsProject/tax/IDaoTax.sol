//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "./IDnsTax.sol";

interface IDaoTax is IDnsTax {
    function AddTax(address erc20Addr_, uint8 taxType_, uint256 defaultTax_, uint256 percent_,uint256 percentBase_) external;
    function UpdateTax(address erc20Addr_,uint8 taxType_, uint256 defaultTax_, uint256 percent_,uint256 percentBase_) external;
    function AddTaxLen(address erc20Addr_,uint256 len_,uint256 tax_,uint256 percent_) external;
    function DelTaxLen(address erc20Addr_,uint256 len_) external;
}