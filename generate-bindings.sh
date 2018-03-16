#!/bin/sh
rm -rf contracts
mkdir contracts contracts/abi
solc --overwrite -o contracts/abi --abi solidity/contracts/SimpleGreeter.sol
abigen --abi contracts/abi/SimpleGreeter.abi --pkg contracts --type SimpleGreeter --out contracts/SimpleGreeter.go
