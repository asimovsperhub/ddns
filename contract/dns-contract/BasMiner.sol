//SPDX-License-Identifier: UNLICENSED
pragma solidity >= 0.5.0;

import "./BasToken.sol";
import "./BasOwnership.sol";


interface ContractReceiver{
    function subAllocate(uint sum) external;
    function getBalance(address addr) external view returns (uint256);
    function withdraw() external;
}

/*
[DEPLOYED]
this contract works as ledger keeper
*/
contract BasAccountant is ManagedByDAO, HasRelations{

    using SafeMath for uint256;

    mapping(address=>uint) public ledger;

    constructor(address rel) HasRelations(rel) {}
    
    event Withdraw(address indexed drawer,
                    uint256 amount,
                    address to);
                    
    event Allocate(address indexed receiver,
                    uint256 amount);

    /*
    this function will be called by contract that costs token,
    like OANN, emailManager
    just keep ledger of an address
    */
    function allocateProfitNormal(address receiver,
                        uint value)
                        external
                        CanBeModified{
        ledger[receiver] = ledger[receiver].add(value);
        emit Allocate(receiver, value);
    }

    /*
    this function keep ledger of contract address
    and triggers its subAllocate interface
    */
    function allocateProfitAdvanced(address receiver,
                            uint value)
                            external
                            CanBeModified{
        ledger[receiver] = ledger[receiver].add(value);
        ContractReceiver(receiver).subAllocate(value);
        emit Allocate(receiver, value);
    }

    function withdrawTo(uint amount,
                address target)
                external { //CanBeModified
        uint256 balance = ledger[msg.sender];
        require(balance >= amount, "balance insufficient");

        ledger[msg.sender] = balance.sub(amount);
        ERC20(rel.token()).transfer(target, balance);
        
        emit Withdraw(msg.sender, amount, target);
    }

    function daoWithdraw(address to,
                        uint256 no)
                        external
                        OnlyDAO{
        ERC20(rel.token()).transfer(to, no);
    }

}
/*
[DEPLOYED]
this contract acts as contract receiver
*/

contract BasMiner is ManagedByDAO, ContractReceiver, HasRelations{
    using SafeMath for uint256;

    address[] public MainNode;
    mapping(address=>uint) public balanceOf;
    uint8 public constant MainNodeSize = 64;

    constructor(address rel) HasRelations(rel) {}

    /*
    we should make Withdraw indexed
    because miner can view how many it already withdrawed
    */
    event MinerAdd(address miner);
    event Withdraw(address indexed drawer, uint256 amout);
    event MinerRemove(address miner);
    event MinerReplace(address oldMiner, address newMiner);


    function GetAllMainNodeAddress()
                public
                view
                returns(address[] memory) {

        return MainNode;
    }

    function addMiner(address m)
                    external
                    OnlyDAO{

        require(MainNode.length < MainNodeSize, "nodes is full");

        MainNode.push(m);
        emit MinerAdd(m);
    }

    function replaceMiner(address oldM,
                    address newM)
                    external
                    OnlyDAO{

        require(MainNode.length > 0, "nodes empty");

        for (uint8 i = 0; i < MainNode.length; i++){
            if (MainNode[i] == oldM){
                MainNode[i] = newM;
                
                BasAccountant(rel.acc()).withdrawTo(balanceOf[oldM], oldM);
                balanceOf[oldM] = 0;
                emit MinerReplace(oldM, newM);
                return;
            }
        }
        revert("replace failed");
    }

    function removeMiner(address miner)
                    external
                    OnlyDAO{
        require(MainNode.length > 0, "nodes empty");
        
        for (uint8 i = 0; i < MainNode.length; i++){
            if (MainNode[i] == miner){
                MainNode[i] = MainNode[MainNode.length - 1];
                MainNode.pop();
                
                BasAccountant(rel.acc()).withdrawTo(balanceOf[miner], miner);
                balanceOf[miner] = 0;
                emit MinerRemove(miner);
                return;
            }
        }
        revert("remove failed");
    }

    function withdraw() override external{
        uint256 balance = balanceOf[msg.sender];
        require(balance > 0,"balance is 0");
        balanceOf[msg.sender] = 0;
        BasAccountant(rel.acc()).withdrawTo(balance, msg.sender);
        emit Withdraw(msg.sender, balance);
    }

    function subAllocate(uint256 sum)
                    override
                    external
                    CanBeModified{
        if(MainNode.length > 0){
            uint256 one_porift = sum.div(MainNode.length);
            for (uint8 i = 0; i < MainNode.length; i++){
                address miner_address = MainNode[i];
                balanceOf[miner_address] = balanceOf[miner_address].add(one_porift);
            }
        }
    }

    function getBalance(address addr)
                    override
                    external
                    view
                    returns (uint256){
        return balanceOf[addr];
    }
}