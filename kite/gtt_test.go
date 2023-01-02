package kite

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGTTs(t *testing.T) {
	t.Parallel()
	gttOrders, err := getKite().GetGTTs()
	if err != nil {
		t.Errorf("Error while fetching GTT orders. %v", err)
	}
	for _, gttOrder := range gttOrders {
		if gttOrder.ID == 0 {
			t.Errorf("Error while parsing order id in GTT orders. %v", err)
		}
	}
}

func TestGetGTT(t *testing.T) {
	t.Parallel()
	gttOrder, err := getKite().GetGTT(123)
	if err != nil {
		t.Errorf("Error while fetching GTT orders. %v", err)
	}
	if gttOrder.ID != 123 {
		t.Errorf("Error while parsing order id in GTT order. %v", err)
	}
}

func TestModifyGTT(t *testing.T) {
	t.Parallel()
	gttOrder, err := getKite().ModifyGTT(123, GTTParams{
		Tradingsymbol:   "INFY",
		Exchange:        "NSE",
		LastPrice:       800,
		TransactionType: TransactionTypeBuy,
		Trigger: &GTTSingleLegTrigger{
			TriggerParams: TriggerParams{
				TriggerValue: 2,
				Quantity:     2,
				LimitPrice:   2,
			},
		},
	})
	if err != nil {
		t.Errorf("Error while fetching GTT orders. %v", err)
	}
	if gttOrder.TriggerID != 123 {
		t.Errorf("Error while parsing order id in GTT order. %v", err)
	}
}

func TestPlaceGTT(t *testing.T) {
	t.Parallel()
	gttOrder, err := getKite().PlaceGTT(GTTParams{
		Tradingsymbol:   "INFY",
		Exchange:        "NSE",
		LastPrice:       800,
		TransactionType: TransactionTypeBuy,
		Trigger: &GTTSingleLegTrigger{
			TriggerParams: TriggerParams{
				TriggerValue: 1,
				Quantity:     1,
				LimitPrice:   1,
			},
		},
	})

	assert.Nil(t, err)
	assert.NotNil(t, gttOrder)
}

func TestDeleteGTT(t *testing.T) {
	t.Parallel()
	gttOrder, err := getKite().DeleteGTT(132359442)
	assert.Nil(t, err)
	assert.NotNil(t, gttOrder)
}
