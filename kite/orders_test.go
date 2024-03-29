package kite

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetOrders(t *testing.T) {
	t.Parallel()
	orders, err := getKite().GetOrders()
	if err != nil {
		t.Errorf("Error while fetching orders. %v", err)
	}
	t.Run("test empty/unparsed orders", func(t *testing.T) {
		for _, order := range orders {
			require.NotEqual(t, "", order.OrderID)
		}
	})
	t.Run("test tag parsing", func(t *testing.T) {
		require.Equal(t, "", orders[0].Tag)
		require.Equal(t, "connect test order1", orders[1].Tag)
		require.Equal(t, []string{"connect test order2", "XXXXX"}, orders[2].Tags)
	})
	t.Run("test ice-berg and TTL orders", func(t *testing.T) {
		require.Equal(t, "iceberg", orders[3].Variety)
		require.Equal(t, "TTL", orders[3].Validity)
		require.Equal(t, 200.0, orders[3].Meta["iceberg"].(map[string]interface{})["leg_quantity"])
		require.Equal(t, 1000.0, orders[3].Meta["iceberg"].(map[string]interface{})["total_quantity"])
	})
}

func TestGetOrdersOMS(t *testing.T) {
	t.Parallel()
	orders, err := getKite().GetOrdersOms()
	if err != nil {
		t.Errorf("Error while fetching orders. %v", err)
	}
	t.Run("test empty/unparsed orders", func(t *testing.T) {
		for _, order := range orders {
			require.NotEqual(t, "", order.OrderID)
		}
	})
	t.Run("test tag parsing", func(t *testing.T) {
		require.Equal(t, "", orders[0].Tag)
		require.Equal(t, "connect test order1", orders[1].Tag)
		require.Equal(t, []string{"connect test order2", "XXXXX"}, orders[2].Tags)
	})
	t.Run("test ice-berg and TTL orders", func(t *testing.T) {
		require.Equal(t, "iceberg", orders[3].Variety)
		require.Equal(t, "TTL", orders[3].Validity)
		require.Equal(t, 200.0, orders[3].Meta["iceberg"].(map[string]interface{})["leg_quantity"])
		require.Equal(t, 1000.0, orders[3].Meta["iceberg"].(map[string]interface{})["total_quantity"])
	})
}

func TestGetTrades(t *testing.T) {
	t.Parallel()
	trades, err := getKite().GetTrades()
	if err != nil {
		t.Errorf("Error while fetching trades. %v", err)
	}
	for _, trade := range trades {
		if trade.TradeID == "" {
			t.Errorf("Error while fetching trade id in trades. %v", err)
		}
	}
}

func TestGetOrderHistory(t *testing.T) {
	t.Parallel()
	orderHistory, err := getKite().GetOrderHistory("test")
	if err != nil {
		t.Errorf("Error while fetching trades. %v", err)
	}
	for _, order := range orderHistory {
		if order.OrderID == "" {
			t.Errorf("Error while fetching order id in order history. %v", err)
		}
	}
}

func TestGetOrderTrades(t *testing.T) {
	t.Parallel()
	tradeHistory, err := getKite().GetOrderTrades("test")
	if err != nil {
		t.Errorf("Error while fetching trades. %v", err)
	}
	for _, trade := range tradeHistory {
		if trade.TradeID == "" {
			t.Errorf("Error while fetching trade id in trade history. %v", err)
		}
	}
}

func TestPlaceOrder(t *testing.T) {
	t.Parallel()
	params := OrderParams{
		Exchange:          "test",
		Tradingsymbol:     "test",
		Validity:          "test",
		Product:           "test",
		OrderType:         "test",
		TransactionType:   "test",
		Quantity:          100,
		DisclosedQuantity: 100,
		Price:             100,
		TriggerPrice:      100,
		Squareoff:         100,
		Stoploss:          100,
		TrailingStoploss:  100,
		Tag:               "test",
	}
	orderResponse, err := getKite().PlaceOrder("test", params)
	if err != nil {
		t.Errorf("Error while placing order. %v", err)
	}
	if orderResponse.OrderID == "" {
		t.Errorf("No order id returned. Error %v", err)
	}
}

func TestModifyOrder(t *testing.T) {
	t.Parallel()
	params := OrderParams{
		Exchange:          "test",
		Tradingsymbol:     "test",
		Validity:          "test",
		Product:           "test",
		OrderType:         "test",
		TransactionType:   "test",
		Quantity:          100,
		DisclosedQuantity: 100,
		Price:             100,
		TriggerPrice:      100,
		Squareoff:         100,
		Stoploss:          100,
		TrailingStoploss:  100,
		Tag:               "test",
	}
	orderResponse, err := getKite().ModifyOrder("test", "test", params)
	if err != nil {
		t.Errorf("Error while placing order. %v", err)
	}
	if orderResponse.OrderID == "" {
		t.Errorf("No order id returned. Error %v", err)
	}
}

func TestCancelOrder(t *testing.T) {
	t.Parallel()
	parentOrderID := "test"

	orderResponse, err := getKite().CancelOrder("test", "test", &parentOrderID)
	if err != nil || orderResponse.OrderID == "" {
		t.Errorf("Error while placing cancel order. %v", err)
	}
}

func TestExitOrder(t *testing.T) {
	t.Parallel()
	parentOrderID := "test"

	orderResponse, err := getKite().ExitOrder("test", "test", &parentOrderID)
	if err != nil {
		t.Errorf("Error while placing order. %v", err)
	}
	if orderResponse.OrderID == "" {
		t.Errorf("No order id returned. Error %v", err)
	}
}

func TestIssue64(t *testing.T) {
	t.Parallel()
	orders, err := getKite().GetOrders()
	if err != nil {
		t.Errorf("Error while fetching orders. %v", err)
	}

	// Check if marshal followed by unmarshall correctly parses timestamps
	ord := orders[0]
	js, err := json.Marshal(ord)
	if err != nil {
		t.Errorf("Error while marshalling order. %v", err)
	}

	var outOrd Order
	err = json.Unmarshal(js, &outOrd)
	if err != nil {
		t.Errorf("Error while unmarshalling order. %v", err)
	}

	if !ord.ExchangeTimestamp.Equal(outOrd.ExchangeTimestamp.Time) {
		t.Errorf("Incorrect timestamp parsing.\nwant:\t%v\ngot:\t%v", ord.ExchangeTimestamp, outOrd.ExchangeTimestamp)
	}
}

func TestChargeOrder(t *testing.T) {
	t.Parallel()
	params := []ChargeOrderParams{{
		AveragePrice:    91,
		Exchange:        "NFO",
		OrderID:         "230714200050319",
		OrderType:       "MARKET",
		Product:         "MIS",
		Quantity:        50,
		TradingSymbol:   "NIFTY2372019500CE",
		TransactionType: "BUY",
		Variety:         "regular",
	}}
	charges, err := getKite().ChargeOrders(params)
	assert.Nil(t, err)
	assert.NotNil(t, charges)
}
