//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.5.0;

import "./BasDomain.sol";
import "./BasOwnership.sol";
import "./BasToken.sol";
import "./BasMiner.sol";
import "./BasLib.sol";

/*
[DEPLOYED]
this contract keeps mail data
*/
contract BasMail is HasRelations{
    
    /*
    we keep mailHash rather than mail string
    to prevent mail address crawler aquiring all mails
    */
    struct MailRecord {
        bytes32     domainHash;
        bytes32     mailHash;
        bool        valid;
        bytes       aliasName;
        bytes       bcAddress;
    }
    
    event NewMail(bytes32 domainHash,
                bytes32 nameHash,
                address owner);
    event MailUpdate(bytes32 domainHash,
                    bytes aliasName,
                    bytes bcAddress);
    event AbandMail(bytes32 domainHash);
    event MailRecharged(bytes32 domainHash,
                    uint extendTime);
                    
    constructor(address rel) HasRelations(rel) {}
    
    mapping(bytes32 => MailRecord) public MailRecords;
    
    modifier OnlyAlive(bytes32 hash){
        require(MailRecords[hash].valid, "the mail address has been abandoned");
        _;
    }
    
    modifier OnlyOwner(address owner, 
                    bytes32 hash){
        (address curOwner, uint256 expire) = BasExpiredOwnership(rel.exo()).ownerOfWithExpire(hash);
        require(curOwner == owner && expire > block.timestamp, "only owner is invalid");
        _;
    }

    function domainHashOf(bytes32 mailHash) 
                    view 
                    external 
                    returns(bytes32){
        return MailRecords[mailHash].domainHash;
    }
    
    function updateMail(bytes32 mailHash,
                        bytes calldata aliasName,
                        bytes calldata bcAddress)
                        OnlyOwner(msg.sender, mailHash)
                        OnlyAlive(mailHash)
                        external{

        MailRecords[mailHash].aliasName = aliasName;
        MailRecords[mailHash].bcAddress = bcAddress;

        emit MailUpdate(mailHash, aliasName, bcAddress);
    }
    
    function newEmail(bytes32 domainHash,
                    bytes32 mailHash,
                    address owner,
                    uint expire,
                    bytes calldata aliasName)
                    external 
                    CanBeModified{
        
        require(MailRecords[mailHash].domainHash == bytes32(0), "email has been registered");
        
        BasExpiredOwnership(rel.exo()).newOwnership(mailHash, owner, expire);
        
        MailRecords[mailHash] = MailRecord(domainHash, mailHash, true, aliasName, '');
        
        emit NewMail(domainHash, mailHash, owner);
    }
    
    function recharge(bytes32 mailHash,
                    uint extendTime)
                    external
                    OnlyAlive(mailHash) 
                    CanBeModified{
        
        BasExpiredOwnership ownership = BasExpiredOwnership(rel.exo());
        
        ownership.extendTime(mailHash, extendTime);
        
        emit MailRecharged(mailHash, extendTime);
    }
    
    /*
    we don't simply delete mail data, because there is a chance the mail is re-registered
    which may cause secure problem
    */
    function abandon(bytes32 hash)
                OnlyOwner(msg.sender, hash)
                external{
        
        MailRecords[hash].valid = false;
        
        MailRecords[hash].aliasName = "";
        MailRecords[hash].bcAddress = "";
        
        emit AbandMail(hash);
    }
}

/*
[DEPLOYED]
this contract manages mail registry and mail admin options
*/
contract BasMailManager is ManagedByDAO, HasRelations{
    using SafeMath for uint256;
    
    
    event PriceChanged();
    event AllocationChanged();
    event MailServerActive();
    event MailServerInactive();
    event MailServerOpenChanged(bytes32 domainHash,
                                bool isOpen);
    event Cost(uint256 amount);
                
    event UpdateConf();

    
    uint256 public MAX_MAIL_YEAR        = 5;
    uint256 public OPEN_ACTION_GAS      = 100 * (10 ** 18);
    uint256 public REG_MAIL_GAS         = 2 * (10 ** 18);
    uint256 public BASIC_M              = 40;
    uint256 public TOP_DOMAIN_M         = 40;
    uint256 public TOP_DOMAIN_O         = 20;

    
    bool commonRootCanServe = true;
    bool commonRootCanPublic = false;
    bool subCanServe = false;
    bool subCanPublic = false;
    
    function changeSettings(bool cRootServe,
                            bool cRootPub,
                            bool subServe,
                            bool subPub)
                            external
                            OnlyDAO{
                commonRootCanServe  = cRootServe;
                commonRootCanPublic = cRootPub;
                subCanServe         = subServe;
                subCanPublic        = subPub;
    }
    
    
    function priceSetting(uint256 max_mail_year,
                        uint256 open_action_gas,
                        uint256 reg_mail_gas)
                        public
                        OnlyDAO{
                MAX_MAIL_YEAR           = max_mail_year;
                OPEN_ACTION_GAS         = open_action_gas;
                REG_MAIL_GAS            = reg_mail_gas;
                
                emit PriceChanged();
    }
    
    function allocationSetting(uint256 basic_m,
                            uint256 top_domain_m,
                            uint256 top_domain_o)
                            public
                            OnlyDAO{
                                
                require(basic_m <= 100 &&
                        top_domain_m.add(top_domain_o) <= 100,
                        "param error");
                BASIC_M                 = basic_m;
                TOP_DOMAIN_M            = top_domain_m;
                TOP_DOMAIN_O            = top_domain_o;
                
                emit AllocationChanged();
    }
    
    struct MailServiceConf{
        bool    active;
        bool    openToPublic;
    }
    
    mapping(bytes32 => MailServiceConf) public    mailConfigs;

    constructor(address rel) HasRelations(rel) {}
    
    modifier ValidEmailDuration(uint8 durationInYear){
        require(durationInYear > 0 && 
                durationInYear <= MAX_MAIL_YEAR, 
                "Invalid email duration");
        _;
    }
    
    modifier OnlyDomainOwner(bytes32 domainHash){
        (address owner, uint256 expire) = BasTradableOwnership(rel.tro()).ownerOfWithExpire(domainHash);
        require(owner == msg.sender && expire > block.timestamp, "not owned or expired");
        _;
    }
    
    /*
    to check openMailService, set active = false, means current active status should be false
    to check setPublic, set active = true, mean current active status should be true
    */
    modifier ruleCheck(bytes32 domainHash,
                        bool active,
                        bool openToPublic){
        require(active == mailConfigs[domainHash].active, "active not same");
        if (BasSubDomain(rel.sub()).hasDomain(domainHash)){     //sub
            if(!active){
                require(subCanServe, "sub domain can't serve");
            }
            if (openToPublic){
                require(subCanPublic, "sub domian can't be public");
            }
        }else if(!BasRootDomain(rel.root()).isRare(domainHash)){  //common root
            if(!active){
                require(commonRootCanServe, "common root can't serve");
            }
            if (openToPublic){
                require(commonRootCanPublic, "common root can't be public");
            }
        }
        //rare root skips all check
        _;
    }
    
    function changeConfByDaoProposal(bytes32 nameHash,
                                    bool active,
                                    bool openToPublic)
                                    external
                                    OnlyDAO{
        mailConfigs[nameHash] = MailServiceConf(active, openToPublic);
        emit UpdateConf();
    }
    

    function openMailService(bytes32 domainHash,
                            bool openToPublic)
                        external
                        OnlyDomainOwner(domainHash)
                        ruleCheck(domainHash, false, openToPublic){
        
        ERC20(rel.token()).transferFrom(msg.sender, rel.acc(), OPEN_ACTION_GAS);
        
        mailConfigs[domainHash] = MailServiceConf(true, openToPublic);
        
        BasAccountant(rel.acc()).allocateProfitAdvanced(rel.miner(), OPEN_ACTION_GAS.mul(BASIC_M).div(100));
        
        emit MailServerActive();
        if (openToPublic){
            emit MailServerOpenChanged(domainHash, true);
        }
    }


    function removeMailServer(bytes32 domainHash)
                        OnlyDomainOwner(domainHash) external{
        require(mailConfigs[domainHash].active, "not active");
        delete mailConfigs[domainHash];
        
        emit MailServerInactive();
        if (mailConfigs[domainHash].openToPublic){
            emit MailServerOpenChanged(domainHash, false);
        }
    }

    function setPublic(bytes32 domainHash,
                        bool openToPublic)
                        external
                        ruleCheck(domainHash, true, openToPublic){

        mailConfigs[domainHash].openToPublic = openToPublic;
        emit MailServerOpenChanged(domainHash, openToPublic);
    }
    
    /*
    there are generally two type of mail domain
    1. rare root domain, user can register mail when domain is open to public
    2. other domain, only domain user or domain admin can register mail
    */
    function registerMail(bytes32 domainHash,
                         bytes32 mailHash,
                        uint8 durationInYear,
                        bytes calldata aliasName)
                        ValidEmailDuration(durationInYear)
                        external{
        
        uint256 cost = _payForMail(domainHash, durationInYear);
        
        BasMail(rel.mail()).newEmail(domainHash, mailHash, msg.sender, (block.timestamp).add(durationInYear * 365 days), aliasName);
        emit Cost(cost);
    }
    
    function recharge(bytes32 mailHash, 
                    uint8 durationInYear) 
                    ValidEmailDuration(durationInYear)
                    external{
        
        BasMail mail = BasMail(rel.mail());
        bytes32 domainHash = mail.domainHashOf(mailHash);

        uint256 cost = _payForMail(domainHash, durationInYear);
        
        mail.recharge(mailHash, durationInYear * 365 days);
        emit Cost(cost);
    }
    
    function _payForMail(bytes32 domainHash,
                        uint256 durationInYear) 
                        internal
                        returns (uint256){
        MailServiceConf memory conf = mailConfigs[domainHash];
        require(conf.active, "domain hasn't open email service");
        (address domainOwner, uint256 expire) = BasTradableOwnership(rel.tro()).ownerOfWithExpire(domainHash);
        require(expire > 0, "domain is expired");
        require( msg.sender == domainOwner ||
                conf.openToPublic,
                "not allowed to register");
        
        uint cost = REG_MAIL_GAS.mul(durationInYear);
        
        ERC20(rel.token()).transferFrom(msg.sender, rel.acc(), cost);
        _allocateProfit(cost,conf.openToPublic, domainOwner);
        return cost;
    }
    
    function _allocateProfit(uint cost,
                    bool isOpenToPub,
                    address topDomainOwner) private{
        BasAccountant acc = BasAccountant(rel.acc());
         if (isOpenToPub){
            acc.allocateProfitAdvanced(rel.miner(), cost.mul(TOP_DOMAIN_M).div(100));
            acc.allocateProfitNormal(topDomainOwner, cost.mul(TOP_DOMAIN_O).div(100));
        }else{
            acc.allocateProfitAdvanced(rel.miner(), cost.mul(BASIC_M).div(100));
        }
    }
}