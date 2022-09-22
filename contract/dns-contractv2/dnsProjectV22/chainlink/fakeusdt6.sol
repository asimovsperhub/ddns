//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import "../../ERC20/ERC20.sol";

//0x8CaEB7aAA241F371B6A56f3495865A9Bf25d4EE9 rinkeby
contract FakeUSDT6 is ERC20{
    string  public constant  name = "Fake USDT6";
    string  public constant  symbol = "FUSDT6";
    uint8   public constant  decimals = 6;
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