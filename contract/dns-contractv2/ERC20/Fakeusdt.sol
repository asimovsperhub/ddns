//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./ERC20.sol";

contract FakeUSDT is ERC20{
    string  public constant  name = "Fake USDT";
    string  public constant  symbol = "FUSDT";
    uint8   public constant  decimals = 18;
    uint256 public constant INITIAL_SUPPLY = 4.2e8 * (10 ** uint256(decimals));

    constructor() {
        _mint(msg.sender, INITIAL_SUPPLY);
    }

    function burn(uint amount) external{
        _burn(msg.sender, amount);
    }

    function burnFrom(address account, uint amount) external{
        _burnFrom(account, amount);
    }
}