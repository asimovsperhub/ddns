//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "../lib/LibDnsArray.sol";
import "./Icost.sol";
import "../../utils/Owner.sol";
import "../IDnsTopLevelName.sol";
import "../dnserc721/DnsErc721Owner.sol";
import "../chainlink/AggregatorV3Interface.sol";

contract DnsCost is Icost{
    using LibAddressArray for address[];

    address public usdtAddr;
    uint8 public baseDecimal;
    uint256[8] public SLTax;
    uint256[8] public TLPrice;
    uint constant lengthCount = 8;
    mapping(bytes32=>uint256[8]) public SLPrice;
    mapping(address=>uint8) public paymentMap;
    address[] public paymentList;
    mapping(address=>address) public chainLinkRateAddr;
    IDnsTopLevelName dnsTop;

    constructor(address top_){
        //rinkeby initialize
        for(uint i=0;i<TLPrice.length;i++){
            TLPrice[i] = 3000*(10**18);
        }

        for(uint i=0;i<SLTax.length;i++){
            SLTax[i] = 5*(10**18);
        }
        usdtAddr = 0x6043bfe64a866c7fb17D1855fe3eBC4342Ca9c15;
        baseDecimal = 18;

        paymentMap[0x0000000000000000000000000000000000000000] = 18;
        paymentList.addNoDup(0x0000000000000000000000000000000000000000);

        //eth chainlink
        dnsTop = IDnsTopLevelName(top_);
        chainLinkRateAddr[0x0000000000000000000000000000000000000000] = 0x8A753747A1Fa494EC906cE90E9f37563A8AF630e;
        //mainnet initialize
    }

    modifier onlyOwner{
        require(msg.sender == owned(address(dnsTop)).owner(),"not allowed");
        _;
    }

    function setRelation(address dnstop_) external onlyOwner{
        dnsTop = IDnsTopLevelName(dnstop_);
    }

    function setUsdtAddr(address usdtTokenAddr_,uint8 baseDecimal_) external onlyOwner{
        usdtAddr = usdtTokenAddr_;
        baseDecimal = baseDecimal_;
    }

    function SetSecondLevelNameTax(uint256[8] calldata tax_) external onlyOwner{
        for(uint i=0;i<tax_.length;i++){
            SLTax[i] = tax_[i];
        }
    }
    function SetTopLevelNamePrice(uint256[8] calldata price_) external onlyOwner{
        for(uint i=0;i<price_.length;i++){
            TLPrice[i] = price_[i];
        }
    }

    function setSecondLevelNamePrice(bytes32 ownerNameHash_,uint256[8] calldata prices_) external {
        require(ownerNameHash_ != bytes32(0) && Erc721Owner(dnsTop.getErc721Contract(ownerNameHash_)).owner() == msg.sender,"not allowed");
        for(uint i=0;i<prices_.length;i++){
            SLPrice[ownerNameHash_][i] = prices_[i];
        }
    }

    function addPayment(address erc20Addr_,uint8 decimal_) external onlyOwner {
        paymentMap[erc20Addr_] = decimal_;
        paymentList.addNoDup(erc20Addr_);
    }

    function delPayment(address erc20Addr_) external onlyOwner{
        delete paymentMap[erc20Addr_];
        paymentList.del(erc20Addr_);
    }
    function addChainLinkAddr(address erc20Addr_, address chainLinkContact_) external onlyOwner{
        chainLinkRateAddr[erc20Addr_] = chainLinkContact_;
    }

    function getAllSecondLevelNameTax() external view returns (uint256[8] memory){
        return SLTax;
    }
    function getAllTopLevelNamePrice() external view returns (uint256[8] memory){
        return TLPrice;
    }
    function getAllSecondLevelNamePrice(bytes32 ownerNameHash_) external view returns(uint256[8] memory){
        return SLPrice[ownerNameHash_];
    }

    function _getSecondLevelNameTax(uint8 nameLength_) internal view returns(uint256){
        if(nameLength_ >= SLTax.length){
            nameLength_ = uint8(SLTax.length-1);
        }
        return SLTax[nameLength_];
    }

    function _getSecondLevelNamePrice(bytes32 ownerNameHash_,uint256 nameLength_) internal view returns(uint256){
        if(nameLength_ >= SLPrice[ownerNameHash_].length){
            nameLength_ = SLPrice[ownerNameHash_].length -1;
        }

        return SLPrice[ownerNameHash_][nameLength_];
    }

    function _getTopLevelNamePrice(uint8 nameLength_) internal view returns(uint256){
        if(nameLength_ >= TLPrice.length){
            nameLength_ = uint8(TLPrice.length) -1;
        }
        return TLPrice[nameLength_];
    }

    function paymentIsExist(address erc20Addr_) external view returns(uint8){
        return paymentMap[erc20Addr_];
    }

    function getPaymentsCount() external view returns(uint256){
        return paymentList.length;
    }

    function getAllPayments() external view returns(bytes memory){
        bytes memory r;
        for(uint i=0;i<paymentList.length;i++){
            r = abi.encodePacked(r,paymentList[i]);
        }
        return r;
    }

    function _base(uint8 baseDecimal_, uint8 exDecimal_,uint256 answer_) internal pure returns(uint256){
        if(baseDecimal_ == exDecimal_){
            return answer_;
        }
        if(baseDecimal_ > exDecimal_){
            return answer_/(10**(baseDecimal_-exDecimal_));
        }else{
            return answer_ * (10**(exDecimal_-baseDecimal_));
        }
    }

    function getLatestExchangeValue(address erc20Addr_,uint256 price_) external view returns(uint256,uint80){
        AggregatorV3Interface pr = AggregatorV3Interface(chainLinkRateAddr[erc20Addr_]);
        (
        uint80 roundId,
        int256 answer,
        ,
        ,
        ) = pr.latestRoundData();

        uint256 cost = price_*(10**pr.decimals())/uint256(answer);

        return (_base(baseDecimal,paymentMap[erc20Addr_],cost),roundId);
    }

    function _getExchangeWithId(address erc20Addr_, uint80 lastPriceId_, uint256 price_) internal view returns (uint256,bool) {
        AggregatorV3Interface pr = AggregatorV3Interface(chainLinkRateAddr[erc20Addr_]);
        (
        uint80 roundId,
        int256 answer,
        ,
        ,
        ) = pr.latestRoundData();

        uint256 cost = price_*(10**pr.decimals())/uint256(answer);

        if(roundId == lastPriceId_){
            return (_base(baseDecimal,paymentMap[erc20Addr_],cost),true);
        }else if (roundId == (lastPriceId_ + 1)){
            (
            ,
            answer,
            ,
            ,
            ) = pr.getRoundData(lastPriceId_);
            return (_base(baseDecimal,paymentMap[erc20Addr_],cost),true);
        }else{
            return (0,false);
        }
    }

    function getExchangeWithId(address erc20Addr_, uint80 lastPriceId_, uint256 price_) external view returns (uint256,bool) {
        return _getExchangeWithId(erc20Addr_,lastPriceId_,price_);
    }

    function getTopLevelNamePrice(address erc20Addr_,uint80 lastPriceId_, uint8 nameLength_,uint8 year_) external override view returns (uint256,bool){
        uint256 _price = _getTopLevelNamePrice(nameLength_) * year_;
        if(lastPriceId_==0 && erc20Addr_ == usdtAddr){
            return (_price,true);
        }
        return _getExchangeWithId(erc20Addr_,lastPriceId_,_price);
    }

    function getSecondLevelNamePrice(bytes32 fatherHash_,address erc20Addr_,uint80 lastPriceId_, uint8 nameLength_, uint8 year_)
    external override view returns(uint256,uint256,bool){
        uint256 _price = _getSecondLevelNamePrice(fatherHash_,nameLength_)*year_;
        uint256 _tax = _getSecondLevelNameTax(nameLength_) * year_;

        if(lastPriceId_ == 0 && erc20Addr_ == usdtAddr){
            return (_tax,_price,true);
        }

        (uint256 cost,bool valid) = _getExchangeWithId(erc20Addr_,lastPriceId_,_price);
        require(valid,"price not correct");
        (uint256 tax,bool validtax) = _getExchangeWithId(erc20Addr_,lastPriceId_,_tax);
        require(validtax,"tax not correct");
        return (tax,cost,valid);

    }

    function priceIdIsValid(address erc20Addr_,uint80 lastPriceId) external override view returns(bool){
        AggregatorV3Interface pr = AggregatorV3Interface(chainLinkRateAddr[erc20Addr_]);
        (
        uint80 roundId,
        ,
        ,
        ,
        ) = pr.latestRoundData();
        if(lastPriceId == roundId || lastPriceId + 1 == roundId){
            return true;
        }
        return false;
    }

}
