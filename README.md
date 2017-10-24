# Realex HPP Golang SDK
You can sign up for a Realex account at https://www.realexpayments.com.
## Requirements
Golang 1.7+
## Installation
```sh
$ go get github.com/Fatsoma/rxp-hpp-go
```

Or using dep dependency manager
```sh
dep ensure -add github.com/Fatsoma/rxp-hpp-go
```
You then import it with this import path:

```go
import hpp "github.com/Fatsoma/rxp-hpp-go"
```

## Usage
### Creating Request JSON for Realex JS SDK
```golang

req := hpp.Request{
  req.Amount = 100
  req.Currency = "EUR"
  req.MerchantID = "merchantID"
}
json, err := hpp.New("secret").ToJSON(req)
if err != nil {
    // make request with built JSON
}

```
### Consuming Response JSON from Realex JS SDK
```golang
resp, err := hpp.New("secret").FromJSON(json)
```
## License
See the LICENSE file.
