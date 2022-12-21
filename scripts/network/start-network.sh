#!/bin/bash

source scripts/vars.sh
cd scripts/network

start_docker() {
    name=$1

    docker-compose up -d $name
    docker-compose logs -f $name | sed -r -u "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" > logs/$name.log 2>&1 &

    printf "Waiting for $name to start..."

    ( tail -f -n0 logs/$name.log & ) | grep -q "finalizing commit of block"
    echo "Done"
}

# cleanup any stale state
docker-compose down
rm -rf config logs
mkdir logs

# init chain
bash init-chain.sh $ROOT/build/binary/junod ujuno $ROOT/scripts/network/config/juno
bash init-chain.sh $ROOT/build/binary/osmosisd uosmo $ROOT/scripts/network/config/osmosis

# start docker
start_docker juno
start_docker osmosis

# start relayer
bash setup-relayer.sh

# setup pool on osmosis
${BINARY[0]} tx ibc-transfer transfer transfer channel-0 $OSMO_2 1000000000ujuno --from test1 --keyring-backend test --home ${DIR[0]} --chain-id test-juno --fees 100000ujuno --yes --node ${NODE[0]}

sleep 10

# create pool
IBC_DENOM=$(${BINARY[1]} q ibc-transfer denom-hash transfer/channel-0/ujuno --node ${NODE[1]} -o json | jq -r .hash)
sed "s/IBCDENOM/$IBC_DENOM/g" $ROOT/scripts/network/juno-osmosis-pool.json > $ROOT/scripts/network/config/juno-osmosis-pool.json

${BINARY[1]} tx gamm create-pool --pool-file $ROOT/scripts/network/config/juno-osmosis-pool.json --from test1 --keyring-backend test --home ${DIR[1]} --chain-id test-osmo --fees 100000uosmo --yes --node ${NODE[1]}

sleep 10

POOL_NUM=$(${BINARY[1]} q gamm pools --node tcp://localhost:26357 -o json | jq '.pools | length')
echo "Pool number: $POOL_NUM"

cd ../..