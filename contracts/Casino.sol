// SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

import './ERC721A.sol';
import './CasinoToken.sol';


contract Casino is ERC721A {

      CasinoToken public casinoToken; 
      uint256 private constant _PRICE = 1e18;
      uint256 private constant quantity = 3;
      address private owner;
      constructor() ERC721A("Casino", "CASINO") {
          owner = msg.sender;
      }
    
    modifier onlyOwner {
        require( msg.sender == owner, "not owner");
        _;
    }
       
    function mint() external payable {
        // `_mint`'s second argument now takes in a `quantity`, not a `tokenId`.
        require(_tokenBalance() >= _PRICE, "not enough");
        casinoToken.transferFrom(msg.sender, address(this), _PRICE);
        _mint(msg.sender, quantity);
    }

    function _tokenBalance() internal view returns(uint bal) {
       //(bool success, bytes memory data ) = casinoToken.call(abi.encodeWithSignature("balanceOf(address)",msg.sender));
       bal = casinoToken.balanceOf(msg.sender);
    }
    function handleToken(address tokenAddr) external onlyOwner {
        casinoToken = CasinoToken(tokenAddr);
    }
}