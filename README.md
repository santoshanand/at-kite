# The Kite Connect API Go client

NOTE: All code taken from [Zerodha] [github.com/zerodha/gokiteconnect]
The NON official Go client for communicating with the Kite Connect API.

Kite Connect is a set of REST-like APIs that expose many capabilities required
to build a complete investment and trading platform. Execute orders in real
time, manage user portfolio, stream live market data (WebSockets), and more,
with the simple HTTP API collection.

## Installation

```
go get github.com/santoshanand/at-kite
```

## API usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/santoshanand/at-kite/kite"
)

func main() {
	accessToken := ""
	// Create a new Kite connect instance
	kc := kite.New(accessToken)
	kc.SetDebug(true)
	// kc.SetAccessToken(accessToken)

	// Get margins
	margins, err := kc.GetUserMargins()
	if err != nil {
		fmt.Printf("Error getting margins: %v", err)
	}
	fmt.Println("margins: ", margins)

	dt := time.Now()

	to := dt.AddDate(0, 0, -1)
	rr, errr := kc.GetHistoricalData(13088002, "5minute", to, dt, false, true)
	if errr != nil {
		fmt.Println(errr)
	}

	fmt.Println(rr)
}

```

## Kite ticker usage

```go
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/santoshanand/at-kite/kite"
	"github.com/santoshanand/at-kite/models"
	"github.com/santoshanand/at-kite/realtime"
)

var (
	// apiSecret string = getEnv("KITE_API_SECRET", "my_api_secret")
	instToken uint32 = getEnvUint32("KITE_INSTRUMENT_TOKEN", 62285063)
)

var (
	ticker *realtime.Ticker
)

// Triggered when any error is raised
func onError(err error) {
	fmt.Println("Error: ", err)
}

// Triggered when websocket connection is closed
func onClose(code int, reason string) {
	fmt.Println("Close: ", code, reason)
}

// Triggered when connection is established and ready to send and accept data
func onConnect() {
	fmt.Println("Connected")
	fmt.Println("Subscribing to", instToken)
	err := ticker.Subscribe([]uint32{instToken})
	if err != nil {
		fmt.Println("err: ", err)
	}
	// Set subscription mode for given list of tokens
	// Default mode is Quote
	err = ticker.SetMode(realtime.ModeFull, []uint32{instToken})
	if err != nil {
		fmt.Println("err: ", err)
	}
}

// Triggered when tick is recevived
func onTick(tick models.Tick) {
	fmt.Println("Tick: ", tick)
}

// Triggered when reconnection is attempted which is enabled by default
func onReconnect(attempt int, delay time.Duration) {
	fmt.Printf("Reconnect attempt %d in %fs\n", attempt, delay.Seconds())
}

// Triggered when maximum number of reconnect attempt is made and the program is terminated
func onNoReconnect(attempt int) {
	fmt.Printf("Maximum no of reconnect attempt reached: %d", attempt)
}

// Triggered when order update is received
func onOrderUpdate(order kite.Order) {
	fmt.Printf("Order: %s", order.OrderID)
}

func main() {

	accessToken := ""
	// Create new Kite ticker instance
	ticker = realtime.New(accessToken)

	// Assign callbacks
	ticker.OnError(onError)
	ticker.OnClose(onClose)
	ticker.OnConnect(onConnect)
	ticker.OnReconnect(onReconnect)
	ticker.OnNoReconnect(onNoReconnect)
	ticker.OnTick(onTick)
	ticker.OnOrderUpdate(onOrderUpdate)

	// Start the connection
	ticker.Serve()
}

// getEnv returns the value of the environment variable provided.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// getEnvUint32 returns the value of the environment variable provided converted as Uint32.
func getEnvUint32(key string, fallback int) uint32 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return uint32(fallback)
		}
		return uint32(i)
	}
	return uint32(fallback)
}

```

## Examples

You can run the following after updating the API Keys in the examples:

```bash
go run examples/connect/basic/connect.go
```

## Development

#### Fetch mock responses for testcases

This needs to be run initially

```
git submodule update --init --recursive
```

#### Run unit tests

```
go test -v
```
