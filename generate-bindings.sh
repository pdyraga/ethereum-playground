#!/bin/sh
rm -rf contracts
mkdir contracts contracts/abi
solc --overwrite -o contracts/abi --abi solidity/contracts/SimpleGreeter.sol
solc --overwrite -o contracts/abi --abi solidity/contracts/EchoGreeter.sol
abigen --abi contracts/abi/SimpleGreeter.abi --pkg contracts --type SimpleGreeter --out contracts/SimpleGreeter.go
abigen --abi contracts/abi/EchoGreeter.abi --pkg contracts --type EchoGreeter --out contracts/EchoGreeter.go