#!/usr/bin/env bash

contract_path="./contracts/evm-contracts/contracts/"
test_contracts="SimpleEmitter"
contracts="Receiver BridgeAsset Emitter Safe"

echo "Compiling contracts..."

pushd $contract_path

# npm install --silent
# node_modules/.bin/truffle compile

for value in $test_contracts
do
	solcjs --bin -o bin/ tests/${value}.sol 
	solcjs --abi -o abi/ tests/${value}.sol 
	abigen --abi abi/tests_${value}_sol_${value}.abi --pkg $value --type $value --bin bin/tests_${value}_sol_${value}.bin --out ../bindings/$value.go
    echo "finished generating bindings for $value at $contract_path/bindings/$value.go"
done

solcjs --abi -o abi/ Receiver.sol BridgeAsset.sol interfaces/IHandler.sol helpers/Ownable.sol 
solcjs --bin -o bin/ Receiver.sol BridgeAsset.sol interfaces/IHandler.sol helpers/Ownable.sol 
# abigen --abi abi/Receiver_sol_Receiver.abi --pkg $value --type $value --bin bin/contracts_${value}_sol_${value}.bin --out ../bindings/$value.go
# echo "finished generating bindings for Receiver, BridgeAsset at $contract_path/bindings/"

solcjs --abi -o abi/ Emitter.sol Safe.sol erc/ERC20.sol interfaces/ISafe.sol helpers/SafeMath.sol erc/IERC20.sol 
solcjs --bin -o bin/ Emitter.sol Safe.sol erc/ERC20.sol interfaces/ISafe.sol helpers/SafeMath.sol erc/IERC20.sol 
for value in $contracts
do
	abigen --abi abi/${value}_sol_${value}.abi --pkg $value --type $value --bin bin/${value}_sol_${value}.bin --out ../bindings/$value.go
	echo "finished generating bindings for $value at $contract_path/bindings/$value.go"
done

popd