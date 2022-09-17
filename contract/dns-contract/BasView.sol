//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.5.0;

import "./BasToken.sol";
import "./BasOwnership.sol";
import "./BasDomain.sol";
import "./BasOANN.sol";
import "./BasMail.sol";
import "./BasMarket.sol";
import "./BasLib.sol";


contract BasView is ManagedByDAO{
    
    using SafeMath for uint256;
    
    BasRelations public rel;
    
    function getOANNParams() external view returns(
            uint256 MAX_YEAR,
            uint256 AROOT_GAS,
            uint256 BROOT_GAS,
            uint256 SUB_GAS,
            uint256 CUSTOM_PRICE_GAS){
            BasOANN oann = BasOANN(rel.oann());
            MAX_YEAR = oann.MAX_YEAR();
            AROOT_GAS = oann.AROOT_GAS();
            BROOT_GAS = oann.BROOT_GAS();
            SUB_GAS = oann.SUB_GAS();
            CUSTOM_PRICE_GAS = oann.CUSTOM_PRICE_GAS();
    }
    
    mapping(uint8 => string) public ErrorCode;
    
    //set initial error code
    constructor(address rel_addr) {
        
        ErrorCode[0]  = "no error";
        
        //this covers domain registry
        ErrorCode[1]  = "invalid string";
        ErrorCode[2]  = "domain is taken";
        ErrorCode[3]  = "invalid expiration";
        ErrorCode[4]  = "custom price below default";
        ErrorCode[5]  = "root not open";
        
        rel = BasRelations(rel_addr);
    }
    
    function changeRelation(address new_rel) 
                        external
                        OnlyDAO{
        rel = BasRelations(new_rel);                        
    }
    
    function setErrorCode(uint8 index, 
                        string calldata reason) 
                        external
                        OnlyDAO{
        ErrorCode[index] = reason;
    }
    
    function domainIsWild(bytes32 hash)
                    public
                    view
                    returns(bool){
        (, uint expire) = BasTradableOwnership(rel.tro()).ownerOfWithExpire(hash);
        return  expire < block.timestamp;
    }
    
    // this function checks if root domain can be registed, return detail and error code if necessary
    function checkRootRegistry(bytes calldata name,
                            bool isCustom,
                            uint256 cusPrice,
                            uint8 durationInYear)
                            external
                            view
                            returns(uint8,      // errorCode
                                    bool,       // isRare
                                    uint256){   // cost
            BasRootDomain root = BasRootDomain(rel.root());
            BasOANN oann = BasOANN(rel.oann());
            (bool isValid, bool isRare) = root.classifyRoot(name);
            uint256 cost;
            if(!isValid){
                return (1, isRare, cost);
            }
            if(durationInYear > 5){
                return (3, isRare, cost);
            }
            if(cusPrice < oann.SUB_GAS()){
                return (4, isRare, cost);
            }
            if(!domainIsWild(BasHash.Hash(name))){
                return (2, isRare, cost);
            }
            if (isRare) {
                cost = oann.AROOT_GAS().mul(durationInYear);
            } else {
                cost = oann.BROOT_GAS().mul(durationInYear);
            }
            if (isCustom){
                cost = cost.add(oann.CUSTOM_PRICE_GAS());
            }
            return (0, isRare, cost);
    }
    
    // this function checks if sub domain can be registed, return detail and error code if necessary
    function checkSubRegistry(bytes calldata rName,
                            bytes calldata sName,
                            uint256 durationInYear)
                            external
                            view
                            returns(uint8,      // errorCode
                                    uint256){   // cost

            uint256 cost;
            if(!BasSubDomain(rel.sub()).verifySub(sName)){
                return (1, cost);
            }
            if(durationInYear > 5){
                return (3, cost);
            }
            if(!domainIsWild(BasHash.Hash(abi.encodePacked(sName, ".", rName)))){
                return (2, cost);
            }
            bytes32 rootHash = BasHash.Hash(rName);
            bool isOpen;
            bool isCustom;
            uint256 cusPrice;
            address rootOwner;
            uint256 rootExpire;
            (rootOwner, rootExpire) = BasTradableOwnership(rel.tro()).ownerOfWithExpire(rootHash);
            if(rootExpire > block.timestamp){
                (, isOpen, isCustom, cusPrice) = BasRootDomain(rel.root()).Root(rootHash);
                if(!isOpen && msg.sender!=rootOwner){
                    return (5, cost);
                }
            }else if(rootOwner == address(0)){
                if(!BasRootDomain(rel.root()).verifyRoot(rName)){
                    return (1, cost);
                }
            }
            if(rootExpire > block.timestamp && isCustom){
                cost = cusPrice.mul(durationInYear);
            }else{
                cost = BasOANN(rel.oann()).SUB_GAS().mul(durationInYear);
            }
            return (0, cost);
    }
    
    function queryDomainInfo(bytes32 nameHash)
                        external
                        view
                        returns(bytes memory name,
                                address owner,
                                uint256 expiration,
                                bool isRoot,
                                bool rIsOpen,
                                bool rIsCustom,
                                bool rIsRare,
                                uint256 rCusPrice,
                                bytes32 sRootHash,
                                bool isMarketOrder){
        BasRootDomain root = BasRootDomain(rel.root());
        (owner, expiration) = BasTradableOwnership(rel.tro()).ownerOfWithExpire(nameHash);
        if(root.hasDomain(nameHash)){
            isRoot = true;
            (name, rIsOpen, rIsCustom, rCusPrice) = root.Root(nameHash);
            rIsRare = root.isRare(nameHash);
        }else if(BasSubDomain(rel.sub()).hasDomain(nameHash)){
            (name,sRootHash) = BasSubDomain(rel.sub()).Sub(nameHash);
        }
        isMarketOrder = BasMarket(rel.market()).DomainSellOrders(msg.sender, nameHash) > 0;
    }
    
    function queryDomainEmailInfo(bytes32 nameHash)
                    external
                    view
                    returns(bytes memory name,
                            address owner,
                            uint256 expiration,
                            bool isActive,
                            bool openToPublic){
        BasRootDomain root = BasRootDomain(rel.root());
        (owner, expiration) = BasTradableOwnership(rel.tro()).ownerOfWithExpire(nameHash);
        (isActive, openToPublic) = BasMailManager(rel.mm()).mailConfigs(nameHash);
        if(root.hasDomain(nameHash)){
            (name, , ,) = root.Root(nameHash);
        }else if(BasSubDomain(rel.sub()).hasDomain(nameHash)){
            (name,) = BasSubDomain(rel.sub()).Sub(nameHash);
        }
    }
    
    function queryEmailInfo(bytes32 mailHash)
                    external
                    view
                    returns(address owner,
                            uint256 expiration,
                            bytes32 domainHash,
                            bool isValid,
                            bytes memory aliasName,
                            bytes memory bcAddress){
        (domainHash, , isValid, aliasName, bcAddress) = BasMail(rel.mail()).MailRecords(mailHash);
        (owner, expiration) = BasExpiredOwnership(rel.exo()).ownerOfWithExpire(mailHash);
    }
    
    function queryOrderInfo(address seller,
                            bytes32 nameHash)
                    external
                    view
                    returns(bytes memory name,
                            uint256 price,
                            bool isValid){
        BasRootDomain root = BasRootDomain(rel.root());
        if(root.hasDomain(nameHash)){
            (name, , ,) = root.Root(nameHash);
        }else if(BasSubDomain(rel.sub()).hasDomain(nameHash)){
            (name,) = BasSubDomain(rel.sub()).Sub(nameHash);
        }
        BasMarket market = BasMarket(rel.market());
        price = market.DomainSellOrders(seller, nameHash);
        (address owner, uint256 expiration) = BasExpiredOwnership(rel.tro()).ownerOfWithExpire(nameHash);
        isValid = (owner == seller && expiration > (block.timestamp).add(market.AT_LEAST_REMAIN_TIME()));
    }
    
    function queryDomainConfigs(bytes32 nameHash)
                    external
                    view
                    returns(bytes memory A,
                            bytes memory AAAA,
                            bytes memory MX,
                            bytes memory BlockChain,
                            bytes memory IOTA,
                            // bytes memory Optional,
                            bytes memory CName,
                            bytes memory MXBCA){
        BasDomainConf conf = BasDomainConf(rel.conf());
        A           = conf.domainConfData(nameHash, "A");
        AAAA        = conf.domainConfData(nameHash, "AAAA");
        MX          = conf.domainConfData(nameHash, "MX");
        BlockChain  = conf.domainConfData(nameHash, "BlockChain");
        IOTA        = conf.domainConfData(nameHash, "IOTA");
        CName       = conf.domainConfData(nameHash, "CName");
        MXBCA       = conf.domainConfData(nameHash, "MXBCA");
    }
}