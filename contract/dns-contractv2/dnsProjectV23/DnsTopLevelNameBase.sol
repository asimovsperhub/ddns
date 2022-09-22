//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../ERC721/ERC721.sol";
import "../ERC20/IERC20.sol";
import "../utils/Owner.sol";
import "./lib/LibDnsToolKit.sol";
import "./lib/LibDnsSig.sol";
import "./cost/Icost.sol";
import "./dnserc721/IDnsErc721Owner.sol";
import "./dnserc721/IDnsNameErc721.sol";
import "../ERC721/utils/Strings.sol";
import "./IDnsTopLevelName.sol";
import "./accountant/IDnsAccountant.sol";
import "./lib/LibDnsNameErc721.sol";

contract DnsTopLevelNameBase is owned,IDnsTopLevelName {
    using Strings for uint256;
    Icost public costContractAddr;
    IDnsAccountant public accountantC;
    address public dnsSecond;
    uint8 constant TopMintClose = 1;
    uint8 constant TopChargeClose  = 2;
    uint8 constant TopSigMintClose = 4;
    uint8 constant TopSigChargeClose = 8;
    uint8 public mintSwitch = 0;

    struct NameItem{
        string entireName;
        uint32 expireTime;
        uint256 tokenId;
    }

    mapping(bytes32=>address) public erc721Store;
    mapping(bytes32=>NameItem) public nameStore;

//    event EvMintTopLevelName(string entireName, uint8 year, address erc20Addr, uint256 tokenId);
//    event EVChargeTopLevelName(bytes32 nameHash,uint8 year, address erc20Addr, bool isTransfer_);
//    event EvOpen2Reg(bytes32 nameHash);
//    event EvMintTopLevelNameBySig(string entireName, uint8 year, address erc20Addr, uint256 nonce, uint256 price,uint256 tokenId);
//    event EvChargeTopLevelNameBySig(bytes32 nameHash,uint8 year, address erc20Addr, uint256 nonce, uint256 price, bool isTransfer);

    constructor(string memory name, string memory symbol){
        erc721Store[bytes32(0)] = LibDnsNameErc721.NewDnsNameErc721(name,symbol,msg.sender,bytes32(0));
    }

    //rewrite transferOwnership
    function transferOwnership(address payable newOwner) external override onlyOwner{
        owner = newOwner;
        IDnsErc721Owner(erc721Store[bytes32(0)]).transferOwnership(newOwner);
    }

    function getErc721Contract(bytes32 fatherHash) external override view returns(address){
        require(erc721Store[fatherHash] != address(0),"erc721 not found");
        return erc721Store[fatherHash];
    }
    function getOperator() external override view returns(address){
        return dnsSecond;
    }

    function setContract(address costAddr_,address accountantAddr_,address dnsSecond_,uint8 mintSw_) external onlyOwner{
        costContractAddr = Icost(costAddr_);
        accountantC = IDnsAccountant(accountantAddr_);
        dnsSecond = dnsSecond_;
        mintSwitch = mintSw_;
    }



//    function _topLevelNamePrice(address erc20Addr_,uint256 cost_) internal {
//        require(cost_>0,"not a valid cost");
//        if (erc20Addr_ == address(0)){
//            require(msg.value>=cost_,"payout is not enough");
//            accountantC.deposit(erc721Store[bytes32(0)],erc20Addr_,msg.value);
//            payable(address(accountantC)).transfer(msg.value);
//        }else{
//            require(IERC20(erc20Addr_).balanceOf(msg.sender)>=cost_ &&
//                IERC20(erc20Addr_).allowance(msg.sender,address(this))>= cost_,"payout is not enough");
//            accountantC.deposit(erc721Store[bytes32(0)],erc20Addr_,cost_);
//            IERC20(erc20Addr_).transferFrom(msg.sender,address(accountantC),cost_);
//        }
//    }
//
//    function _secondLevelName(bytes32 fHash_,address erc20Addr_,uint256 tax_,uint256 cost_) internal{
//        uint256 max = cost_;
//        if (tax_ > cost_){
//            max = tax_;
//        }
//        uint256 left = max - tax_;
//        if (erc20Addr_ == address (0)){
//            require(msg.value >= max,"payout is not enough");
//            accountantC.deposit(erc721Store[bytes32(0)],erc20Addr_,tax_);
//            if(msg.value > max){
//                left = msg.value - tax_;
//            }
//            accountantC.deposit(erc721Store[fHash_],erc20Addr_,tax_);
//            payable(address(accountantC)).transfer(msg.value);
//        }else{
//            require(IERC20(erc20Addr_).balanceOf(msg.sender)>=max &&
//                IERC20(erc20Addr_).allowance(msg.sender,address(this))>= max,"payout is not enough");
//            accountantC.deposit(erc721Store[bytes32(0)],erc20Addr_,tax_);
//            if(left>0){
//                accountantC.deposit(erc721Store[fHash_],erc20Addr_,left);
//            }
//            IERC20(erc20Addr_).transferFrom(msg.sender,address(accountantC),max);
//        }
//    }
//
//    function MintTopLevelName(string memory entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId)external payable{
//        require(mintSwitch&TopMintClose == 0,"can't mint");
//        require(LibDnsToolKit.verifyRoot(bytes(entireName_)),"not a root Name");
//        bytes32 nameHash = LibDnsToolKit.entireNameHash(entireName_);
//        require(nameStore[nameHash].tokenId == 0,"name was registered");
//        (uint256 cost, bool valid) = costContractAddr.getTopLevelNamePrice(erc20Addr_,lastPriceId,uint8(bytes(entireName_).length),year_);
//        require(valid,"price not valid");
//        _topLevelNamePrice(erc20Addr_,cost);
//        nameStore[nameHash] = NameItem(entireName_,uint32(block.timestamp)+(uint32(year_)*uint32((365 days))),
//            IDnsNameErc721(erc721Store[bytes32(0)]).MintId(nameHash,
//            uint32(block.timestamp)+(uint32(year_)*uint32(365 days)),msg.sender));
//        emit EvMintTopLevelName(entireName_, year_, erc20Addr_,  nameStore[nameHash].tokenId);
//    }
//
//    function _extendExpire(bytes32 nameHash_, uint8 year_, bool isTransfer_) internal{
//        if(block.timestamp > nameStore[nameHash_].expireTime){
//            if(isTransfer_){
//                IDnsNameErc721(erc721Store[bytes32(0)]).DnsTransfer(
//                    msg.sender,
//                    nameStore[nameHash_].tokenId);
//            }
//            nameStore[nameHash_].expireTime = uint32(block.timestamp);
//        }
//        nameStore[nameHash_].expireTime += year_ * 365 days;
//        IDnsNameErc721(erc721Store[bytes32(0)]).DnsExtendExpire(nameStore[nameHash_].tokenId,
//            nameStore[nameHash_].expireTime);
//    }
//
//    function ChargeTopLevelName(bytes32 nameHash_,uint8 year_, address erc20Addr_, uint80 lastPriceId,bool isTransfer_) external payable{
//        require(mintSwitch&TopChargeClose == 0,"can't charge");
//        require(nameStore[nameHash_].expireTime>0,"name not exist");
//        (uint256 cost, bool valid) = costContractAddr.getTopLevelNamePrice(erc20Addr_,
//            lastPriceId,
//            uint8(bytes(nameStore[nameHash_].entireName).length),year_);
//        require(valid,"price not valid");
//        _topLevelNamePrice(erc20Addr_,cost);
//        _extendExpire(nameHash_,year_,isTransfer_);
//        emit EVChargeTopLevelName(nameHash_,year_, erc20Addr_, isTransfer_);
//    }
//
//    function Open2Reg(bytes32 nameHash_) external {
//        require(erc721Store[nameHash_]==address(0),"opened");
//        require(IERC721(erc721Store[bytes32(0)]).ownerOf(nameStore[nameHash_].tokenId)==msg.sender,"not owner");
//        erc721Store[nameHash_] = LibDnsNameErc721.NewDnsNameErc721(
//            nameStore[nameHash_].entireName,
//            string(abi.encodePacked("dnsDao#",nameStore[nameHash_].tokenId.toString())),
//            msg.sender,nameHash_);
//        emit EvOpen2Reg(nameHash_);
//    }
//
//    function MintTopLevelNameBySig(string memory entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId,
//        uint256 nonce, uint256 price_, bytes memory sig) external payable{
//        require(mintSwitch&TopSigMintClose == 0,"can't mint");
//        require(LibDnsToolKit.verifyRoot(bytes(entireName_)),"not a root Name");
//        bytes32 nameHash = LibDnsToolKit.entireNameHash(entireName_);
//        require(nameStore[nameHash].expireTime == 0,"name was registered");
//        require(IDnsNameErc721(erc721Store[bytes32(0)]).SigUserAddr() ==
//            LibDnsSignature.SigUserAddr(keccak256(abi.encodePacked(entireName_,
//                year_,
//                erc20Addr_,
//                lastPriceId,
//                nonce,
//                price_,
//                msg.sender)),sig),"sig not match");
//        if(price_>0){
//            bool valid = costContractAddr.priceIdIsValid(erc20Addr_,lastPriceId);
//            require(valid,"price not valid");
//            _topLevelNamePrice(erc20Addr_,price_);
//        }else{
//            require(block.timestamp > uint32(lastPriceId) && block.timestamp-uint32(lastPriceId)<(1 days),"sig expired");
//        }
//        nameStore[nameHash] = NameItem(entireName_,uint32(block.timestamp)+(uint32(year_)*uint32(365 days)),
//            IDnsNameErc721(erc721Store[bytes32(0)]).MintId(nameHash,
//            uint32(block.timestamp)+(uint32(year_)*uint32(365 days)),msg.sender));
//        emit EvMintTopLevelNameBySig( entireName_, year_, erc20Addr_, nonce, price_,nameStore[nameHash].tokenId);
//    }
//
//    function ChargeTopLevelNameBySig(bytes32 nameHash_,uint8 year_, address erc20Addr_, uint80 lastPriceId,
//        uint256 nonce,  uint256 price_, bytes memory sig,
//        bool isTransfer_) external payable{
//        require(mintSwitch&TopSigChargeClose == 0,"can't charge");
//        require(nameStore[nameHash_].expireTime>0,"name not exist");
//        require(IDnsNameErc721(erc721Store[bytes32(0)]).SigUserAddr() ==
//            LibDnsSignature.SigUserAddr(keccak256(abi.encodePacked(nameHash_,
//                year_,
//                erc20Addr_,
//                lastPriceId,
//                nonce,
//                price_,
//                msg.sender)),sig),"sig not match");
//        if(price_>0){
//            bool valid = costContractAddr.priceIdIsValid(erc20Addr_,lastPriceId);
//            require(valid,"price not valid");
//            _topLevelNamePrice(erc20Addr_,price_);
//        }else{
//            require(block.timestamp > uint32(lastPriceId) && block.timestamp-uint32(lastPriceId)<(1 days),"sig expired");
//        }
//        _extendExpire(nameHash_,year_,isTransfer_);
//
//        emit EvChargeTopLevelNameBySig(nameHash_, year_, erc20Addr_, nonce, price_, isTransfer_);
//    }

}