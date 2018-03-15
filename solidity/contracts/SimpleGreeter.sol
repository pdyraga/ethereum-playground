pragma solidity ^0.4.17;

contract SimpleGreeter {
  string greeting;
    
  function SimpleGreeter(string _greeting) public {
    greeting = _greeting;
  }
    
  function greet() public constant returns (string) {
    return greeting;
  }
}