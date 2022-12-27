#!/bin/bash

ACCOUNT="test"
SLEEP_TIME="15"
KEYRING="test"

source scripts/vars.sh

echo ${BINARY[1]}
CONTRACT="osmo14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sq2r9g9"
ADD_ROUTE="{
    \"set_route\":
        {
            \"input_denom\":\"ibc/04F5F501207C3626A2C14BFEF654D51C2E0B8F7CA578AB8ED272A66FE4E48097\",
            \"output_denom\":\"uosmo\",
            \"pool_route\":[{\"pool_id\":\"1\",\"token_out_denom\":\"uosmo\"}]
        }
    }"
RES=$(${BINARY[1]} tx wasm execute "$CONTRACT" "${ADD_ROUTE}" --from "$ACCOUNT" -y --chain-id "${CHAINID[1]}" --node "${NODE[1]}" --gas 400000 --keyring-backend $KEYRING --home ${DIR[$i]} --fees 0${DENOM[1]} -o json)
echo $RES

TXHASH=$(echo $RES | jq -r .txhash)

echo $TXHASH

# sleep for chain to update
sleep "$SLEEP_TIME"

RAW_LOG=$(${BINARY[1]} query tx "$TXHASH" --chain-id "${CHAINID[1]}" --node "${NODE[1]}" -o json | jq -r .raw_log)

echo $RAW_LOG

# query route
GET_ROUTE="{
    \"get_route\": 
        {
            \"input_denom\":\"ibc/04F5F501207C3626A2C14BFEF654D51C2E0B8F7CA578AB8ED272A66FE4E48097\",
            \"output_denom\":\"uosmo\"
        }
    }"
RES=$(${BINARY[1]} query wasm contract-state smart "$CONTRACT" "$GET_ROUTE" --node "${NODE[1]}" --output json)
echo $RES
