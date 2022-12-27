#!/bin/bash

ACCOUNT="test"
SLEEP_TIME="15"
KEYRING="test"

source scripts/vars.sh

CONTRACT_DIR=( "artifacts/cw_fee_abstraction.wasm" "artifacts/swaprouter.wasm" "artifacts/crosschain_swaps.wasm" )
CONTRACT_NAME=( "cw-fee-abstraction" "swaprouter" "crosschain_swaps" )
CONTRACT_CHAIN=( 0 1 1 )
CONTRACT_ADDRESS=( '' '' '' )
arch=$(uname -m)
if [ $arch == "arm64" ]; then
    CONTRACT_DIR=( "artifacts/cw_fee_abstraction-aarch64.wasm" "artifacts/swaprouter-aarch64.wasm" "artifacts/crosschain_swaps-aarch64.wasm" )
fi
swaprouter_init_msg=$(jq --null-input --arg OWNER $OSMO_1 '{"owner": $OWNER}')
INIT_MSG=( '' "$swaprouter_init_msg" '' )
custom_init_msg() {
    if [ "${CONTRACT_ADDRESS[1]}" != "" ]; then
        INIT_MSG[2]=$(jq --null-input --arg SWAPROUTER_CONTRACT "${CONTRACT_ADDRESS[1]}" '{"swap_contract": $SWAPROUTER_CONTRACT, "track_ibc_sends": false, "channels": [["juno", "channel-0"]]'})
    fi
}

# check if a folder exists
if [ ! -d scripts/contract-interaction/logs ]; then
    mkdir scripts/contract-interaction/logs
fi

# loop contracts
for j in $(seq 0 $((${#CONTRACT_DIR[@]} - 1))); do
    i=${CONTRACT_CHAIN[$j]}
    echo "DEPLOYING ${CONTRACT_DIR[$j]}"
    custom_init_msg
    if [ "${INIT_MSG[$j]}" = "" ]; then
        continue
    fi

    # store contract
    RES=$(${BINARY[$i]} tx wasm store "${CONTRACT_DIR[$j]}" --from "$ACCOUNT" -y --output json --chain-id "${CHAINID[$i]}" --node "${NODE[$i]}" --gas 20000000 --fees 875000${DENOM[$i]} -y --output json --keyring-backend $KEYRING --home ${DIR[$i]})
    echo $RES

    if [ "$(echo $RES | jq -r .raw_log)" != "[]" ]; then
        # exit
        echo "ERROR = $(echo $RES | jq .raw_log)"
        exit 1
    else
        echo "STORE SUCCESS"
    fi

    TXHASH=$(echo $RES | jq -r .txhash)

    echo $TXHASH

    # sleep for chain to update
    sleep "$SLEEP_TIME"

    # query code id
    RAW_LOG=$(${BINARY[$i]} query tx "$TXHASH" --chain-id "${CHAINID[$i]}" --node "${NODE[$i]}" -o json | jq -r .raw_log)

    echo $RAW_LOG

    CODE_ID=$(echo $RAW_LOG | jq -r .[0].events[1].attributes[1].value)

    echo "CODE_ID on ${CHAINID[$i]} = $CODE_ID"

    # instantiate contract
    RES=$(${BINARY[$i]} tx wasm instantiate "$CODE_ID" "${INIT_MSG[$j]}" --from "$ACCOUNT" --no-admin --label "contract" -y --chain-id "${CHAINID[$i]}" --node "${NODE[$i]}" --gas 20000000 --fees 100000${DENOM[$i]} -o json --keyring-backend $KEYRING --home ${DIR[$i]})

    if [ "$(echo $RES | jq -r .raw_log)" != "[]" ]; then
        # exit
        echo "ERROR = $(echo $RES | jq .raw_log)"
        exit 1
    else
        echo "INSTANTIATE SUCCESS"
    fi

    # sleep for chain to update
    sleep "$SLEEP_TIME"

    # query contract address
    RAW_LOG=$(${BINARY[$i]} query tx "$(echo $RES | jq -r .txhash)" --chain-id "${CHAINID[$i]}" --node "${NODE[$i]}" -o json | jq -r .raw_log)
    echo $RAW_LOG
    ADDRESS=$(echo $RAW_LOG | jq -r .[0].events[0].attributes[0].value)
    CONTRACT_ADDRESS[$j]=$ADDRESS
    echo "CONTRACT ADDRESS of ${CONTRACT_NAME[$j]} on ${CHAINID[$i]} with address = $ADDRESS" >> scripts/contract-interaction/logs/contract-addresses.txt

    echo "DONE DEPLOYING ${CONTRACT_DIR[$j]}"
    echo
    echo
    echo
done