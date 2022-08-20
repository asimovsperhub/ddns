//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.5.0;

import "./BasOwnership.sol";

/*
[INHERIT]
this contract just provides two functions
*/
contract BasDomain is ManagedByDAO, HasRelations{

    uint8 constant MAX_DOMAIN_SEG_LEN   = 64;

    constructor(address rel) HasRelations(rel) {}

    //[0~9, a~z, -, _]
    function validChar(bytes1 c)
    internal
    pure
    returns (bool) {
        return
        (c >= 0x30 && c <= 0x39) ||
        (c >= 0x61 && c <= 0x7a) ||
        c == 0x2d ||
        c == 0x5f;
    }

    event Recharge(bytes32 nameHash,
                    uint256 duration);
    
    modifier OnlyOwner(address owner,
                    bytes32 hash){
        (address curOwner, uint256 expire) = BasTradableOwnership(rel.tro()).ownerOfWithExpire(hash);
        require(curOwner == owner && expire > block.timestamp, "only owner is invalid");
        _;
    }
    
    function recharge(bytes32 nameHash,
                uint rechargeTime,
                uint maxEnd)
                external
                CanBeModified{

        BasTradableOwnership ownership = BasTradableOwnership(rel.tro());
        require(ownership.extendTime(nameHash, rechargeTime) < maxEnd, "can't recharge that long");

        emit Recharge(nameHash, rechargeTime);
    }
}

/*
[DEPLOYED]
this contract keeps root domain data
*/
contract BasRootDomain is BasDomain {
    using SafeMath for uint256;

    uint8 public RARE_LENGTH = 6;

    constructor(address rel) BasDomain(rel) {}

    struct RootItem {
        bytes   rootName;
        bool    openToPublic;
        bool    isCustom;
        uint256 customPrice;
    }

    event RootUpdate(bytes32 nameHash,
                    bytes rootName,
                    bool openToPublic,
                    bool isCustom,
                    uint customPrice);

    event NewRootDomain(bytes32 indexed nameHash,
                    bytes rootName,
                    bool openToPublic,
                    bool isCustom,
                    uint customPrice);

    event RootDataReplaced(bytes32 nameHash,
                    bool openToPublic,
                    bool isCustom,
                    uint customPrice);
    event OpenCustomPrice(bytes32 nameHash,
                    uint customPrice);

    event ClosedCustomPrice(bytes32 nameHash);

    event RootPublicChanged(bytes32 nameHash,
                            bool isOpen);

    mapping(bytes32 => RootItem) public Root;

    function changeRareLength(uint8 length)
                    external
                    OnlyDAO{
                       RARE_LENGTH = length;
    }

    function classifyRoot(bytes memory name)
                    public
                    view
                    returns (bool, bool) {

        if (name.length == 0 || name.length >= MAX_DOMAIN_SEG_LEN) {
            return (false, false);
        }

        bool bRare = true;
        for (uint8 i = 0; i < name.length; i++) {

            if (!validChar(name[i])) {
                return (false, false);
            }

            if (bRare) {
                bRare = !(i >= RARE_LENGTH ||
                name[i] == 0x2d ||
                name[i] == 0x5f);
            }
        }

        return (true, bRare);
    }

    function classifyRoot(bytes32 nameHash)
                    external
                    view
                    returns (bool, bool){
            if (hasDomain(nameHash)){
                return classifyRoot(Root[nameHash].rootName);
            }else{
                return (false,false);
            }
    }


    /*
    because we check by nameHash so we can skip validation
    */
    function isRare(bytes32 nameHash)
                    external
                    view
                    returns (bool){

            if(!hasDomain(nameHash)){
                return false;
            }
            bytes memory name = Root[nameHash].rootName;
            for (uint8 i = 0; i < name.length; i++) {
                if (i >= RARE_LENGTH ||
                name[i] == 0x2d ||
                name[i] == 0x5f){
                    return false;
                }
            }
            return true;
    }

    function verifyRoot(bytes memory name)
                    public
                    pure
                    returns (bool) {

        if (name.length == 0 || name.length >= MAX_DOMAIN_SEG_LEN) {
            return false;
        }

        for (uint8 i = 0; i < name.length; i++) {

            if (!validChar(name[i])) {
                return false;
            }
        }

        return true;
    }

    function verifyRoot(bytes32 nameHash)
                    external
                    view
                    returns (bool){
            if(hasDomain(nameHash)){
                return verifyRoot(Root[nameHash].rootName);
            }else{
                return false;
            }
    }

    function hasDomain(bytes32 hash)
                    public
                    view
                    returns (bool){
        return Root[hash].rootName.length > 0;
    }
    function getNameByHash(bytes32 hash)
                    external
                    view
                    returns (bytes memory){
        return Root[hash].rootName;
    }
    
    /*
    create a new item or takeover ownership of expired item
    */
    function replaceOrCreate(bytes calldata rootName,
                        uint expire,
                        bool openToPublic,
                        bool isCustom,
                        uint256 customPrice,
                        address applicant)
                        external
                        CanBeModified {

        bytes32 nameHash = BasHash.Hash(rootName);
        BasTradableOwnership ownership = BasTradableOwnership(rel.tro());

        if(Root[nameHash].openToPublic != openToPublic){
            emit RootPublicChanged(nameHash, openToPublic);
        }

        Root[nameHash].openToPublic = openToPublic;
        Root[nameHash].isCustom = isCustom;
        Root[nameHash].customPrice = customPrice;

        if (ownership.ownerOf(nameHash) != address(0)){
            ownership.takeover(nameHash, applicant, expire);
            emit RootDataReplaced(nameHash, openToPublic, isCustom, customPrice);
        }else{
            ownership.newOwnership(nameHash, applicant, expire);
            Root[nameHash].rootName = rootName;
            emit NewRootDomain(nameHash, rootName, openToPublic, isCustom, customPrice);
        }
    }

    /*
    this function is called by OANN because open custom price costs BAS
    but the ownership is checked here, operator is surposed to be tx.origin
    */
    function openCustomPrice(bytes32 nameHash,
                        uint customPrice,
                        address operator)
                        external
                        OnlyOwner(operator, nameHash)
                        CanBeModified{

        Root[nameHash].isCustom = true;

        if(!Root[nameHash].openToPublic){
            emit RootPublicChanged(nameHash, true);
        }

        Root[nameHash].openToPublic = true;
        Root[nameHash].customPrice = customPrice;

        emit OpenCustomPrice(nameHash, customPrice);
    }

    /*
    called by user
    */
    function closeCustomPrice(bytes32 nameHash)
                            external
                            OnlyOwner(msg.sender, nameHash) {

        if(Root[nameHash].openToPublic){
            Root[nameHash].openToPublic = false;
            emit RootPublicChanged(nameHash, false);
        }

        Root[nameHash].isCustom = false;
        emit ClosedCustomPrice(nameHash);
    }

    function setPublic(bytes32 nameHash,
                    bool isOpen)
                    external
                    OnlyOwner(msg.sender,nameHash) {

        require(Root[nameHash].openToPublic != isOpen, "nothing changed");
        Root[nameHash].openToPublic = isOpen;

        emit RootPublicChanged(nameHash, isOpen);
    }
}

/*
[DEPLOYED]
this contract keeps sub domain data
*/
contract BasSubDomain is BasDomain{
    using SafeMath for uint256;

    uint16 constant MAX_SUB_DOMAIN_LEN = 192;

    bytes1 constant DOT_CHAR_VAL = 0x2e;

    constructor(address rel) BasDomain(rel) {}

    struct SubItem {
        bytes   totalName;
        bytes32 rootHash;
    }

    event SubUpdate(bytes32 nameHash,
                    bytes totalName,
                    bytes32 rootHash);

    event NewSubDomain(bytes32 indexed nameHash,
                    bytes totalName,
                    bytes32 rootHash);

    event SubDataReplaced(bytes32 nameHash);

    mapping(bytes32 => SubItem) public Sub;

    function verifySub(bytes memory name)
                    public
                    pure
                    returns (bool) {

        if (name.length >= MAX_SUB_DOMAIN_LEN){
            return false;
        }

        if (name[0] == DOT_CHAR_VAL || name[name.length-1] == DOT_CHAR_VAL){
            return false;
        }

        uint256 segementLength = 0;

        for (uint256 i = 0; i < name.length; i++) {

            bytes1 char = name[i];
            if (!validChar(char) && char != DOT_CHAR_VAL){
                return false;
            }

            if (char == DOT_CHAR_VAL){
                if (segementLength == 0){
                    return false;
                }

                segementLength = 0;
                continue;
            }

            segementLength += 1;
            if (segementLength >= MAX_DOMAIN_SEG_LEN){
                return false;
            }
        }

        return true;
    }

    function hasDomain(bytes32 hash)
                external
                view
                returns (bool){
        return Sub[hash].totalName.length > 0;
    }

    function getNameByHash(bytes32 hash)
                external
                view
                returns (bytes memory){
        return Sub[hash].totalName;
    }

    function replaceOrCreate(bytes calldata subdomain,
                    bytes32 rootHash,
                    uint newExpire,
                    address applicant)
                    external
                    CanBeModified{

        bytes32 subHash = BasHash.Hash(subdomain);
        BasTradableOwnership ownership = BasTradableOwnership(rel.tro());

        if (ownership.ownerOf(subHash) != address(0)) {
            ownership.takeover(subHash, applicant, newExpire);
            emit SubDataReplaced(subHash);
        } else {
            ownership.newOwnership(subHash, applicant, newExpire);
            Sub[subHash] = SubItem(subdomain, rootHash);
            emit NewSubDomain(subHash, subdomain, rootHash);
        }
    }
}

/*
[DEPLOYED]
this contract is used for DNS server, email server etc
*/
contract BasDomainConf is ManagedByDAO, HasRelations{

    event DomainConfChanged(bytes32 nameHash,
                            string typeName,
                            bytes data);

//those names are recommended, but user can add arbitrary names
    string[] public TypeNames;
    mapping(bytes32 => mapping(string => bytes)) public domainConfData;

    constructor(address rel) HasRelations(rel) {
        TypeNames = ["A","AAAA","MX","BlockChain","IOTA","Optional","CName","MXBCA"];
    }

    modifier OnlyOwner(address owner,
                    bytes32 hash){
        (address curOwner, uint256 expire) = BasTradableOwnership(rel.tro()).ownerOfWithExpire(hash);
        require(curOwner == owner && expire > block.timestamp, "only owner is invalid");
        _;
    }

    function addNewType(string calldata typName)
                    external
                    OnlyDAO{
        TypeNames.push(typName);
    }

    function dataMigrate(address otherConf,
                        bytes32[] calldata hashArray)
                            external
                            OnlyDAO{
        BasDomainConf other = BasDomainConf(otherConf);
        for(uint256 i = 0; i < hashArray.length; i++){
            for(uint256 j = 0; j < TypeNames.length; j++){
                bytes memory data = other.domainConfData(hashArray[i], TypeNames[j]);
                if (data.length > 0){
                    domainConfData[hashArray[i]][TypeNames[j]] = data;
                }
            }
        }
    }

    //support multiple ipv4 data, joined by ","
    function updateByOwner(bytes32 nameHash,
                        string calldata typName,
                        bytes calldata data)
                        external
                        OnlyOwner(msg.sender, nameHash){
        domainConfData[nameHash][typName] = data;
        emit DomainConfChanged(nameHash, typName, data);
    }

    function clearByOwner(bytes32 nameHash)
                        external
                        OnlyOwner(msg.sender, nameHash){
        for(uint256 i=0; i< TypeNames.length; i++){
            if(domainConfData[nameHash][TypeNames[i]].length > 0){
               delete domainConfData[nameHash][TypeNames[i]];
            }
        }
    }

    function query(bytes32 nameHash,
                string calldata typName)
                external
                view
                returns (bytes memory bca) {

        return domainConfData[nameHash][typName];
    }

}