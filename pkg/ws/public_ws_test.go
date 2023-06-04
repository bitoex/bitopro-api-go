package ws

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// unit test for Ws.RunTickerWsConsumer
func TestTickerWsConsumer(t *testing.T) {
	ctx := context.Background()
	ws := NewPublicWs()
	tickers, close := ws.RunTickerWsConsumer(ctx, []string{"BTC_TWD"})
	if tickers == nil {
		t.Error("tickers is nil")
	}
	if close == nil {
		t.Error("close is nil")
	}
	ticker := <-tickers
	fmt.Printf("%+v", ticker)
	if ticker.Err != nil {
		t.Error("ticker.Err is not nil")
	}
	assert.Equal(t, ticker.Event, "TICKER")
	assert.Equal(t, ticker.Pair, "BTC_TWD")
	assert.NotEqual(t, ticker.LastPrice, "")
	assert.NotEqual(t, ticker.PriceChange24hr, "")
	assert.NotEqual(t, ticker.Volume24hr, "")
	assert.NotEqual(t, ticker.High24hr, "")
	assert.NotEqual(t, ticker.Low24hr, "")
	assert.NotEqual(t, ticker.Timestamp, 0)
	assert.NotEqual(t, ticker.DateTime, "")
}

// unit test for Ws.RunOrderbookWsConsumer
func TestOrderbookWsConsumer(t *testing.T) {
	ctx := context.Background()
	ws := NewPublicWs()
	orderbooks, close := ws.RunOrderbookWsConsumer(ctx, []string{"BTC_TWD:1"})
	if orderbooks == nil {
		t.Error("orderbooks is nil")
	}
	if close == nil {
		t.Error("close is nil")
	}
	orderbook := <-orderbooks
	fmt.Printf("%+v", orderbook)
	if orderbook.Err != nil {
		t.Error("orderbook.Err is not nil")
	}
	assert.Equal(t, orderbook.Event, "ORDER_BOOK")
	assert.Equal(t, orderbook.Pair, "BTC_TWD")
	assert.Len(t, orderbook.Asks, 1)
	assert.Len(t, orderbook.Bids, 1)
	assert.NotEqual(t, orderbook.Timestamp, 0)
	assert.NotEqual(t, orderbook.DateTime, "")
}

// unit test for Ws.RunTradesWsConsumer
func TestTradesWsConsumer(t *testing.T) {
	ctx := context.Background()
	ws := NewPublicWs()
	trades, close := ws.RunTradesWsConsumer(ctx, []string{"BTC_TWD"})
	if trades == nil {
		t.Error("trades is nil")
	}
	if close == nil {
		t.Error("close is nil")
	}
	trade := <-trades
	fmt.Printf("%+v", trade)
	if trade.Err != nil {
		t.Error("trade.Err is not nil")
	}
	assert.Equal(t, trade.Event, "TRADE")
	assert.Equal(t, trade.Pair, "BTC_TWD")
	assert.NotEqual(t, trade.Timestamp, 0)
	assert.NotEqual(t, trade.DateTime, "")
	assert.NotZero(t, trade.Data)
	assert.NotEmpty(t, trade.Data[0].Amount)
	assert.NotEmpty(t, trade.Data[0].Price)
	assert.NotEmpty(t, trade.Data[0].Timestamp)
}
