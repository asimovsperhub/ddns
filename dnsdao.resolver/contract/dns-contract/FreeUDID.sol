//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.5.0;

import "./BasLib.sol";
import "./BasToken.sol";
import "./Owner.sol";


contract FreeUDID is owned{
    using SafeMath for uint256;

    IERC20 public UDIDToken;

    uint256 private amountPerDay;

    mapping(address => uint256) userTokenPerDay;

    constructor (address t){
        UDIDToken = IERC20(t);
        amountPerDay = 10 * 1e18;
    }

    function TransferUDID() external {

        require(userTokenPerDay[msg.sender] + 86400 < block.timestamp, "only 10 UDID per day free for user");

        if ( userTokenPerDay[msg.sender] == uint256(0)){
            userTokenPerDay[msg.sender] = block.timestamp;
        }else{
            userTokenPerDay[msg.sender] = userTokenPerDay[msg.sender] + 86400;
        }

        UDIDToken.transfer(msg.sender,amountPerDay);
    }

    function RetrieveUDID() public onlyOwner{
        UDIDToken.transfer(msg.sender,UDIDToken.balanceOf(address(this)));
    }

    function SetAmountPerDay(uint256 amount) public onlyOwner {
        amountPerDay = amount;
    }

    function LastTime() external view returns(uint256 tm){
        return (userTokenPerDay[msg.sender]);
    }

}

