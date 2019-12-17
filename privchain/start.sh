#!/bin/bash

geth \
  --datadir ./data \
  --networkid 15 \
  --rpc \
  --rpccorsdomain "*" \
  --rpcapi "eth,web3,,personal,miner,net,txpool" \
  --allow-insecure-unlock \
  --mine \
  --minerthreads 1 \
  --etherbase 'a370bfe70687e3d2f17d4723d7d5bfec8806a63b' \
  --shh --ws
