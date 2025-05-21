# relaycheck

A lightweight HTTP API to check if the client uses
[iCloud Private Relay](https://support.apple.com/en-us/102602).

## Installation

Either download the
[latest release](https://github.com/bjoernalbers/relaycheck/releases/latest)
or built relaycheck from source:

- install go
- clone this repository and `cd` into it
- build `relaycheck` binary via `go build`

## Usage

Run `./relaycheck` to start the HTTP server, which will listen on ":8080" by
default (can be overwritten with `--addr`).
It will check the address from the HTTP header "X-Forwarded-For" to get the
original client address if behind a reverse proxy.
Otherwise the regular client address will be tested.

Response without iCloud Private Relay:

    $ curl localhost:8080
    {"relay":false}

Response with an iCloud Private Relay address:

    $ curl -H "X-Forwarded-For: 140.248.36.60" localhost:8080
    {"relay":true}
