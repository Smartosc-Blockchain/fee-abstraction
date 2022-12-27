#!/bin/bash

ACCOUNT="test"
SLEEP_TIME="15"
KEYRING="test"

source scripts/vars.sh

CROSSCHAIN_CONTRACT="osmo1nc5tatafv6eyq7llkr2gv50ff9e22mnf70qgjlv737ktmt4eswrqvlx82r"

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
TRANSFER=$(${BINARY[0]} tx ibc-transfer transfer transfer channel-0 $DESTINATION_ACCOUNT 1000000ujuno --from $ACCOUNT --keyring-backend $KEYRING --home ${DIR[0]} --chain-id test-juno -y --node ${NODE[0]})

sleep 15

echo "Check balance after IBC transfer and before crosschain swap"

# check juno balance after transfer
BALANCE=$(${BINARY[0]} query bank balances "$HOST_ACCOUNT" --chain-id "${CHAINID[0]}" --node "${NODE[0]}" --output json | jq -r .balances[])
echo $BALANCE
# check osmosis balance after transfer
BALANCE=$(${BINARY[1]} query bank balances "$DESTINATION_ACCOUNT" --chain-id "${CHAINID[1]}" --node "${NODE[1]}" --output json | jq -r .balances[])
echo $BALANCE

echo "Crosschain swap..."
# exec crosschain swap 
RES=$(${BINARY[0]} tx ibc-transfer transfer transfer channel-0 $DESTINATION_ACCOUNT 1ujuno --from $ACCOUNT --keyring-backend $KEYRING --home ${DIR[0]} --chain-id test-juno -y --node ${NODE[0]} --memo '{"wasm":{"contract": "osmo1nc5tatafv6eyq7llkr2gv50ff9e22mnf70qgjlv737ktmt4eswrqvlx82r", "msg": {"osmosis_swap": {"input_coin":{"denom":"ibc/04F5F501207C3626A2C14BFEF654D51C2E0B8F7CA578AB8ED272A66FE4E48097", "amount":"100000"}, "output_denom":"uosmo", "slippage":{"max_slippage_percentage":"20"}, "receiver":"juno1xrj7hrjg86fdd9ct7j4dluusgd6geghh52ff4j"}}}}')

sleep 15

echo "Check balance after crosschain swap"
# check juno balance after swap
BALANCE=$(${BINARY[0]} query bank balances "$HOST_ACCOUNT" --chain-id "${CHAINID[0]}" --node "${NODE[0]}" --output json | jq -r .balances[])
echo $BALANCE
# check osmosis balance after swap
BALANCE=$(${BINARY[1]} query bank balances "$DESTINATION_ACCOUNT" --chain-id "${CHAINID[1]}" --node "${NODE[1]}" --output json | jq -r .balances[])
echo $BALANCE