#!/usr/bin/env bash

set -e

BIN=./simulate/eth-storaged

rm -rf simulate/keyring-test
rm -rf simulate/config
rm -rf simulate/data

$BIN config chain-id demo --home ./simulate
$BIN config keyring-backend test --home ./simulate
$BIN keys add alice --home ./simulate
$BIN keys add bob --home ./simulate
$BIN init test --chain-id demo --home ./simulate
$BIN add-genesis-account alice 5000000000stake --keyring-backend test --home ./simulate
$BIN add-genesis-account bob 5000000000stake --keyring-backend test --home ./simulate
$BIN gentx alice 1000000stake --chain-id demo --home ./simulate
$BIN collect-gentxs --home ./simulate

# $BIN start --home ./simulate
