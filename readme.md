# Ethereum Storage
This is a Cosmos based blockchain that stores Ethereum state.

## Configuration
To fetch the Ethereum state, you need to configure the Ethereum node RPC endpoint as the environment variable `ETH_RPC_URL`.

The default parameters are set like this:
```go
// NewParams creates a new Params instance
func NewParams() Params {
	return Params{
		Addresses: []string{
			"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", // mainnet WETH
			"0xB753548F6E010e7e680BA186F9Ca1BdAB2E90cf2", // mainnet Uniswap Proxy Admin
		},
		FromBlock: 17380596, // indicates the block number to start fetching the state
		MaxCount:  2, // indicates the max number of slots to fetch at once
	}
}
```

## Run a local testnet
```bash
# build the binary
make build 

# initialize the chain
make init-chain

# start the chain
make run-chain
```

## Cliend commands
```bash
# show-storage [address] [slot] [block] 
./simulate/eth-storaged query ethstorage show-storage 0xB753548F6E010e7e680BA186F9Ca1BdAB2E90cf2 0 17380596 --output json

# show-storage [address] [slot] :: the block number is set to the latest block 
./simulate/eth-storaged query ethstorage show-storage 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2 0 --output json
```