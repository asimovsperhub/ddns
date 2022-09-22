//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


interface IDnsNameErc721 {
    function MintId(bytes32 hash_,uint32 expireTime_,address idOwner_) external returns(uint256);
    function DnsTransfer(address to, uint256 tokenId) external;
    function DnsExtendExpire(uint256 tokenId_,uint32 expireTime_) external;
//    function IdOwner(uint256 tokenId) external view returns(address);
    function SigUserAddr() external view returns(address);
}



