//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../utils/Owner.sol";
import "../dnsProject/lib/LibDnsArray.sol";

interface IDnsNameReserve {
    function IsReservedName(bytes32 nameHash,bytes calldata p) external view returns (bool);
    // function AddReservedName(bytes32 nameHash) external;
    // function AddReservedNameList(bytes32[] calldata nameList) external;
    // function DelReservedName(bytes32 nameHash) external;

}

contract DnsNameReserve is IDnsNameReserve,owned{
    using LibBytes32Array for bytes32[];

    bytes32[] public ReserveNameList;
    mapping(bytes32=>bool) public ReserverHash;

    function _isReservedName(bytes32 nameHash) internal view returns(bool){
        return ReserverHash[nameHash];
    }
    //p is reserved for other white list mode, like merkle tree,
    function IsReservedName(bytes32 nameHash,bytes calldata p) external override view returns (bool){
        return _isReservedName(nameHash);
    }

    function _addReservedName(bytes32 nameHash) internal {
        ReserveNameList.add(nameHash);
        ReserverHash[nameHash] = true;
    }

    function AddReservedName(bytes32 nameHash) external onlyOwner{
        require(!_isReservedName(nameHash),"ReservedName existsed");
        _addReservedName(nameHash);
    }
    function AddReservedNameList(bytes32[] calldata nameList) external onlyOwner{
        for(uint i=0;i<nameList.length;i++){
            ReserveNameList.addNoDup(nameList[i]);
            ReserverHash[nameList[i]] = true;
        }
    }

//    function AddReservedNameArray(bytes[] calldata nameArr) external onlyOwner{
//        for(uint i=0;i<nameArr.length;i++){
//            bytes32 h = keccak256(nameArr[i]);
//            ReserveNameList.addNoDup(h);
//            ReserverHash[h] = true;
//        }
//    }

    function DelReservedName(bytes32 nameHash) external onlyOwner{
        require(_isReservedName(nameHash),"ReservedName not existsed");
        ReserveNameList.del(nameHash);
        delete ReserverHash[nameHash];
    }

}