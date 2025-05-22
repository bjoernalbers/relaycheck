# relaycheck - Simple HTTP API to detect iCloud Private Relay clients

[iCloud Private Relay](https://support.apple.com/en-us/102602) is a privacy
feature by Apple that hides a user's IP address by routing Safari web traffic
through [relay servers](https://mask-api.icloud.com/egress-ip-ranges.csv).

relaycheck detects whether a client is using an iCloud Private Relay address by
providing a simple HTTP API with a minimal JSON response:

```json
{ "relay": true }
```

It can be used in websites to let users
[verify if iCloud Private Relay is working as expected](https://www.bjoernalbers.de/tools/icloud-privat-relay-test/).

When deployed behind a reverse proxy, relaycheck evaluates the
`X-Forwarded-For` header to extract the original client IP.

## Installation

Just download the
[latest release](https://github.com/bjoernalbers/relaycheck/releases/latest)
and make it executable: `chmod +x relaycheck`
Or built `relaycheck` from source:

- install Go, i.e. by using [Homebrew](https://brew.sh) on a Mac: `brew install go`
- clone this repository and `cd` into it
- build `relaycheck` binary via `make`

## Usage

Run `./relaycheck` to start the HTTP server, which will listen on ":8080" by
default (can be overwritten with `-addr` option).

Response without iCloud Private Relay:

    $ curl localhost:8080
    {"relay":false}

Response with an iCloud Private Relay address:

    $ curl -H "X-Forwarded-For: 140.248.36.60" localhost:8080
    {"relay":true}
