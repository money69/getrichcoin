rpcclient
=========

[![Build Status](http://img.shields.io/travis/grhsuite/grhd.svg)](https://travis-ci.org/grhsuite/grhd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/grhsuite/grhd/rpcclient)

rpcclient implements a Websocket-enabled GetRichCoin JSON-RPC client package written
in [Go](http://golang.org/).  It provides a robust and easy to use client for
interfacing with a GetRichCoin RPC server that uses a grhd/getrichcoin core compatible
GetRichCoin JSON-RPC API.

## Status

This package is currently under active development.  It is already stable and
the infrastructure is complete.  However, there are still several RPCs left to
implement and the API is not stable yet.

## Documentation

* [API Reference](http://godoc.org/github.com/grhsuite/grhd/rpcclient)
* [grhd Websockets Example](https://github.com/grhsuite/grhd/tree/master/rpcclient/examples/grhdwebsockets)  
  Connects to a grhd RPC server using TLS-secured websockets, registers for
  block connected and block disconnected notifications, and gets the current
  block count
* [grhwallet Websockets Example](https://github.com/grhsuite/grhd/tree/master/rpcclient/examples/grhwalletwebsockets)  
  Connects to a grhwallet RPC server using TLS-secured websockets, registers for
  notifications about changes to account balances, and gets a list of unspent
  transaction outputs (utxos) the wallet can sign
* [GetRichCoin Core HTTP POST Example](https://github.com/grhsuite/grhd/tree/master/rpcclient/examples/getrichcoincorehttp)  
  Connects to a getrichcoin core RPC server using HTTP POST mode with TLS disabled
  and gets the current block count

## Major Features

* Supports Websockets (grhd/grhwallet) and HTTP POST mode (getrichcoin core)
* Provides callback and registration functions for grhd/grhwallet notifications
* Supports grhd extensions
* Translates to and from higher-level and easier to use Go types
* Offers a synchronous (blocking) and asynchronous API
* When running in Websockets mode (the default):
  * Automatic reconnect handling (can be disabled)
  * Outstanding commands are automatically reissued
  * Registered notifications are automatically reregistered
  * Back-off support on reconnect attempts

## Installation

```bash
$ go get -u github.com/grhsuite/grhd/rpcclient
```

## License

Package rpcclient is licensed under the [copyfree](http://copyfree.org) ISC
License.
