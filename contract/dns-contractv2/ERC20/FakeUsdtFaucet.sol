//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "./IERC20.sol";
import "../utils/Owner.sol";


contract FakeUSDTFaucet is owned {
    IERC20 public fusdt;
    uint256 defaultU = 1000*1e18;
    uint256 defaultE = 1e16;

    constructor(address fusdtAddr_){
        fusdt = IERC20(fusdtAddr_);
    }

    function setDefaultU(uint256 defaultU_) external onlyOwner{
        defaultU = defaultU_;
    }

    function setDefaultE(uint256 defaultE_) external onlyOwner{
        defaultE = defaultE_;
    }

    function GetFreeFUSDT() external{
        require(fusdt.balanceOf(msg.sender) <  defaultU,"fusdt enough");

        fusdt.transfer(msg.sender,defaultU);

    }

    function GetFreeEth() external{
        require(msg.sender.balance < defaultE,"eth enough");
        payable(msg.sender).transfer(defaultE);
    }

    function retriveAllUsdt() external onlyOwner{
        require(fusdt.balanceOf(address(this))>0,"no usdt");
        fusdt.transfer(msg.sender,fusdt.balanceOf(address(this)));

    }

    function retriveAllEth() external onlyOwner{
        require(address(this).balance>0,"no eth");
        payable(msg.sender).transfer(address(this).balance);
    }

    receive()  payable external {}
    fallback() payable external{}
}
