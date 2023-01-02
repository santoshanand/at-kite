package kite

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetQuote(t *testing.T) {
	t.Parallel()
	marketQuote, err := getKite().GetQuote()
	if err != nil {
		t.Errorf("Error while fetching MF orders. %v", err)
	}

	if q, ok := marketQuote["NSE:INFY"]; ok {
		if q.InstrumentToken != 408065 {
			t.Errorf("Incorrect values set. %v", err)
		}
	} else {
		t.Errorf("Key wanted but not found. %v", err)
	}
}

func TestGetLTP(t *testing.T) {
	t.Parallel()
	marketLTP, err := getKite().GetLTP()
	if err != nil {
		t.Errorf("Error while fetching MF orders. %v", err)
	}

	if ltp, ok := marketLTP["NSE:INFY"]; ok {
		if ltp.InstrumentToken != 408065 {
			t.Errorf("Incorrect values set. %v", err)
		}
	} else {
		t.Errorf("Key wanted but not found. %v", err)
	}
}

func TestGetHistoricalData(t *testing.T) {
	t.Parallel()
	marketHistorical, err := getKite().GetHistoricalData(123, "myinterval", time.Unix(0, 0), time.Unix(1, 0), true, false)
	if err != nil {
		t.Errorf("Error while fetching MF orders. %v", err)
	}

	for i := 0; i < len(marketHistorical)-1; i++ {
		if marketHistorical[i].Date.Unix() > marketHistorical[i+1].Date.Unix() {
			t.Errorf("Unsorted candles returned. %v", err)
			return
		}
	}
}

func TestGetHistoricalDataWithOI(t *testing.T) {
	t.Parallel()
	marketHistorical, err := getKite().GetHistoricalData(456, "myinterval", time.Unix(0, 0), time.Unix(1, 0), true, true)
	require.Nil(t, err)
	require.Equal(t, 6, len(marketHistorical))

	for i := 0; i < len(marketHistorical)-1; i++ {
		require.Greater(t, marketHistorical[i+1].Date.Unix(), marketHistorical[i].Date.Unix())
		require.NotEqual(t, marketHistorical[i].OI, 0)
	}
}

func TestGetOHLC(t *testing.T) {
	t.Parallel()
	marketOHLC, err := getKite().GetOHLC()
	if err != nil {
		t.Errorf("Error while fetching MF orders. %v", err)
	}

	if ohlc, ok := marketOHLC["NSE:INFY"]; ok {
		if ohlc.InstrumentToken != 408065 {
			t.Errorf("Incorrect values set. %v", err)
		}
	} else {
		t.Errorf("Key wanted but not found. %v", err)
	}
}

func TestGetInstruments(t *testing.T) {
	t.Parallel()
	marketInstruments, err := getKite().GetInstruments()
	if err != nil {
		t.Errorf("Error while fetching MF orders. %v", err)
	}

	for _, mInstr := range marketInstruments {
		if mInstr.InstrumentToken == 0 {
			t.Errorf("Incorrect data loaded. %v", err)
		}

		if mInstr.InstrumentToken == 12074242 {
			if mInstr.Expiry.Year() != 2018 {
				t.Errorf("Incorrectly parsed timestamp for instruments")
			}
		}
	}
}

func TestGetInstrumentsByExchange(t *testing.T) {
	t.Parallel()
	marketInstruments, err := getKite().GetInstrumentsByExchange("nse")
	if err != nil {
		t.Errorf("Error while fetching MF orders. %v", err)
	}

	for _, mInstr := range marketInstruments {
		if mInstr.Exchange != "NSE" {
			t.Errorf("Incorrect data loaded. %v", err)
		}
	}
}

func TestGetMFInstruments(t *testing.T) {
	t.Parallel()
	marketInstruments, err := getKite().GetMFInstruments()
	if err != nil {
		t.Errorf("Error while fetching MF orders. %v", err)
	}

	for _, mInstr := range marketInstruments {
		if mInstr.Tradingsymbol == "" {
			t.Errorf("Incorrect data loaded. %v", err)
		}
	}
}
