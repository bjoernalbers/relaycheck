# relaycheck - Simple HTTP API to detect iCloud Private Relay clients

[iCloud Private Relay](https://support.apple.com/en-us/102602) is a privacy
feature by Apple that hides a user's IP address by routing Safari web traffic
through relay servers.

relaycheck detects whether a client is using an
[iCloud Private Relay address](https://mask-api.icloud.com/egress-ip-ranges.csv)
and returns the result in a minimal JSON response:

```json
{
   "ip" : "172.225.6.92",
   "location" : {
      "city" : "Berlin",
      "country_code" : "DE",
      "region_code" : "DE-BE"
   },
   "relay" : true
}
```

It can be used in websites to let users
[verify if iCloud Private Relay is working as expected](https://www.bjoernalbers.de/icloud-privat-relay-test/).

When deployed behind a reverse proxy, relaycheck evaluates the
`X-Forwarded-For` header to extract the original client IP.

## Installation

Just download the
[latest release](https://github.com/bjoernalbers/relaycheck/releases/latest)
and make it executable with `chmod +x relaycheck`.
Mac users will have to remote the quarantine attribute as well:

```
$ xattr -r -d com.apple.quarantine ~/Downloads/relaycheck-darwin-arm64
```

Or built `relaycheck` from source:

- install Go, i.e. by using [Homebrew](https://brew.sh) on a Mac: `brew install go`
- clone this repository and `cd` into it
- build `relaycheck` binary via `make`

## Usage

Run `./relaycheck` to start the HTTP server, which will listen on ":8080" by
default (can be overwritten with `-addr` option).

Response for a regular (non-relay) address:

```
$ curl -sH "X-Forwarded-For: 1.1.1.1" localhost:8080 | json_pp
{
   "ip" : "1.1.1.1",
   "relay" : false
}
```

Response for an iCloud Private Relay address:

```
$ curl -sH "X-Forwarded-For: 172.225.6.92" localhost:8080 | json_pp
{
   "ip" : "172.225.6.92",
   "location" : {
      "city" : "Berlin",
      "country_code" : "DE",
      "region_code" : "DE-BE"
   },
   "relay" : true
}
```
