#!/bin/bash

export ROOT=$(pwd)

export MNEMONIC_1="dry repeat crush category laugh proud pretty crew record crash neglect road valley soon solution flat poet fantasy space resist april owner ship business"
export MNEMONIC_2="supreme era pool truth shop essay source wall steel rely local wing convince enact champion warm food grunt siege obey kiss crane squeeze original"
export MNEMONIC_3="confirm select whale obtain toe fortune wisdom truck hospital cement spring when idea cupboard machine glory mouse kitchen moral fiber bomb rabbit fog raven"

export JUNO_1=juno1xrj7hrjg86fdd9ct7j4dluusgd6geghh52ff4j
export JUNO_2=juno15hh4c5dzwdy6alx6uzc5c2hzd3eu2dn22whs0p

export OSMO_1=osmo1xrj7hrjg86fdd9ct7j4dluusgd6geghh2rezyu
export OSMO_2=osmo15hh4c5dzwdy6alx6uzc5c2hzd3eu2dn2588m70

# Juno, Osmosis
export BINARY=( "$ROOT/build/binary/junod" "$ROOT/build/binary/osmosisd" )
export DIR=( "$ROOT/scripts/network/config/juno" "$ROOT/scripts/network/config/osmosis" )
export DENOM=( "ujuno" "uosmo" )
export NODE=( "http://localhost:26657" "http://localhost:26357" )
export CHAINID=( "test-juno" "test-osmo" )

# export cross-chain contract addresses
export SWAPROUTER_CONTRACT=osmo1eyfccmjm6732k7wp4p6gdjwhxjwsvje44j0hfx8nkgrm8fs7vqfsn92ayh
export CROSSCHAIN_CONTRACT=osmo1pvrwmjuusn9wh34j7y520g8gumuy9xtl3gvprlljfdpwju3x7ucsxrqwu2