package ws

import (
	"context"
	"fmt"
	"testing"
)

func TestAccountBalanceWsConsumer(t *testing.T) {
	ctx := context.Background()
	ws := NewWs("", "", "", "")
	accountBalances, close := ws.RunAccountBalancesWsConsumer(ctx)
	if accountBalances == nil {
		t.Error("accountBalances is nil")
	}
	if close == nil {
		t.Error("close is nil")
	}
	accountBalance := <-accountBalances
	fmt.Printf("%+v", accountBalance)
	if accountBalance.Err != nil {
		t.Error("trade.Err is not nil")
	}
}

func TestOrderWsConsumer(t *testing.T) {
	ctx := context.Background()
	ws := NewWs("", "", "", "")
	orders, close := ws.RunOrdersWsConsumer(ctx)
	if orders == nil {
		t.Error("Orders is nil")
	}
	if close == nil {
		t.Error("close is nil")
	}
	order := <-orders
	fmt.Printf("%+v", order)
	if order.Err != nil {
		t.Error("trade.Err is not nil")
	}
}

func TestHistoryOrdersWsConsumer(t *testing.T) {
	ctx := context.Background()
	ws := NewWs("", "", "", "")
	historyOrders, close := ws.RunOrdersWsConsumer(ctx)
	if historyOrders == nil {
		t.Error("HistoryOrders is nil")
	}
	if close == nil {
		t.Error("close is nil")
	}
	historyOrder := <-historyOrders
	fmt.Printf("%+v", historyOrder)
	if historyOrder.Err != nil {
		t.Error("trade.Err is not nil")
	}
}

func TestMatchesOrdersWsConsumer(t *testing.T) {
	ctx := context.Background()
	ws := NewWs("", "", "", "")
	matchesOrders, close := ws.RunOrdersWsConsumer(ctx)
	if matchesOrders == nil {
		t.Error("MatchesOrders is nil")
	}
	if close == nil {
		t.Error("close is nil")
	}
	matchesOrder := <-matchesOrders
	fmt.Printf("%+v", matchesOrder)
	if matchesOrder.Err != nil {
		t.Error("trade.Err is not nil")
	}
}
