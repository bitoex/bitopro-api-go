package ws

import (
	"context"
	"encoding/json"
)

func (ws Ws) RunAccountBalancesWsConsumer(ctx context.Context) (account chan AccountBalanceData, close chan struct{}) {
	path := "ws/v1/pub/auth/account-balance"
	account = make(chan AccountBalanceData, 1000)
	close = ws.runWsConnConsumer(ctx, path, true, func(msg []byte, err error) {
		if err != nil {
			account <- AccountBalanceData{Err: err}
			return
		}

		var accData AccountBalanceData
		err = json.Unmarshal(msg, &accData)
		if err != nil {
			account <- AccountBalanceData{Err: err}
			return
		}
		account <- accData
	})
	return account, close
}

func (ws Ws) RunOrdersWsConsumer(ctx context.Context) (orders chan OrdersData, close chan struct{}) {
	path := "ws/v1/pub/auth/orders"
	orders = make(chan OrdersData, 1000)
	close = ws.runWsConnConsumer(ctx, path, true, func(msg []byte, err error) {
		if err != nil {
			orders <- OrdersData{Err: err}
			return
		}

		var order OrdersData
		err = json.Unmarshal(msg, &order)
		if err != nil {
			orders <- OrdersData{Err: err}
			return
		}
		orders <- order
	})
	return orders, close
}
