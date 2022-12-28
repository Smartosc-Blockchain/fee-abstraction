#!/bin/bash

ACCOUNT="test"
SLEEP_TIME="20"
KEYRING="test"

source scripts/vars.sh

# check if CROSSCHAIN_CONTRACT is empty env var
if [ $CROSSCHAIN_CONTRACT = "" ]; then
    echo "run scripts/deploy_contract.sh first"
    exit 1
fi

HOST_ACCOUNT=$(${BINARY[0]} keys show $ACCOUNT -a --keyring-backend $KEYRING --home ${DIR[0]})
DESTINATION_ACCOUNT=$(${BINARY[1]} keys show $ACCOUNT -a --keyring-backend $KEYRING --home ${DIR[1]})

echo "Check balance before IBC transfer and crosschain swap"

# check juno balance
BALANCE=$(${BINARY[0]} query bank balances "$HOST_ACCOUNT" --chain-id "${CHAINID[0]}" --node "${NODE[0]}" --output json | jq -r .balances[])
echo $BALANCE
# check osmosis balance
BALANCE=$(${BINARY[1]} query bank balances "$DESTINATION_ACCOUNT" --chain-id "${CHAINID[1]}" --node "${NODE[1]}" --output json | jq -r .balances[])
echo $BALANCE

echo "Transfer..."
# transfer some ujuno to osmosis to swap after
TRANSFER=$(${BINARY[0]} tx ibc-transfer transfer transfer channel-0 $DESTINATION_ACCOUNT 1000000ujuno --from $ACCOUNT --keyring-backend $KEYRING --home ${DIR[0]} --chain-id "${CHAINID[0]}" -y --node ${NODE[0]})

sleep $SLEEP_TIME

echo "Check balance after IBC transfer and before crosschain swap"

# check juno balance after transfer
BALANCE=$(${BINARY[0]} query bank balances "$HOST_ACCOUNT" --chain-id "${CHAINID[0]}" --node "${NODE[0]}" --output json | jq -r .balances[])
echo $BALANCE
# check osmosis balance after transfer
BALANCE=$(${BINARY[1]} query bank balances "$DESTINATION_ACCOUNT" --chain-id "${CHAINID[1]}" --node "${NODE[1]}" --output json | jq -r .balances[])
echo $BALANCE

echo "Crosschain swap..."
# exec crosschain swap 
MEMO=$(jq -c --null-input --arg CROSSCHAIN_CONTRACT "$CROSSCHAIN_CONTRACT" --arg HOST_ACCOUNT "$HOST_ACCOUNT" '{
    "wasm":{
        "contract":$CROSSCHAIN_CONTRACT, 
        "msg": {
            "osmosis_swap": {
                "input_coin": {
                    "denom":"ibc/04F5F501207C3626A2C14BFEF654D51C2E0B8F7CA578AB8ED272A66FE4E48097",
                    "amount":"100000"
                }, 
                "output_denom":"uosmo", 
                "slippage":{
                    "max_slippage_percentage":"20"
                }, 
                "receiver":$HOST_ACCOUNT
            }
        }
    }
}')
RES=$(${BINARY[0]} tx ibc-transfer transfer transfer channel-0 $DESTINATION_ACCOUNT 1ujuno --from $ACCOUNT --keyring-backend $KEYRING --home ${DIR[0]} --chain-id "${CHAINID[0]}" -y --node ${NODE[0]} --memo $MEMO -o json)
echo $RES

sleep $SLEEP_TIME

echo "Check balance after crosschain swap"
# check juno balance after swap
BALANCE=$(${BINARY[0]} query bank balances "$HOST_ACCOUNT" --chain-id "${CHAINID[0]}" --node "${NODE[0]}" --output json | jq -r .balances[])
echo $BALANCE
# check osmosis balance after swap
BALANCE=$(${BINARY[1]} query bank balances "$DESTINATION_ACCOUNT" --chain-id "${CHAINID[1]}" --node "${NODE[1]}" --output json | jq -r .balances[])
echo $BALANCE