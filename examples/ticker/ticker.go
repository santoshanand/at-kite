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
