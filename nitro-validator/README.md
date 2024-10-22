## Nitro Validator

This directory is used to build the nitro validator. It currently needs both solidity ^0.8.24 and to be compiled with via-ir and therefore cant be in the same src directory.

This should be a singleton deployed on the L1. 


### Build

```shell
$ forge build
```

### Test

```shell 
$ forge test
```

### Format

```shell
$ forge fmt
```

### Deploy

```shell
$ forge script script/DeployNitroValidator.s.sol:DeployNitroValidator --rpc-url <your_rpc_url>
```


### Cast

```shell
$ cast <subcommand>
```

### Help

```shell
$ forge --help
$ anvil --help
$ cast --help
```
