#!/usr/bin/env bash

contract_path="./on-chain/evm-contracts/contracts"
test_contracts="SimpleEmitter"
contracts="Receiver BridgeAsset Emitter Safe"

echo "Compiling contracts..."

pushd $contract_path

# npm install --silent
# node_modules/.bin/truffle compile

rm -r ../bindings
mkdir ../bindings

for value in $test_contracts
do
	solcjs --bin -o bin/ tests/${value}.sol 
	solcjs --abi -o abi/ tests/${value}.sol 
	abigen --abi abi/tests_${value}_sol_${value}.abi --pkg $value --type $value --bin bin/tests_${value}_sol_${value}.bin --out ../bindings/$value.go
    echo "finished generating bindings for $value at $contract_path/bindings/$value.go"
done

echo "@@@@@@@@@@@@start receiver"
solcjs --abi -o abi/ Receiver.sol BridgeAsset.sol interfaces/IHandler.sol helpers/Ownable.sol 
solcjs --bin -o bin/ Receiver.sol BridgeAsset.sol interfaces/IHandler.sol helpers/Ownable.sol 
echo "@@@@@@@@@@@@done receiver"
echo "@@@@@@@@@@@@start emitter"
solcjs --abi -o abi/ Emitter.sol Safe.sol erc/ERC20/ERC20.sol interfaces/ISafe.sol helpers/SafeMath.sol erc/ERC20/IERC20.sol erc/ERC721/IERC721.sol erc/ERC721/ERC721.sol erc/ERC721/IERC721Receiver.sol helpers/SafeMath.sol helpers/Address.sol helpers/Counters.sol erc/ERC165/ERC165.sol erc/ERC165/IERC165.sol
solcjs --bin -o bin/ Emitter.sol Safe.sol erc/ERC20/ERC20.sol interfaces/ISafe.sol helpers/SafeMath.sol erc/ERC20/IERC20.sol erc/ERC721/IERC721.sol erc/ERC721/ERC721.sol erc/ERC721/IERC721Receiver.sol helpers/SafeMath.sol helpers/Address.sol helpers/Counters.sol erc/ERC165/ERC165.sol erc/ERC165/IERC165.sol
echo "@@@@@@@@@@@@done emitter"

for value in $contracts
do
	abigen --abi abi/${value}_sol_${value}.abi --pkg $value --type $value --bin bin/${value}_sol_${value}.bin --out ../bindings/$value.go
	echo "finished generating bindings for $value at $contract_path/bindings/$value.go"
done

popd

mkdir ./contracts

for value in $test_contracts
do
	rm -r ./contracts/$value
	mkdir "./contracts/$value"
	mv "./on-chain/evm-contracts/bindings/$value.go" "./contracts/$value"
	echo "moved $value.go to go package ./contracts/$value"
done

for value in $contracts
do
	rm -r ./contracts/$value
	mkdir "./contracts/$value"
	mv "./on-chain/evm-contracts/bindings/$value.go" "./contracts/$value"
	echo "moved $value.go to go package ./contracts/$value"
done