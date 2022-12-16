package main

import (
	"log"

	"github.com/santoshanand/at-kite/kite"
)

func main() {
	accessToken := ""

	// Create a new Kite connect instance
	kc := kite.New(accessToken)

	log.Println("Fetching GTTs...")
	orders, err := kc.GetGTTs()
	if err != nil {
		log.Fatalf("Error getting GTTs: %v", err)
	}
	log.Printf("gtt: %v", orders)

	log.Println("Placing GTT...")
	// Place GTT
	gttResp, err := kc.PlaceGTT(kite.GTTParams{
		Tradingsymbol:   "INFY",
		Exchange:        "NSE",
		LastPrice:       800,
		TransactionType: kite.TransactionTypeBuy,
		Trigger: &kite.GTTSingleLegTrigger{
			TriggerParams: kite.TriggerParams{
				TriggerValue: 1,
				Quantity:     1,
				LimitPrice:   1,
			},
		},
	})
	if err != nil {
		log.Fatalf("error placing gtt: %v", err)
	}

	log.Println("placed GTT trigger_id = ", gttResp.TriggerID)

	log.Println("Fetching details of placed GTT...")

	order, err := kc.GetGTT(gttResp.TriggerID)
	if err != nil {
		log.Fatalf("Error getting GTTs: %v", err)
	}
	log.Printf("gtt: %v", order)

	log.Println("Modify existing GTT...")

	gttModifyResp, err := kc.ModifyGTT(gttResp.TriggerID, kite.GTTParams{
		Tradingsymbol:   "INFY",
		Exchange:        "NSE",
		LastPrice:       800,
		TransactionType: kite.TransactionTypeBuy,
		Trigger: &kite.GTTSingleLegTrigger{
			TriggerParams: kite.TriggerParams{
				TriggerValue: 2,
				Quantity:     2,
				LimitPrice:   2,
			},
		},
	})
	if err != nil {
		log.Fatalf("error placing gtt: %v", err)
	}

	log.Println("modified GTT trigger_id = ", gttModifyResp.TriggerID)

	gttDeleteResp, err := kc.DeleteGTT(gttResp.TriggerID)
	if err != nil {
		log.Fatalf("Error getting GTTs: %v", err)
	}
	log.Printf("gtt deleted: %v", gttDeleteResp)
}
