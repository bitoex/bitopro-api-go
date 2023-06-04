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
	accountBalances, close := ws.RunOrdersWsConsumer(ctx)
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
