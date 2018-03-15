#!/bin/sh
rm -rf target
mkdir target target/abi target/go
solc --overwrite -o target/abi --abi ../solidity/contracts/SimpleGreeter.sol
abigen --abi target/abi/SimpleGreeter.abi --pkg main --type SimpleGreeter --out target/go/SimpleGreeter.go
