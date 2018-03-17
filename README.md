```
$ geth --syncmode "light" --rinkeby --rpc --unlock="0xa46603651a444307c4a36dccce939ed922a51173"
```

```
$ truffle migrate --network rinkeby
Using network 'rinkeby'.

Running migration: 1_initial_migration.js
  Deploying Migrations...
  ... 0x984f41952e35a95b5093798d6cd724af06e13177ce4e8525d02c60a341a62a51
  Migrations: 0x86ece1c62d1e3d260cd0ff5f417ef642827f0bd5
Saving successful migration to network...
  ... 0xacc9cb296b2d9bbb51b2207ae2ba0497487b169b0984bbe70d3ef054a47913fd
Saving artifacts...
Running migration: 2_deploy_simple_greeter.js
  Deploying SimpleGreeter...
  ... 0x539af9d7803939ae77c89071f18fe1622db5a21117fdd6d0e83f2229192d67af
  SimpleGreeter: 0xb57da315adcfa75ba0eedfef02886102d0df7a27
Saving successful migration to network...
  ... 0x12e6b061e5b0865fd81d0fbc03783908c1881acd32dedc9942a6bfa51c41fe09
Saving artifacts...
```
