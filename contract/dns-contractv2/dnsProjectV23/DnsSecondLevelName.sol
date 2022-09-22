//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../utils/Owner.sol";
import "../ERC721/ERC721.sol";
import "../ERC20/IERC20.sol";
import "./cost/Icost.sol";
import "./IDnsTopLevelName.sol";
import "./accountant/IDnsAccountant.sol";
import "./dnserc721/IDnsNameErc721.sol";
import "./lib/LibDnsToolKit.sol";
import "./lib/LibDnsSig.sol";

contract DnsSecondLevelName {
    using Strings for uint256;
    Icost public costContractAddr;
    IDnsAccountant public accountantC;
    IDnsTopLevelName public dnsTop;
    uint8 constant SecondMintClose = 1;
    uint8 constant SecondChargeClose  = 2;
    uint8 constant SecondSigMintClose = 4;
    uint8 constant SecondSigChargeClose = 8;
    uint8 public mintSwitch = 0;

    struct NameItem{
        string entireName;
        uint32 expireTime;
        uint256 tokenId;
    }

    mapping(bytes32=>mapping(bytes32=>string)) public subNameStore;
    mapping(bytes32=>mapping(bytes32=>NameItem)) public nameStore;

    event EvMintSecondLevelName(string entireName, uint8 year, address erc20Addr, uint256 tokenId);
    event EvChargeSecondLevelName(bytes32 fatherHash, bytes32 nameHash,
        uint8 year, address erc20Addr, bool isTransfer);
    event EvMintSecondLevelNameBySig(string entireName, uint8 year, address erc20Addr, uint256 nonce, uint256 price, uint256 tokenId);
    event EvChargeSecondLevelNameBySig(bytes32 fatherHash, bytes32 nameHash, uint8 year, address erc20Addr,
        uint256 nonce, uint256 price, bool isTransfer);
    event EvAddSubName(string entireSubName, bytes32 level2FatherHash);
    event EvDelSubName(bytes32 nameHash, bytes32 level2FatherHash, bytes32 topHash);

    constructor(address cost_,address acct_, address dnsTop_){
        costContractAddr = Icost(cost_);
        accountantC = IDnsAccountant(acct_);
        dnsTop = IDnsTopLevelName(dnsTop_);
    }

    modifier onlyOwner{
        require(msg.sender == owned(address(dnsTop)).owner(),"not owner");
        _;
    }

    function setContract(address cost_,address acct_, address dnsTop_,uint8 mintSw_) external onlyOwner{
        costContractAddr = Icost(cost_);
        accountantC = IDnsAccountant(acct_);
        dnsTop = IDnsTopLevelName(dnsTop_);
        mintSwitch = mintSw_;
    }

    function _secondLevelName(bytes32 fHash_,address erc20Addr_,uint256 tax_,uint256 cost_) internal{
        uint256 max = cost_;
        if (tax_ > cost_){
            max = tax_;
        }
        uint256 left = max - tax_;
        if (erc20Addr_ == address (0)){
            require(msg.value >= max,"payout is not enough");
            accountantC.deposit(dnsTop.getErc721Contract(bytes32(0)),erc20Addr_,tax_);
            if(msg.value > max){
                left = msg.value - tax_;
            }
            accountantC.deposit(dnsTop.getErc721Contract(fHash_),erc20Addr_,tax_);
            payable(address(accountantC)).transfer(msg.value);
        }else{
            require(IERC20(erc20Addr_).balanceOf(msg.sender)>=max &&
                IERC20(erc20Addr_).allowance(msg.sender,address(this))>= max,"payout is not enough");
            accountantC.deposit(dnsTop.getErc721Contract(bytes32(0)),erc20Addr_,tax_);
            if(left>0){
                accountantC.deposit(dnsTop.getErc721Contract(fHash_),erc20Addr_,left);
            }
            IERC20(erc20Addr_).transferFrom(msg.sender,address(accountantC),max);
        }
    }


    function _extendExpire(bytes32 fHash_,bytes32 nameHash_, uint8 year_, bool isTransfer_) internal{
        if(block.timestamp > nameStore[fHash_][nameHash_].expireTime){
            if(isTransfer_){
                IDnsNameErc721(dnsTop.getErc721Contract(fHash_)).DnsTransfer(
                    msg.sender,
                    nameStore[fHash_][nameHash_].tokenId);
            }
            nameStore[fHash_][nameHash_].expireTime = uint32(block.timestamp )+ uint32(year_)*uint32(365 days);
        }
        nameStore[fHash_][nameHash_].expireTime += uint32(year_) * uint32(365 days);
        IDnsNameErc721(dnsTop.getErc721Contract(fHash_)).DnsExtendExpire(nameStore[fHash_][nameHash_].tokenId,
            nameStore[fHash_][nameHash_].expireTime);
    }


    function MintSecondLevelName(string memory entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId) external payable{
        require(mintSwitch&SecondMintClose == 0,"can't mint");
        bytes32 fhash = LibDnsToolKit.fatherNameHash(bytes(entireName_));
        require(dnsTop.getErc721Contract(fhash)!=address(0),"not open to reg");
        bytes memory sName = LibDnsToolKit.getSecondLevelName(bytes(entireName_));
        (uint256 tax,uint256 cost,) = costContractAddr.getSecondLevelNamePrice(fhash,erc20Addr_,
            lastPriceId,
            uint8(sName.length),year_);
        _secondLevelName(fhash,erc20Addr_,tax,cost);
        bytes32 nameHash = keccak256(bytes(entireName_));
        nameStore[fhash][nameHash] = NameItem(entireName_,uint32(block.timestamp )+ uint32(year_)*uint32(365 days),
            IDnsNameErc721(dnsTop.getErc721Contract(fhash)).MintId(nameHash,
            uint32(block.timestamp )+ uint32(year_)*uint32(365 days),msg.sender));
        emit EvMintSecondLevelName( entireName_, year_, erc20Addr_, nameStore[fhash][nameHash].tokenId);
    }

    function ChargeSecondLevelName(bytes32 fatherHash_,
        bytes32 nameHash_,
        uint8 year_,
        address erc20Addr_,
        uint80 lastPriceId,
        bool isTransfer_) external payable{
        require(mintSwitch&SecondChargeClose == 0,"can't charge");
        require(nameStore[fatherHash_][nameHash_].expireTime>0,"name not exist");
        bytes memory sName = LibDnsToolKit.getSecondLevelName(bytes(nameStore[fatherHash_][nameHash_].entireName));
        (uint256 tax,uint256 cost,) = costContractAddr.getSecondLevelNamePrice(fatherHash_,erc20Addr_,
            lastPriceId,
            uint8(sName.length),year_);
        _secondLevelName(fatherHash_,erc20Addr_,tax,cost);
        _extendExpire(fatherHash_,nameHash_,year_,isTransfer_);
        emit EvChargeSecondLevelName( fatherHash_, nameHash_, year_,  erc20Addr_, isTransfer_);
    }

    function MintSecondLevelNameBySig(string memory entireName_, uint8 year_, address erc20Addr_, uint80 lastPriceId,
        uint256 nonce, uint256 price_, bytes memory sig) external payable{
        require(mintSwitch&SecondSigMintClose == 0,"can't mint");
        bytes32 fhash = LibDnsToolKit.fatherNameHash(bytes(entireName_));
        bytes32 nameHash = keccak256(bytes(entireName_));
        require(nameStore[fhash][nameHash].expireTime==0,"name was registered");
        require(dnsTop.getErc721Contract(fhash)!=address(0),"not open to reg");
        require(IDnsNameErc721(dnsTop.getErc721Contract(fhash)).SigUserAddr() ==
            LibDnsSignature.SigUserAddr(keccak256(abi.encodePacked(entireName_,
                year_,
                erc20Addr_,
                lastPriceId,
                nonce,
                price_,
                msg.sender)),sig),"sig not match");
        bytes memory sName = LibDnsToolKit.getSecondLevelName(bytes(entireName_));

        if(price_ > 0){
            (uint256 tax,,) = costContractAddr.getSecondLevelNamePrice(fhash,erc20Addr_,
                lastPriceId,
                uint8(sName.length),uint8(year_));
            if (price_ < tax){
                tax = price_;
            }
            _secondLevelName(fhash,erc20Addr_,tax,price_);
        }else{
            require(block.timestamp > uint32(lastPriceId) && block.timestamp-uint32(lastPriceId)<(1 days),"sig expired");
        }

        nameStore[fhash][nameHash] = NameItem(entireName_,uint32(block.timestamp )+ uint32(year_)*uint32(365 days),
            IDnsNameErc721(dnsTop.getErc721Contract(fhash)).MintId(nameHash,
            uint32(block.timestamp )+ uint32(year_)*uint32(365 days),msg.sender));
        emit EvMintSecondLevelNameBySig( entireName_,  year_,  erc20Addr_, nonce,  price_, nameStore[fhash][nameHash].tokenId);
    }

    function ChargeSecondLevelNameBySig(bytes32 fatherHash_,
        bytes32 nameHash_,
        uint8 year_,
        address erc20Addr_,
        uint80 lastPriceId,
        uint256 nonce, uint256 price_, bytes memory sig,
        bool isTransfer_) external payable{
        require(mintSwitch&SecondSigChargeClose == 0,"can't charge");
        require(nameStore[fatherHash_][nameHash_].expireTime>0,"name not exist");
        require(IDnsNameErc721(dnsTop.getErc721Contract(fatherHash_)).SigUserAddr() ==
            LibDnsSignature.SigUserAddr(keccak256(abi.encodePacked(fatherHash_,
                year_,
                erc20Addr_,
                lastPriceId,
                nonce,
                price_,
                msg.sender)),sig),"sig not match");
        bytes memory sName = LibDnsToolKit.getSecondLevelName(bytes(nameStore[fatherHash_][nameHash_].entireName));

        if(price_ > 0){
            (uint256 tax,,) = costContractAddr.getSecondLevelNamePrice(fatherHash_,erc20Addr_,
                lastPriceId,
                uint8(sName.length),uint8(year_));
            if (price_ < tax){
                tax = price_;
            }
            _secondLevelName(fatherHash_,erc20Addr_,tax,price_);
        }else{
            require(block.timestamp > uint32(lastPriceId) && block.timestamp-uint32(lastPriceId)<(1 days),"sig expired");
        }
        _extendExpire(fatherHash_,nameHash_,year_,isTransfer_);
        emit EvChargeSecondLevelNameBySig( fatherHash_,  nameHash_, year_, erc20Addr_,
            nonce, price_, isTransfer_);
    }

    function AddSubName(string memory entireSubName_, bytes32 level2FatherHash_) external{
        require(LibDnsToolKit.getLeve2FatherNameHash(bytes(entireSubName_))==level2FatherHash_,"father name not correct");
        bytes32 fatherHash = LibDnsToolKit.getTopFatherNameHash(bytes(entireSubName_));
        require(msg.sender == IERC721(dnsTop.getErc721Contract(fatherHash)).
        ownerOf(nameStore[fatherHash][level2FatherHash_].tokenId),"not owner");
        subNameStore[level2FatherHash_][LibDnsToolKit.entireNameHash(entireSubName_)] = entireSubName_;
        emit EvAddSubName( entireSubName_, level2FatherHash_);
    }

    function DelSubName(bytes32 nameHash_, bytes32 level2FatherHash_, bytes32 topHash_) external{
        require(msg.sender == IERC721(dnsTop.getErc721Contract(topHash_)).
        ownerOf(nameStore[topHash_][level2FatherHash_].tokenId),"not owner");

        delete subNameStore[level2FatherHash_][nameHash_];

        emit EvDelSubName(nameHash_, level2FatherHash_, topHash_);
    }

}