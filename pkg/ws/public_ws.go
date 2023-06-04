package ws

import (
	"context"
	"encoding/json"
	"strings"
)

func (ws Ws) RunTickerWsConsumer(ctx context.Context, pairs []string) (tickers chan TickerData, close chan struct{}) {
	path := "ws/v1/pub/tickers?pairs=" + strings.Join(pairs, ",")
	tickers = make(chan TickerData, 1000)
	close = ws.runWsConnConsumer(ctx, path, false, func(msg []byte, err error) {
		if err != nil {
			tickers <- TickerData{Err: err}
			return
		}

		var ticker TickerData
		err = json.Unmarshal(msg, &ticker)
		if err != nil {
			tickers <- TickerData{Err: err}
			return
		}
		tickers <- ticker
	})
	return tickers, close
}

func (ws Ws) RunOrderbookWsConsumer(ctx context.Context, pairLimits []string) (orderbooks chan OrderBookData, close chan struct{}) {
	path := "ws/v1/pub/order-books?pairs=" + strings.Join(pairLimits, ",")
	orderbooks = make(chan OrderBookData, 1000)
	close = ws.runWsConnConsumer(ctx, path, false, func(msg []byte, err error) {
		if err != nil {
			orderbooks <- OrderBookData{Err: err}
			return
		}

		var orderbook OrderBookData
		err = json.Unmarshal(msg, &orderbook)
		if err != nil {
			orderbooks <- OrderBookData{Err: err}
			return
		}
		orderbooks <- orderbook
	})
	return orderbooks, close
}

func (ws Ws) RunTradesWsConsumer(ctx context.Context, pairs []string) (trades chan TradeData, close chan struct{}) {
	path := "ws/v1/pub/trades?pairs=" + strings.Join(pairs, ",")
	trades = make(chan TradeData, 1000)
	close = ws.runWsConnConsumer(ctx, path, false, func(msg []byte, err error) {
		if err != nil {
			trades <- TradeData{Err: err}
			return
		}

		var trade TradeData
		err = json.Unmarshal(msg, &trade)
		if err != nil {
			trades <- TradeData{Err: err}
			return
		}
		trades <- trade
	})
	return trades, close
}
