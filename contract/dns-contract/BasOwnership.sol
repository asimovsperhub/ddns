//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.5.0;

import "./BasLib.sol";

/*
[INHERIT]
contract inherit this will be partially under control of DAO 
*/
contract  ManagedByDAO {
    
    address public DAOAddress;

    constructor() {
        DAOAddress = msg.sender;
    }

    modifier OnlyDAO {
        require(msg.sender == DAOAddress, "admin only");
        _;
    }

    function ChangeDAO(address newDao)
                            external
                            OnlyDAO {
        DAOAddress = newDao;
    }
}

interface UpdateAble{
    function getPrecursor() external view returns(address);
}

/*
[DEPLOYED]
this contract governs the relationship between contracts
*/
contract BasRelations is ManagedByDAO, UpdateAble{

    address precursor;

    address public token;
    address public exo;
    address public tro;
    address public root;
    address public sub;
    address public conf;
    address public acc;
    address public miner;
    address public oann;
    address public market;
    address public mail;
    address public mm;

    constructor(address pre) {
        precursor = pre;
    }

    //this show one contract can be modified by which contract
    mapping(address => mapping(address => bool)) public canBeModified;

    function getPrecursor() override external view returns(address){
        return precursor;
    }

    function setTokenAddr(address token_addr) external OnlyDAO{
        token = token_addr;
    }
    function setTroAddr(address tro_addr) external OnlyDAO{
        tro = tro_addr;
    }
    function setExoAddr(address exo_addr) external OnlyDAO{
        exo = exo_addr;
    }
    function setRootAddr(address root_addr) external OnlyDAO{
        root = root_addr;
    }
    function setSubAddr(address sub_addr) external OnlyDAO{
        sub = sub_addr;
    }
    function setConfAddr(address conf_addr) external OnlyDAO{
        conf = conf_addr;
    }
    function setAccAddr(address acc_addr) external OnlyDAO{
        acc = acc_addr;
    }

    function setMinerAddr(address miner_addr) external OnlyDAO{
        miner = miner_addr;
    }
    function setOANNAddr(address oann_addr) external OnlyDAO{
        oann = oann_addr;
    }
    function setMarketAddr(address market_addr) external OnlyDAO{
        market = market_addr;
    }
    function setMailAddr(address mail_addr) external OnlyDAO{
        mail = mail_addr;
    }
    function setMailManagerAddr(address mm_addr) external OnlyDAO{
        mm = mm_addr;
    }

    function SwitchOn()  external  OnlyDAO{
        canBeModified[tro][root] = true;
        canBeModified[tro][sub] = true;
        canBeModified[tro][market] = true;
        canBeModified[root][oann] = true;
        canBeModified[sub][oann] = true;
        canBeModified[acc][oann] = true;
        canBeModified[acc][mm] = true;
        canBeModified[miner][acc] = true;
        canBeModified[mail][mm] = true;
        canBeModified[exo][mail] = true;
    }

    function addCanByModified(address data, address caller)
                                external
                                OnlyDAO{
        canBeModified[data][caller] = true;
    }

    function changeRelationAll(address new_rel)
                                external
                                OnlyDAO{
        //this check is just for misoperation, can't guard against intentional damage
        require(UpdateAble(new_rel).getPrecursor() == address(this), "update check failed");
        HasRelations(exo).changeRelation(new_rel);
        HasRelations(tro).changeRelation(new_rel);
        HasRelations(root).changeRelation(new_rel);
        HasRelations(sub).changeRelation(new_rel);
        HasRelations(acc).changeRelation(new_rel);
        HasRelations(miner).changeRelation(new_rel);
        HasRelations(oann).changeRelation(new_rel);
        HasRelations(market).changeRelation(new_rel);
        HasRelations(mail).changeRelation(new_rel);
        HasRelations(mm).changeRelation(new_rel);
    }
}

/*
[INHERIT]
contract inherit this will have relations managed by BasRelations
*/
contract HasRelations {

    BasRelations public rel;

    constructor(address addr) {
        rel = BasRelations(addr);
    }

    modifier CanBeModified(){
        require(rel.canBeModified(address(this), msg.sender), "not allowed to modify");
        _;
    }

    function changeRelation(address new_rel) external {
        require(msg.sender == address(rel), "not accepted");
        rel = BasRelations(new_rel);
    }

}

/*
[DEPLOYED]
contract like this is a none tradable ownership database
(e.g email ownership)
*/
contract BasExpiredOwnership is ManagedByDAO, HasRelations{

    using SafeMath for uint256;
    using BasSet for BasSet.IndexedAsset;

    /*
    there is no need to index any of those events
    */
    event Add(bytes32 nameHash, address owner);
    event Update(bytes32 nameHash, address owner);
    event Extend(bytes32 nameHash, uint256 time);
    event Takeover(bytes32 nameHash, address from, address to);

    struct ownerShip{
        address ownerAddr;
        uint256 expireDate;
    }

    mapping(bytes32 => ownerShip) internal _ownerShips;
    mapping(address => BasSet.IndexedAsset) internal _assetsOf;

    constructor(address rel) HasRelations(rel) {}

    modifier OnlyOwner(bytes32 nameHash) {
        ownerShip memory os = _ownerShips[nameHash];
        require(os.ownerAddr == msg.sender && os.expireDate > block.timestamp, "owner only");
        _;
    }

    modifier CanRegister(bytes32 nameHash) {
        ownerShip memory os = _ownerShips[nameHash];
        require(os.ownerAddr == address(0) || os.expireDate < block.timestamp, "not valid");
        _;
    }

    function ownerOf(bytes32 nameHash)
                    external
                    view
                    returns (address){

        return _ownerShips[nameHash].ownerAddr;
    }

    function ownerOfWithExpire(bytes32 nameHash)
                    public
                    view
                    returns (address,uint256){

        ownerShip memory os = _ownerShips[nameHash];
        return (os.ownerAddr, os.expireDate);
    }

    function updateByDaoProposal(bytes32 nameHash,
                                address owner,
                                uint256 expire)
                                external
                                OnlyDAO {

        ownerShip storage os = _ownerShips[nameHash];

        //enforce  to clear data
        _assetsOf[os.ownerAddr].trimIfExist(nameHash);
        _assetsOf[owner].trimIfExist(nameHash);
        _assetsOf[owner].append(nameHash);
        os.ownerAddr = owner;
        os.expireDate = expire;

        emit Update(nameHash, owner);
    }

    function newOwnership(bytes32 nameHash,
                        address owner,
                        uint256 expire)
                        external
                        CanBeModified
                        CanRegister(nameHash) {

        _ownerShips[nameHash] = ownerShip(owner, expire);
        _assetsOf[owner].append(nameHash);

        emit Add(nameHash, owner);
    }

    function extendTime(bytes32 nameHash,
                        uint256 extend)
                        external
                        CanBeModified
                        returns (uint256){

        ownerShip storage os = _ownerShips[nameHash];
        require(os.ownerAddr != address(0), "extend to unknown");
//        uint256 memory nowtime = block.timestamp;
        if (os.expireDate < block.timestamp){
            os.expireDate = extend.add(block.timestamp);
        }else{
            os.expireDate = os.expireDate.add(extend);
        }
        emit Extend(nameHash, extend);
        return os.expireDate;
    }

    function _relpaceOwnership(address oldOwner,
                            address newOwner,
                            bytes32 nameHash)
                            internal{
        _assetsOf[oldOwner].trimIfExist(nameHash);
        _assetsOf[newOwner].append(nameHash);
        _ownerShips[nameHash].ownerAddr = newOwner;
    }
    
    function takeover(bytes32 nameHash,
                    address owner,
                    uint256 expire)
                    external
                    CanBeModified
                    CanRegister(nameHash) {

        address oldOwner = _ownerShips[nameHash].ownerAddr;

        _relpaceOwnership(oldOwner, owner, nameHash);
        _ownerShips[nameHash].expireDate = expire;

        emit Takeover(nameHash, oldOwner, owner);
    }

    function assetsCountsOf()
                external
                view
                returns(uint256){
                return _assetsOf[msg.sender].counts();
    }

    function assetsOf(uint256 start,
                    uint256 end)
                external
                view
                returns(bytes32[] memory){
                return _assetsOf[msg.sender].slice(start,end);
    }
}

/*
[DEPLOYED]
contract like this is a tradable ownership database
(e.g domain ownership)
*/
contract BasTradableOwnership is BasExpiredOwnership{

    /*
    there is also no need to keep index on events
    because Transfer or TransferFrom can happen on different
    scenerios and we should keep records on business level
    */
    event Transfer(bytes32 nameHash, address from, address to);
    event TransferFrom(bytes32 nameHash, address from, address to, address by);
    event Approval(address from, address to, bytes32 nameHash);
    event Revoke(address from, bytes32 nameHash);

    mapping(address => mapping(bytes32 => address)) internal _allowed;

    constructor(address rel) BasExpiredOwnership(rel)  {}

    function transfer(bytes32 nameHash,
                    address to)
                    external
                    OnlyOwner(nameHash) {

        require(msg.sender != to, "transfer to self");

        _relpaceOwnership(msg.sender, to, nameHash);

        emit Transfer(nameHash, msg.sender, to);
    }

    function approve(bytes32 nameHash,
                    address spender)
                    external
                    OnlyOwner(nameHash){

        _allowed[msg.sender][nameHash] = spender;
        emit Approval(msg.sender, spender, nameHash);
    }

    function allowance(address owner,
                        bytes32 nameHash)
                        external
                        view
                        returns (address) {

        return _allowed[owner][nameHash];
    }

    function revoke(bytes32 nameHash)
                    external
                    OnlyOwner(nameHash) {

        delete _allowed[msg.sender][nameHash];
        emit Revoke(msg.sender, nameHash);
    }

    function transferFrom(bytes32 nameHash,
                        address from,
                        address to)
                        external {

        require(_ownerShips[nameHash].ownerAddr == from, "owned no more");
        require(_allowed[from][nameHash] == msg.sender, "not allowed");
        require(from != to, "transfer to self");

        delete _allowed[from][nameHash];

        _relpaceOwnership(from, to, nameHash);

        emit TransferFrom(nameHash, from, to, msg.sender);
    }
}