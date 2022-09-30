//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "../../ERC20/ERC20.sol";
import "./AggregatorV3Interface.sol";

//0xaEa429C5aBe96a7305885D445fa73672dEC6CFfD rinkeby
contract FakeSol is ERC20{
    string  public constant  name = "Fake SOL";
    string  public constant  symbol = "FSOL";
    uint8   public constant  decimals = 18;
    uint256 public constant INITIAL_SUPPLY = 4.2e8 * (10 ** uint256(decimals));

    constructor() {
        _mint(msg.sender, INITIAL_SUPPLY);
    }

    function burn(uint amount) external{
        _burn(msg.sender, amount);
    }

    function burnFrom(address account, uint amount) external{
        _burnFrom(account, amount);
    }
}

//0x68175939eC2dA237e981bC6542d00bBD25911503 rinkeby
contract SolUSDT is AggregatorV3Interface {
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
        desc = "sol-usdt";
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