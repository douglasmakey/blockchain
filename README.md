
## Blockchain with Go

Simple example to understand and learn about Blockchain. 

```bash
go build
```

Next
```bash
$ ./blockchain history
Mining the block with data "GenesisBlock"
00002bcb38688ce4b1ca7b4b7dd28a079d5de93c01e8fae26e326f8b1c983fc3

Prev. hash:
Data: GenesisBlock
Hash: 00002bcb38688ce4b1ca7b4b7dd28a079d5de93c01e8fae26e326f8b1c983fc3
PoW Validation: true

$ ./blockchain add -data "Pay 1 BTC for a car"
Mining the block with data "Pay 1 BTC for a car"
0000613f831bf20420fe5465405b23631d5422514161d8fa9db4b1abe69d6d1b

Done!

$ ./blockchain history
Prev. hash: 00002bcb38688ce4b1ca7b4b7dd28a079d5de93c01e8fae26e326f8b1c983fc3
Data: Pay 1 BTC for a car
Hash: 0000613f831bf20420fe5465405b23631d5422514161d8fa9db4b1abe69d6d1b
PoW Validation: true

Prev. hash:
Data: GenesisBlock
Hash: 00002bcb38688ce4b1ca7b4b7dd28a079d5de93c01e8fae26e326f8b1c983fc3
PoW Validation: true

```
