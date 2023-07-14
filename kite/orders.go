package kite

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// Order represents a individual order response.
type Order struct {
	AccountID string `json:"account_id"`
	PlacedBy  string `json:"placed_by"`

	OrderID                 string                 `json:"order_id"`
	ExchangeOrderID         string                 `json:"exchange_order_id"`
	ParentOrderID           string                 `json:"parent_order_id"`
	Status                  string                 `json:"status"`
	StatusMessage           string                 `json:"status_message"`
	StatusMessageRaw        string                 `json:"status_message_raw"`
	OrderTimestamp          Time                   `json:"order_timestamp"`
	ExchangeUpdateTimestamp Time                   `json:"exchange_update_timestamp"`
	ExchangeTimestamp       Time                   `json:"exchange_timestamp"`
	Variety                 string                 `json:"variety"`
	Meta                    map[string]interface{} `json:"meta"`

	Exchange        string `json:"exchange"`
	TradingSymbol   string `json:"tradingsymbol"`
	InstrumentToken uint32 `json:"instrument_token"`

	OrderType         string  `json:"order_type"`
	TransactionType   string  `json:"transaction_type"`
	Validity          string  `json:"validity"`
	ValidityTTL       int     `json:"validity_ttl"`
	Product           string  `json:"product"`
	Quantity          float64 `json:"quantity"`
	DisclosedQuantity float64 `json:"disclosed_quantity"`
	Price             float64 `json:"price"`
	TriggerPrice      float64 `json:"trigger_price"`

	AveragePrice      float64 `json:"average_price"`
	FilledQuantity    float64 `json:"filled_quantity"`
	PendingQuantity   float64 `json:"pending_quantity"`
	CancelledQuantity float64 `json:"cancelled_quantity"`

	Tag  string   `json:"tag"`
	Tags []string `json:"tags"`
}

// Orders is a list of orders.
type Orders []Order

// OrderParams represents parameters for placing an order.
type OrderParams struct {
	Exchange        string `url:"exchange,omitempty"`
	Tradingsymbol   string `url:"tradingsymbol,omitempty"`
	Validity        string `url:"validity,omitempty"`
	ValidityTTL     int    `url:"validity_ttl,omitempty"`
	Product         string `url:"product,omitempty"`
	OrderType       string `url:"order_type,omitempty"`
	TransactionType string `url:"transaction_type,omitempty"`

	Quantity          int     `url:"quantity,omitempty"`
	DisclosedQuantity int     `url:"disclosed_quantity,omitempty"`
	Price             float64 `url:"price,omitempty"`
	TriggerPrice      float64 `url:"trigger_price,omitempty"`

	Squareoff        float64 `url:"squareoff,omitempty"`
	Stoploss         float64 `url:"stoploss,omitempty"`
	TrailingStoploss float64 `url:"trailing_stoploss,omitempty"`

	IcebergLegs int `url:"iceberg_legs,omitempty"`
	IcebergQty  int `url:"iceberg_quantity,omitempty"`

	Tag string `json:"tag" url:"tag,omitempty"`
}

// OrderResponse represents the order place success response.
type OrderResponse struct {
	OrderID string `json:"order_id"`
}

// Trade represents an individual trade response.
type Trade struct {
	AveragePrice      float64 `json:"average_price"`
	Quantity          float64 `json:"quantity"`
	TradeID           string  `json:"trade_id"`
	Product           string  `json:"product"`
	FillTimestamp     Time    `json:"fill_timestamp"`
	ExchangeTimestamp Time    `json:"exchange_timestamp"`
	ExchangeOrderID   string  `json:"exchange_order_id"`
	OrderID           string  `json:"order_id"`
	TransactionType   string  `json:"transaction_type"`
	TradingSymbol     string  `json:"tradingsymbol"`
	Exchange          string  `json:"exchange"`
	InstrumentToken   uint32  `json:"instrument_token"`
}

// Trades is a list of trades.
type Trades []Trade

// ChargeOrderParams represents an individual charge order.
type ChargeOrderParams struct {
	AveragePrice    float64 `json:"average_price"`
	Exchange        string  `json:"exchange"`
	OrderID         string  `json:"order_id"`
	OrderType       string  `json:"order_type"`
	Product         string  `json:"product"`
	Quantity        int     `json:"quantity"`
	TradingSymbol   string  `json:"tradingsymbol"`
	TransactionType string  `json:"transaction_type"`
	Variety         string  `json:"variety"`
}

// ChargeOrders is a list of orders.
type ChargeOrders []ChargeOrderParams

type Charges struct {
	TransactionTax         float64 `json:"transaction_tax"`
	TransactionTaxType     string  `json:"transaction_tax_type"`
	ExchangeTurnoverCharge float64 `json:"exchange_turnover_charge"`
	SebiTurnoverCharge     float64 `json:"sebi_turnover_charge"`
	Brokerage              float64 `json:"brokerage"`
	StampDuty              float64 `json:"stamp_duty"`
	Total                  float64 `json:"total"`
	GST                    GST     `json:"gst"`
}

// ChargeOrderResponse -
type ChargeOrderResponse struct {
	TransactionType string  `json:"transaction_type"`
	Tradingsymbol   string  `json:"tradingsymbol"`
	Exchange        string  `json:"exchange"`
	Variety         string  `json:"variety"`
	Product         string  `json:"product"`
	OrderType       string  `json:"order_type"`
	Quantity        int     `json:"quantity"`
	Price           float64 `json:"price"`
	Charges         Charges `json:"charges"`
}

type GST struct {
	Igst  float64 `json:"igst"`
	Cgst  float64 `json:"cgst"`
	Sgst  float64 `json:"sgst"`
	Total float64 `json:"total"`
}

// GetOrders gets list of orders.
func (c *Client) GetOrders() (Orders, error) {
	var orders Orders
	err := c.doEnvelope(http.MethodGet, URIGetOrders, nil, nil, &orders)
	return orders, err
}

// GetOrdersOms - get oms orders
func (c *Client) GetOrdersOms() (Orders, error) {
	var orders Orders
	err := c.doEnvelope(http.MethodGet, URIGetOMSOrders, nil, nil, &orders)
	return orders, err
}

// GetTrades gets list of trades.
func (c *Client) GetTrades() (Trades, error) {
	var trades Trades
	err := c.doEnvelope(http.MethodGet, URIGetTrades, nil, nil, &trades)
	return trades, err
}

// GetOrderHistory gets history of an individual order.
func (c *Client) GetOrderHistory(OrderID string) ([]Order, error) {
	var orderHistory []Order
	err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIGetOrderHistory, OrderID), nil, nil, &orderHistory)
	return orderHistory, err
}

// GetOrderTrades gets list of trades executed for a particular order.
func (c *Client) GetOrderTrades(OrderID string) ([]Trade, error) {
	var orderTrades []Trade
	err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIGetOrderTrades, OrderID), nil, nil, &orderTrades)
	return orderTrades, err
}

// ChargeOrders - place a request for order charge
func (c *Client) ChargeOrders(chargeOrderParams []ChargeOrderParams) ([]ChargeOrderResponse, error) {
	var (
		chargeOrderResponse []ChargeOrderResponse
		// params              url.Values
		err error
	)

	// if params, err = query.Values(chargeOrderParams); err != nil {
	// 	return chargeOrderResponse, NewError(InputError, fmt.Sprintf("Error decoding order params: %v", err), nil)
	// }

	b, _ := json.Marshal(chargeOrderParams)

	// str := `[{"order_id":"230714200050319","exchange":"NFO","tradingsymbol":"NIFTY2372019500CE","transaction_type":"BUY","variety":"regular","product":"MIS","order_type":"MARKET","quantity":50,"average_price":91},{"order_id":"230714200050321","exchange":"NFO","tradingsymbol":"BANKNIFTY2372044900CE","transaction_type":"BUY","variety":"regular","product":"MIS","order_type":"MARKET","quantity":25,"average_price":226},{"order_id":"230714200087703","exchange":"NFO","tradingsymbol":"NIFTY2372019500CE","transaction_type":"SELL","variety":"regular","product":"MIS","order_type":"LIMIT","quantity":50,"average_price":103.6},{"order_id":"230714200440856","exchange":"NFO","tradingsymbol":"BANKNIFTY2372044900CE","transaction_type":"SELL","variety":"regular","product":"MIS","order_type":"MARKET","quantity":25,"average_price":242.8},{"order_id":"230714200514192","exchange":"BFO","tradingsymbol":"SENSEX2371465600PE","transaction_type":"BUY","variety":"regular","product":"NRML","order_type":"LIMIT","quantity":10,"average_price":105.7},{"order_id":"230714200518968","exchange":"BFO","tradingsymbol":"SENSEX2371465600PE","transaction_type":"BUY","variety":"regular","product":"NRML","order_type":"LIMIT","quantity":10,"average_price":100},{"order_id":"230714200572772","exchange":"BFO","tradingsymbol":"SENSEX2371465600PE","transaction_type":"SELL","variety":"regular","product":"NRML","order_type":"LIMIT","quantity":20,"average_price":103}]`
	// b = []byte(str)
	headers := http.Header{}
	headers.Add("content-type", "application/json")
	resp, err := c.doRaw(http.MethodPost, URIPlaceCharges, b, headers)
	err = readEnvelope(resp, &chargeOrderResponse)
	if err != nil {
		if _, ok := err.(Error); !ok {
			fmt.Printf("Error parsing JSON response: %v", err)
		}
	}
	return chargeOrderResponse, err
}

// PlaceOrder places an order.
func (c *Client) PlaceOrder(variety string, orderParams OrderParams) (OrderResponse, error) {
	var (
		orderResponse OrderResponse
		params        url.Values
		err           error
	)

	if params, err = query.Values(orderParams); err != nil {
		return orderResponse, NewError(InputError, fmt.Sprintf("Error decoding order params: %v", err), nil)
	}

	err = c.doEnvelope(http.MethodPost, fmt.Sprintf(URIPlaceOrder, variety), params, nil, &orderResponse)
	return orderResponse, err
}

// ModifyOrder modifies an order.
func (c *Client) ModifyOrder(variety string, orderID string, orderParams OrderParams) (OrderResponse, error) {
	var (
		orderResponse OrderResponse
		params        url.Values
		err           error
	)

	if params, err = query.Values(orderParams); err != nil {
		return orderResponse, NewError(InputError, fmt.Sprintf("Error decoding order params: %v", err), nil)
	}

	err = c.doEnvelope(http.MethodPut, fmt.Sprintf(URIModifyOrder, variety, orderID), params, nil, &orderResponse)
	return orderResponse, err
}

// CancelOrder cancels/exits an order.
func (c *Client) CancelOrder(variety string, orderID string, parentOrderID *string) (OrderResponse, error) {
	var (
		orderResponse OrderResponse
		params        url.Values
	)

	if parentOrderID != nil {
		// initialize the params map first
		params := url.Values{}
		params.Add("parent_order_id", *parentOrderID)
	}

	err := c.doEnvelope(http.MethodDelete, fmt.Sprintf(URICancelOrder, variety, orderID), params, nil, &orderResponse)
	return orderResponse, err
}

// ExitOrder is an alias for CancelOrder which is used to cancel/exit an order.
func (c *Client) ExitOrder(variety string, orderID string, parentOrderID *string) (OrderResponse, error) {
	return c.CancelOrder(variety, orderID, parentOrderID)
}
