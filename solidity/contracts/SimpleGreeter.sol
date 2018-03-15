pragma solidity ^0.4.17;

contract Greeter {
  string greeting;
    
  function Greeter(string _greeting) public {
    greeting = _greeting;
  }
    
  function greet() public constant returns (string) {
    return greeting;
  }
}