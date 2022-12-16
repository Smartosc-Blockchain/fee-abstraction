# Design for fee-abstraction MVP

[fee-abstraction-design](docs/fee_abstraction.drawio.png)

the design above can be split into two smart contracts:
1. fee-abstraction (Juno)
2. cross-swap (Osmosis)

these two smart contracts will be supported by:
1. Juno localnet (docker)
2. Osmosis localnet (docker)
3. Deployment script
4. Relayer

## I. Fee Abstraction
The following feature is needed:
1. Tx handler:  
    * handling signed tx from user, and extract fee
    * substitute fee and submit it to chain
2. Swap handler: IBC component to make a cross - swap
3. Error handler: not in scope for this MVP

## II. Cross Swap
The following feature is needed:
1. Swap handler: IBC component to make the swap on Osmosis

Notice: May have to roll up several tx together