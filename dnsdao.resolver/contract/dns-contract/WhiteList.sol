//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.5.0;

import "./BasLib.sol";
import "./Owner.sol";



contract WhiteListContract is owned{
    using SafeMath for uint256;

    uint256 mintCnt;

    mapping(address=>uint256)  userList;

    constructor() {
        mintCnt = 5;
    }

    function AddWhiteList() external {
        require(userList[msg.sender] == 0,"user is in white list");

        userList[msg.sender] = 1;
    }

    function DelWhiteList() external {
        require(userList[msg.sender] >0, "user is not in white list");

        delete(userList[msg.sender]);
    }

    function SetMintCount(uint256 cnt) external onlyOwner{
        mintCnt = cnt;
    }

    function IsInWhiteList() external view returns (bool)  {
        if (userList[msg.sender] > 0){
            return (true);
        }
        return (false);
    }

}



