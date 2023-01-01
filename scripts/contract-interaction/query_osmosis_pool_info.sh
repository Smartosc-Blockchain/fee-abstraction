#!/bin/bash

ACCOUNT="test"
SLEEP_TIME="15"
KEYRING="test"

source scripts/vars.sh

echo ${BINARY[1]}

OSMOSIS_POOL=$(${BINARY[1]} q gamm pools --node tcp://localhost:26357 -o json | jq '.pools')
echo $OSMOSIS_POOL

IBC_DENOM_OF_POOL_1=$(${BINARY[1]} q gamm pools --node tcp://localhost:26357 -o json | jq '.pools[0].pool_assets[0].token.denom')
echo $IBC_DENOM_OF_POOL_1