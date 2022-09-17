//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.5.0;
//ICANN（The Internet Corporation for Assigned Names and Numbers）
//BOANN（The Blockchain Organizaiton for Assigned Names and Numbers）

import "./BasToken.sol";
import "./BasLib.sol";
import "./BasMiner.sol";
import "./BasDomain.sol";

/*
[DEPLOYED]
this contract manages domain registry and some data change options that costs bas
and keeps all receiptID with profit allocation plan meanwhile
*/
contract BasOANN is ManagedByDAO, HasRelations{
    using SafeMath for uint256;

    event PriceChanged();

    event AllocationChanged();

    event Cost(uint256 amount);

    uint256 public MAX_YEAR             ;//= 5
    uint256 public AROOT_GAS            ;//= 2000 * (10**18);
    uint256 public BROOT_GAS            ;//= 200 * (10**18);
    uint256 public SUB_GAS              ;//= 4 * (10**18);
    uint256 public CUSTOM_PRICE_GAS     ;//= 100 * (10**18);
    uint256 public REG_ROOT_M           ;//= 40;
    uint256 public REG_SELF_SUB_M       ;//= 40;
    uint256 public REG_SELF_SUB_O       ;//= 0;
    uint256 public REG_NORMAL_SUB_M     ;//= 40;
    uint256 public REG_NORMAL_SUB_O     ;//= 20;
    uint256 public REG_CUSTOEMED_SUB_M  ;//= 40;
    uint256 public REG_CUSTOEMED_SUB_O  ;//= 15;
    
    function priceSetting(uint256 max_year,
                        uint256 aroot_gas,
                        uint256 broot_gas,
                        uint256 sub_gas,
                        uint256 custom_price_gas)
                        public
                        OnlyDAO{
            MAX_YEAR                    =   max_year;
            AROOT_GAS                   =   aroot_gas;
            BROOT_GAS                   =   broot_gas;
            SUB_GAS                     =   sub_gas;
            CUSTOM_PRICE_GAS            =   custom_price_gas;

            emit PriceChanged();
    }


    function allocationSetting(uint256 reg_root_m,
                        uint256 reg_self_sub_m,
                        uint256 reg_self_sub_o,
                        uint256 reg_normal_sub_m,
                        uint256 reg_normal_sub_o,
                        uint256 reg_custom_sub_m,
                        uint256 reg_custom_sub_o)
                        public
                        OnlyDAO{

            require(reg_root_m <= 100 &&
                    reg_self_sub_m.add(reg_self_sub_o) <= 100 &&
                    reg_normal_sub_m.add(reg_normal_sub_o) <= 100 &&
                    reg_custom_sub_m.add(reg_custom_sub_o) <= 100,
                    "param error");

            REG_ROOT_M                  =   reg_root_m;
            REG_SELF_SUB_M              =   reg_self_sub_m;
            REG_SELF_SUB_O              =   reg_self_sub_o;
            REG_NORMAL_SUB_M            =   reg_normal_sub_m;
            REG_NORMAL_SUB_O            =   reg_normal_sub_o;
            REG_CUSTOEMED_SUB_M         =   reg_custom_sub_m;
            REG_CUSTOEMED_SUB_O         =   reg_custom_sub_o;


            emit AllocationChanged();
    }

    constructor(address rel) HasRelations(rel)  {
            priceSetting(5, 2000 * (10**18), 200 * (10**18), 4 * (10**18), 100 * (10**18));
            allocationSetting(40, 40, 0, 40, 20, 40, 15);

    }

    modifier validDuration(uint8 y)  {
        require (y <= MAX_YEAR && y > 0, "year out of range");
        _;
    }

    /*
    this profit allocation is for root domain options
    */
    function _allocateProfit_m(uint256 share_m,
                        uint cost)
                        internal{
        if (share_m > 0){
            BasAccountant(rel.acc()).allocateProfitAdvanced(rel.miner(), cost.mul(share_m).div(100));
        }
    }
    
    /*
    this profit allocation is for sub domain options
    */
    function _allocateProfit_m_o(uint256 share_m,
                        uint256 share_o,
                        uint cost,
                        address domainOwner)
                        internal{
        BasAccountant acc = BasAccountant(rel.acc());
        if (share_m > 0){
            acc.allocateProfitAdvanced(rel.miner(), cost.mul(share_m).div(100));
        }
        if (domainOwner != address(0) && share_o > 0){
            acc.allocateProfitNormal(domainOwner, cost.mul(share_o).div(100));
        }
    }

    function _payToAccountant(uint256 amount) internal {
        ERC20(rel.token()).transferFrom(msg.sender, rel.acc(), amount);
    }
 
    /*
    when register a root, there are some checks and some changes about storage,
    first, validity of characters are checked here, so is rareness,
    then the cost is sumed and token transfered to miners
    then the domain contract decide if this registry is new or takeover
    lastly the ownership record the new expiration date
    */
    function registerRoot(bytes calldata name,
                    bool isOpen,
                    bool isCustom,
                    uint256 cusPrice,
                    uint8 durationInYear)
                    external
                    validDuration(durationInYear){

        BasRootDomain root = BasRootDomain(rel.root());
        (bool isValid, bool isRare) = root.classifyRoot(name);
        require(isValid, "invalid domain name");

        uint256 cost;
        if (isRare) {
            cost = AROOT_GAS.mul(durationInYear);
        } else {
            cost = BROOT_GAS.mul(durationInYear);
        }

        if (isCustom) {
            require(cusPrice >= SUB_GAS, "can't be lower than system price");
            cost = cost.add(CUSTOM_PRICE_GAS);
        }

        _payToAccountant(cost);
        _allocateProfit_m(REG_ROOT_M, cost);

        root.replaceOrCreate(name, (block.timestamp).add(durationInYear * 365 days), isOpen, isCustom, cusPrice, msg.sender);
        emit Cost(cost);
    }

    /*
    any one can recharge an existing domain
    */
    function rechargeRoot(bytes32 nameHash,
                    uint8 durationInYear)
                    external
                    validDuration(durationInYear){

        BasRootDomain root = BasRootDomain(rel.root());
        bool isRare = root.isRare(nameHash);

        uint256 cost;
        if (isRare){
            cost = AROOT_GAS.mul(durationInYear);
        }else{
            cost = BROOT_GAS.mul(durationInYear);
        }

        _payToAccountant(cost);
        _allocateProfit_m(REG_ROOT_M, cost);

        root.recharge(nameHash, durationInYear * 365 days, (block.timestamp).add(MAX_YEAR * 365 days));
        emit Cost(cost);
    }

    /*
    each time user want to change price, it should pay some token
    to prevent from constantly changes
    */
    function openCustomPrice(bytes32 nameHash,
                            uint256 price)
                            external{

        require(price > SUB_GAS, "can't set price lower than default");

        _payToAccountant(CUSTOM_PRICE_GAS);
        _allocateProfit_m(REG_ROOT_M, CUSTOM_PRICE_GAS);

        BasRootDomain(rel.root()).openCustomPrice(nameHash, price, msg.sender);
        emit Cost(CUSTOM_PRICE_GAS);
    }

    /*
    registering sub domain takes more steps than root,
    if we are registering a subdomain of an existing root, we can skip the root validity check
    otherwise we should perform a full check
    second, based on existence and openness and account address, decide if able to register
    then sum cost and transfer token
    */
    function registerSub(bytes calldata rName,
                    bytes calldata sName,
                    uint8 durationInYear)
                    external
                    validDuration(durationInYear){

        require(BasSubDomain(rel.sub()).verifySub(sName), "subname not valid");
        bytes32 rootHash = BasHash.Hash(rName);
        bool isOpen;
        bool isCustom;
        uint256 customPrice;
        (address rootOwner, uint256 expire) = BasTradableOwnership(rel.tro()).ownerOfWithExpire(rootHash);
        if (expire > block.timestamp){      //root exists and incharge
            (,isOpen, isCustom, customPrice) = BasRootDomain(rel.root()).Root(rootHash);
            require(isOpen || msg.sender == rootOwner, "root disallow registry");
        }else if(rootOwner == address(0)){      //domain not exist
                require(BasRootDomain(rel.root()).verifyRoot(rName),"root not valid");
        }
        //skip else: root exists but expired
        
        uint256 cost = _decideCostAndDeliverToken(expire > block.timestamp,isCustom, msg.sender == rootOwner, customPrice, durationInYear, rootOwner);
        bytes memory totalName = abi.encodePacked(sName, ".", rName);

        BasSubDomain(rel.sub()).replaceOrCreate(totalName, rootHash, (block.timestamp).add((durationInYear * 365 days)), msg.sender);

        emit Cost(cost);
    }


    /*
    in most case, we should not check root validity again, but in case root of rareness changes,
    recheck is needed,
    */
    function rechargeSub(bytes32 nameHash,
                    uint8 durationInYear)
                    external
                    validDuration(durationInYear){
        BasSubDomain sub = BasSubDomain(rel.sub());
        (,bytes32 rootHash) = sub.Sub(nameHash);
        
        bool isOpen;
        bool isCustom;
        uint256 customPrice;
        (address rootOwner, uint256 expire) = BasTradableOwnership(rel.tro()).ownerOfWithExpire(rootHash);
        if(expire > block.timestamp){
            (,isOpen, isCustom, customPrice) = BasRootDomain(rel.root()).Root(rootHash);
            require(isOpen || msg.sender == rootOwner, "root disallow recharge");
        }
        uint256 cost = _decideCostAndDeliverToken(expire > block.timestamp, isCustom, msg.sender == rootOwner, customPrice, durationInYear, rootOwner);
        sub.recharge(nameHash, durationInYear * 365 days, block.timestamp.add(MAX_YEAR * 365 days));
        emit Cost(cost);
    }

    function _decideCostAndDeliverToken(bool rootValid,
                                    bool isCustom,
                                    bool isSelf,
                                    uint256 customPrice,
                                    uint8 durationInYear,
                                    address rootOwner)
                            internal
                            returns (uint256 cost){
        if (rootValid && isCustom){
            cost = customPrice.mul(durationInYear);
        }else{
            cost = SUB_GAS.mul(durationInYear);
        }
        if (rootValid){
            if(isSelf){
                _allocateProfit_m_o(REG_SELF_SUB_M, REG_SELF_SUB_O, cost, rootOwner);
            }else if(isCustom){
                _allocateProfit_m_o(REG_CUSTOEMED_SUB_M, REG_CUSTOEMED_SUB_O, cost, rootOwner);
            }else{
                _allocateProfit_m_o(REG_NORMAL_SUB_M, REG_NORMAL_SUB_O, cost, rootOwner);
            }
        }else{
            _allocateProfit_m_o(REG_NORMAL_SUB_M, REG_NORMAL_SUB_O, cost, address(0));
        }
        _payToAccountant(cost);
    }
}
