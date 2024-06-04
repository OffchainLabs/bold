#!/bin/bash

anvil --fork-url $L1_RPC > /dev/null &

anvil_pid=$!

chain_id=$(cast chain-id --rpc-url $L1_RPC)

# if chain id is 1, set some env vars. if chain id is 42, set some other env vars
if [ $chain_id -eq 1 ]; then
    export EXECUTOR="0xE6841D92B0C345144506576eC13ECf5103aC7f49"
    export CONFIG_LOCATION="./scripts/files/mainnetConfig.json"
    export DEPLOYED_CONTRACTS_LOCATION="./scripts/files/mainnetDeployedContracts.json"
elif [ $chain_id -eq 11155111 ]; then
    export EXECUTOR="0x6EC62D826aDc24AeA360be9cF2647c42b9Cdb19b"
    export CONFIG_LOCATION="./scripts/files/sepoliaConfig.json"
    export DEPLOYED_CONTRACTS_LOCATION="./scripts/files/sepoliaDeployedContracts.json"
else
  echo "Unsupported chain id: $chain_id"
  exit 1
fi

yarn script:bold-prepare && \
yarn script:bold-populate-lookup && \
yarn script:bold-local-execute

kill $anvil_pid
