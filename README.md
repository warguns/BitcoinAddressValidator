# BitcoinAddressValidator

A simple, easy to use golang Bitcoin address validator, ported from:
https://github.com/LinusU/php-bitcoin-address-validator


## Usage

Quick start:

```golang
// This will return false, indicating invalid address.
IsValid("blah");

// This is a valid address and will thus return true.
IsValid("1AGNa15ZQXAZUgFiqJ2i7Z2DPU2J6hW62i");

// This is a Testnet address, it"s valid and the function will return true.
IsValid("mo9ncXisMeAoXwqcV5EWuyncbmCcQN4rVs", TestNet);
```

## API

### `IsValid(addr, version)`

- `addr`: A bitcoin address
- `version`: The version to test against, defaults to `MAINNET`

Returns a boolean indicating if the address is valid or not.

### `TypeOf(addr)`

- `addr`: A bitcoin address

Returns a boolean if it"s a well-formed address and a string with the type of the address.

## Constants

The library exposes the following constants.

- `MAINNET`: Indicates any mainnet address type
- `TESTNET`: Indicates any testnet address type
- `MAINNET_PUBKEY`: Indicates a mainnet pay to pubkey hash address
- `MAINNET_SCRIPT`: Indicates a mainnet pay to script hash address
- `TESTNET_PUBKEY`: Indicates a testnet pay to pubkey hash address
- `TESTNET_SCRIPT`: Indicates a testnet pay to script hash address
