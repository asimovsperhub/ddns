//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./AggregatorV3Interface.sol";

//0x884f7B3CF833626194881D22e14Fd9e2A4Eb59F4  rinkeby
contract EthUsdt is AggregatorV3Interface {
    struct RoundInfo{
        uint80 roundId;
        int256 answer;
        uint256 startedAt;
        uint256 updatedAt;
        uint80 answeredInRound;
    }

    uint8 private decimal;
    string private desc;
    uint256 private ver;

    constructor(){
        decimal = 8;
        desc = "eth-usdt";
        ver = 1;
    }

    mapping(uint80=>RoundInfo) public roundInfoMap;
    uint80 public latestRoundId;

    function decimals() external override view returns (uint8){
        return decimal;
    }

    function description() external override view returns (string memory){
        return desc;
    }

    function version() external override view returns (uint256){
        return ver;
    }

    function _getRoundData(uint80 _roundId)internal
    view
    returns (
        uint80 roundId,
        int256 answer,
        uint256 startedAt,
        uint256 updatedAt,
        uint80 answeredInRound
    ){
        return (roundInfoMap[_roundId].roundId,
        roundInfoMap[_roundId].answer,
        roundInfoMap[_roundId].startedAt,
        roundInfoMap[_roundId].updatedAt,
        roundInfoMap[_roundId].answeredInRound);
    }

    function getRoundData(uint80 _roundId)
    external override
    view
    returns (
        uint80 roundId,
        int256 answer,
        uint256 startedAt,
        uint256 updatedAt,
        uint80 answeredInRound
    ){
        return _getRoundData(_roundId);
    }

    function latestRoundData()
    external override
    view
    returns (
        uint80 roundId,
        int256 answer,
        uint256 startedAt,
        uint256 updatedAt,
        uint80 answeredInRound
    ){
        return _getRoundData(latestRoundId);
    }

    function insertIntoRoundInfo(uint80 roundId,int256 answer,
        uint256 startedAt,
        uint256 updatedAt,
        uint80 answeredInRound) external {
        if (roundId == latestRoundId){
            return;
        }
        if (roundId>latestRoundId) {
            latestRoundId = roundId;
        }
        roundInfoMap[latestRoundId] = RoundInfo(roundId,answer,startedAt,updatedAt,answeredInRound);
    }

}