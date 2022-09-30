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

    struct PaymentItem {
        uint8 decimal;
        bool enable;
    }

    address public usdtAddr;
    uint8 public baseDecimal;
    uint256[8] public SLTax;
    uint256[8] public TLPrice;
    uint constant lengthCount = 8;
    mapping(bytes32=>uint256[8]) public SLPrice;
    mapping(address=>PaymentItem) public paymentMap;
    address[] public paymentList;
    mapping(address=>address) public chainLinkRateAddr;
    IDnsTopLevelName public dnsTop;

    constructor(address top_){
        //rinkeby initialize
        //top price
        for(uint i=0;i<TLPrice.length;i++){
            TLPrice[i] = 3000*(10**6);
        }
        //tax
        for(uint i=0;i<SLTax.length;i++){
            SLTax[i] = 5*(10**6);
        }
        //usdt
        usdtAddr = address(0x4F4f07917E13785bfd55cd3485B254BdE6f09265);
        baseDecimal = 6;
        //payment
//        paymentMap[0x4F4f07917E13785bfd55cd3485B254BdE6f09265] = PaymentItem(6,true);
        paymentMap[0x0000000000000000000000000000000000000000] = PaymentItem(18,true);
        paymentList.addNoDup(0x0000000000000000000000000000000000000000);
        paymentMap[0x571787b5E033bF8bb2A1e5930265828ef3fFca00] = PaymentItem(18,true);
        paymentList.addNoDup(0x571787b5E033bF8bb2A1e5930265828ef3fFca00);
        paymentMap[0x65eEdD3b0B0A7c5cfBe97b22bBF19Cd492f2237E] = PaymentItem(6,true);
        paymentList.addNoDup(0x65eEdD3b0B0A7c5cfBe97b22bBF19Cd492f2237E);
        //eth chainlink
        dnsTop = IDnsTopLevelName(top_);
        chainLinkRateAddr[0x0000000000000000000000000000000000000000] = 0xe0005488e0D9Be7b24ccE51F722D24749391327F;
        chainLinkRateAddr[0x571787b5E033bF8bb2A1e5930265828ef3fFca00] = 0x136a734040F258054dA7ebF42BD25b73e913839a;
        chainLinkRateAddr[0x65eEdD3b0B0A7c5cfBe97b22bBF19Cd492f2237E] = 0x8857AAEa5c663a4408963e8D964D22e438d596F2;
        //mainnet initialize
//        usdtAddr = address(0xdAC17F958D2ee523a2206206994597C13D831ec7);
//        baseDecimal = 6;

//        paymentMap[0x0000000000000000000000000000000000000000] = PaymentItem(18,true);    //eth
//        paymentList.addNoDup(0x0000000000000000000000000000000000000000);
//        paymentMap[0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48] = PaymentItem(18,true);    //usdc
//        paymentList.addNoDup(0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48);
//        paymentMap[0xD31a59c85aE9D8edEFeC411D448f90841571b89c] = PaymentItem(6,true);     //sol
//        paymentList.addNoDup(0xD31a59c85aE9D8edEFeC411D448f90841571b89c);

//        chainLinkRateAddr[0x0000000000000000000000000000000000000000] = 0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419; //eth-usd
//        chainLinkRateAddr[0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48]= 0x8fFfFfd4AfB6115b954Bd326cbe7B4BA576818f6;//usdc-usd
//        chainLinkRateAddr[0xD31a59c85aE9D8edEFeC411D448f90841571b89c]= 0x4ffC43a60e009B551865A93d232E33Fce9f01507;//sol-usd

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

    function addPayment(address erc20Addr_,uint8 decimal_,address chainLinkContact_) external onlyOwner {
        paymentMap[erc20Addr_] = PaymentItem(decimal_,true);
        paymentList.addNoDup(erc20Addr_);
        chainLinkRateAddr[erc20Addr_] = chainLinkContact_;
    }

    function setPayment(address erc20Addr_,uint8 decimal_,bool enable) external onlyOwner{
        require(paymentMap[erc20Addr_].decimal>0,"payment is not exists");
        paymentMap[erc20Addr_] = PaymentItem(decimal_,enable);
    }

//    function delPayment(address erc20Addr_) external onlyOwner{
//        delete paymentMap[erc20Addr_];
//        paymentList.del(erc20Addr_);
//    }

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
        return paymentMap[erc20Addr_].decimal;
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

        return (_base(baseDecimal,paymentMap[erc20Addr_].decimal,cost),roundId);
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
            return (_base(baseDecimal,paymentMap[erc20Addr_].decimal,cost),true);
        }else if (roundId == (lastPriceId_ + 1)){
            (
            ,
            answer,
            ,
            ,
            ) = pr.getRoundData(lastPriceId_);
            return (_base(baseDecimal,paymentMap[erc20Addr_].decimal,cost),true);
        }else{
            return (0,false);
        }
    }

    function getExchangeWithId(address erc20Addr_, uint80 lastPriceId_, uint256 price_) external view returns (uint256,bool) {
        return _getExchangeWithId(erc20Addr_,lastPriceId_,price_);
    }

    function getTopLevelNamePrice(address erc20Addr_,uint80 lastPriceId_, uint8 nameLength_,uint8 year_) external override view returns (uint256,bool){
        uint256 _price = _getTopLevelNamePrice(nameLength_) * year_;
        if(erc20Addr_ == usdtAddr){
            return (_price,true);
        }
        if(paymentMap[erc20Addr_].enable == false){
            return (0,false);
        }
        return _getExchangeWithId(erc20Addr_,lastPriceId_,_price);
    }

    function getSecondLevelNamePrice(bytes32 fatherHash_,address erc20Addr_,uint80 lastPriceId_, uint8 nameLength_, uint8 year_)
    external override view returns(uint256,uint256,bool){
        uint256 _price = _getSecondLevelNamePrice(fatherHash_,nameLength_)*year_;
        uint256 _tax = _getSecondLevelNameTax(nameLength_) * year_;

        if( erc20Addr_ == usdtAddr ){
            return (_tax,_price,true);
        }
//        if(paymentMap[erc20Addr_].enable == false){
//            return (0,0,false);
//        }
        require(paymentMap[erc20Addr_].enable,"payment disabled");

        (uint256 cost,bool valid) = _getExchangeWithId(erc20Addr_,lastPriceId_,_price);
        require(valid,"price not correct");
        (uint256 tax,bool validtax) = _getExchangeWithId(erc20Addr_,lastPriceId_,_tax);
        require(validtax,"tax not correct");
        return (tax,cost,true);

    }

    function priceIdIsValid(address erc20Addr_,uint80 lastPriceId) external override view returns(bool){
        if (erc20Addr_ == usdtAddr){
            return true;
        }

        if(paymentMap[erc20Addr_].enable ==false){
            return false;
        }

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
