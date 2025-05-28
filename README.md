# relaycheck - Simple HTTP API to detect iCloud Private Relay clients

[iCloud Private Relay](https://support.apple.com/en-us/102602) is a privacy
feature by Apple that hides a user's IP address by routing Safari web traffic
through relay servers.

relaycheck detects whether a client is using an
[iCloud Private Relay address](https://mask-api.icloud.com/egress-ip-ranges.csv)
and returns the result in a minimal JSON response.

Response for a regular (non-relay) address:

```json
{
   "ip" : "1.1.1.1",
   "relay" : false
}
```

Response for an iCloud Private Relay address:

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

### Download binary

Just download the
[latest release](https://github.com/bjoernalbers/relaycheck/releases/latest)
and make it executable.
Example for a Mac with Apple Silicon:

```
curl -LO https://github.com/bjoernalbers/relaycheck/releases/latest/download/relaycheck-darwin-arm64
chmod +x relaycheck-darwin-arm64
xattr -r -d com.apple.quarantine relaycheck-darwin-arm64 # remove from quarantine (only required on macOS)
```

### Build from source

This requires Go, which can be installed on a Mac by using
[Homebrew](https://brew.sh): `brew install go`

The steps for building the binary are:

```
git clone https://github.com/bjoernalbers/relaycheck.git
cd relaycheck
make
```

### Docker

There is a Docker Image as well:

```
docker run --rm ghcr.io/bjoernalbers/relaycheck
```

## Usage

Run `./relaycheck` to start the HTTP server, which will listen on ":8080" by
default (can be overwritten with `-addr` option).

Then send a test request:

```
curl -H "X-Forwarded-For: 172.225.6.92" localhost:8080
```
