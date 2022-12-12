#!/bin/bash

source $ROOT/scripts/network/vars.sh
relayer_config=$ROOT/scripts/network/config/relayer-config/config
relayer_logs=$ROOT/scripts/network/logs/relayer.log
relayer_exec="docker-compose run --rm relayer"

mkdir -p $relayer_config
cp $ROOT/scripts/network/relayer-config.yaml $relayer_config/config.yaml

$relayer_exec rly keys restore juno rly-juno "$MNEMONIC_3" >> $relayer_logs 2>&1
$relayer_exec rly keys restore osmosis rly-osmo "$MNEMONIC_3" >> $relayer_logs 2>&1

printf "Waiting for relayer to start..."

$relayer_exec rly transact link juno-osmosis >> $relayer_logs 2>&1

docker-compose up -d relayer
docker-compose logs -f relayer | sed -r -u "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" >> $relayer_logs 2>&1 &

echo "Done"