# nameservicenel

**nameservicenel** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

Goals:

-   Create a blockchain without a default module
-   Create a Cosmos SDK nameservice module with a dependency on another module
-   Create CRUD actions for a type stored as a map
-   Declare functions of the bank module to be available to the nameservice module
-   Implement keeper functions that implement the logic

## Get started

First, clone this repository:

```
git clone https://github.com/mitchellnel/nameservice-nel
```

Then, invoke:

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Buy a new name

```
nameservice-neld tx nameservice buy-name foo 20token \
  --from alice --chain-id nameservicenel
```

This sends a transaction to buy the name `foo` for `20token`s from the account with name `alice`.

### Query the chain for a list of names

```
nameservice-neld q nameservice list-whois
```

This outputs:

```
pagination:
  next_key: null
  total: "0"
whois:
- index: foo
  name: foo
  owner: cosmos1svlkn5htc2nx568fdrk6je7utqw53gng4dmtn7
  price: 20token
  value: ""
```

We can see that there is not yet a value tied to the `foo` name.

### Set a value to the name

Now that `alice` is an owner of the name, she can set the value to anything that she wants.

Use the `set-name` command to set the value to `bar`:

```
nameservice-neld tx nameservice set-name foo bar \
  --from alice --chain-id nameservicenel
```

And now when we query the chain for a list of names, we get the output:

```
pagination:
  next_key: null
  total: "0"
whois:
- index: foo
  name: foo
  owner: cosmos1svlkn5htc2nx568fdrk6je7utqw53gng4dmtn7
  price: 20token
  value: bar
```
