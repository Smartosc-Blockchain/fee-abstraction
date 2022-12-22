#!/bin/bash

# build cargo
cargo build

# detect architecture and decide image
arch=$(uname -m)
workspace_optimizer_image="cosmwasm/workspace-optimizer:0.12.10"
if [ $arch == "arm64" ]; then
  workspace_optimizer_image="cosmwasm/workspace-optimizer-arm64:0.12.8"
fi

# compile contract
sudo docker run --rm -v "$(pwd)":/code \
  --mount type=volume,source="$(basename "$(pwd)")_cache",target=/code/target \
  --mount type=volume,source=registry_cache,target=/usr/local/cargo/registry \
  $workspace_optimizer_image