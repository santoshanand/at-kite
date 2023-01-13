//go:build exclude

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
