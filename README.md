# relaycheck - Simple HTTP API to detect iCloud Private Relay clients

[iCloud Private Relay](https://support.apple.com/en-us/102602) is a privacy
feature by Apple that hides a user's IP address by routing Safari web traffic
through [relay servers](https://mask-api.icloud.com/egress-ip-ranges.csv).

relaycheck lets you detect whether a client is using a iCloud Private Relay
address by providing a simple HTTP API with a clear JSON response.

```json
{ "relay": true }
```

Designed for developers and end users alike, relaycheck can be used in websites
to let users verify if iCloud Private Relay is working as expected.

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
