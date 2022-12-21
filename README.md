# Fee abstraction Contracts

Enable fee abstraction for Cosmos app chains.

## Description

Fee abstraction allows users on any cosmos app-chain to pay gas fees in the supported IBC asset of their choice, by using the Osmosis AMM to facilitate the swap in the background.

For example, even if a user doesn't have JUNO, they could still submit a tx on Juno Network by selecting another asset such as USDC to pay gas fees with.

check build-and-start-network.sh to start

## Note for starting
1. pip install toml-cli
2. turn off sudo for docker: https://github.com/sindresorhus/guides/blob/main/docker-without-sudo.md