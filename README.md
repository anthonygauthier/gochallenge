# Golang challenge

## Overview
This little Go program first loads an SIP (Session Initiation Protocol) registrations data file. Then it exposes the data via a REST API.

## Endpoints

```go
GET /aor/:addressOfRecord
```
This returns the following JSON data:

```json
{

    "addressOfRecord":"someAOR",
    "tenantId":"000000-0f0f-0000-0fff-000000000000",
    "uri":"sip:0000f0ff0000ff00ff000000000000@255.255.255.255;jbcuser=user",
    "contact":"<sip:0000f0ff0000ff00ff000000000000@255.255.255.255;jbcuser=user>;methods=\"INVITE, ACK, BYE, CANCEL, OPTIONS, INFO, MESSAGE, SUBSCRIBE, NOTIFY, PRACK, UPDATE, REFER\"",
    "path":["<sip:Ff0fFFffFFffFfF0FF0f0FFfFFf0FfF@127.0.0.1:3000;lr>"],
    "source":"144.144.144.144:27015",
    "target":"144.144.144.145:27016",
    "userAgent":"some.domain.000",
    "rawUserAgent":"SomeDomainFFF-FFF_000-UA/144.144.144.144",
    "created":"2016-12-12T22:40:40.764Z",
    "lineId":"000000-0f0f-0000-0fff-000000000000"
}
```