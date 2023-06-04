# BitoPro API Client for Go

- [BitoPro API Client for Go](#BitoPro-API-Client-for-Go)
  - [Requirement](#Requirement)
  - [Installation](#Installation)
    - [Limitations](#Limitations)
      - [Rate Limit](#Rate-Limit)
      - [Precisions](#Precisions)
      - [Minimum order amount](#Minimum-order-amount)
  - [Getting Started](#Getting-Started)
      - [Public API](#Public-API)
      - [Authenticated API](#Authenticated-API)
        - [Basic](#Basic)
        - [From Viper](#From-Viper)
    - [Public REST Endpoints](#Public-REST-Endpoints)
      - [GetOrderBook](#GetOrderBook)
      - [GetTicker](#GetTicker)
      - [GetTickers](#GetTickers)
      - [GetTrades](#GetTrades)
    - [Authenticated REST Endpoints](#Authenticated-REST-Endpoints)
      - [GetAccountBalance](#GetAccountBalance)
      - [GetOrderHistory](#GetOrderHistory)
      - [GetOrderList](#GetOrderList)
      - [CreateOrderLimitBuy/CreateOrderLimitSell/CreateOrderMarketBuy/CreateOrderMarketSell](#CreateOrderLimitBuyCreateOrderLimitSellCreateOrderMarketBuyCreateOrderMarketSell)
      - [CancelOrder](#CancelOrder)
      - [GetOrder](#GetOrder)
    - [Websocket](#Websocket)
      - [Tickers Stream](#ticker-stream)
      - [Trades Stream](#trade-stream)
      - [OrderBook Stream](#orderbook-stream)
      - [AccountBalance Stream](#accountbalance-stream)
      - [UserOrder Stream](#userorders-stream)
  - [Contributing](#Contributing)
  - [License](#License)

## Requirement

Requires minimum of Go 1.11.

## Installation

```sh
$ go get github.com/bitoex/bitopro-api-go
```

### Limitations

#### Rate Limit

There is rate limits applied to each API, please check [API documentation](https://developer.bitopro.com/docs) for more detail.

#### Precisions

Both price and amount are subject to decimal restrictions, please check [official settings](https://www.bitopro.com/fees) for more detail.

#### Minimum order amount

Checkout the [official settings](https://www.bitopro.com/fees) of minimum amount.

## Getting Started

#### Public API

Methods for public APIs are packaged in `PubAPI` struct which can be created through `GetPubClient`.

```go
import "github.com/bitoex/bitopro-api-go/pkg/bitopro"

pubClient := bitopro.GetPubClient()
```

#### Authenticated API

Methods for authenticated APIs are packaged in `AuthAPI` struct which can be created in various ways through `GetAuthClient` depends on how you setup your authenticate information. To use the authenticated APIs, you need the following information, **API Key**, **API Secret**, **Identity (Account Email)**. You can create an API key
[here](https://www.bitopro.com/api). 

##### Basic

```go
import "github.com/bitoex/bitopro-api-go/pkg/bitopro"

authClient := bitopro.GetAuthClient("your identity (email)", "your key", "your secret")
```

##### From Viper

```yaml
# secret.yml
identity: <<your email>>
key: <<your key>>
secret: <<your secret>>
```

```go
import (
  "github.com/spf13/viper"
  
  "github.com/bitoex/bitopro-api-go/pkg/bitopro"
)

viper.AddConfigPath(".")
viper.SetConfigName("secret")
viper.ReadInConfig()

authAPI := GetAuthClient(viper.GetString("identity"), viper.GetString("key"), viper.GetString("secret"))
```

### Public REST Endpoints

#### GetOrderBook

[example](https://github.com/bitoex/bitopro-api-go/blob/master/pkg/bitopro/order_book_test.go)

```go
pubClient.GetOrderBook("btc_twd")
//or
authClient.GetOrderBook("btc_twd")
```

<details>
<summary>JSON Response</summary>

```js
{
  "bids": [
    {
      "price": "180500",
      "amount": "0.12817687",
      "count": 1,
      "total": "0.12817687"
    },
    {
      "price": "180010",
      "amount": "0.32292",
      "count": 2,
      "total": "0.45109687"
    },
    {
      "price": "180000",
      "amount": "0.24236",
      "count": 3,
      "total": "0.69345687"
    }
  ],
  "asks": [
    {
      "price": "180599",
      "amount": "0.00326056",
      "count": 1,
      "total": "0.00326056"
    },
    {
      "price": "180600",
      "amount": "0.04202575",
      "count": 1,
      "total": "0.04528631"
    }
  ]
}
```
</details>

#### GetTicker

[example](https://github.com/bitoex/bitopro-api-go/blob/master/pkg/bitopro/ticker_test.go)

```go
pubClient.GetTicker("btc_twd")
//or
authClient.GetTicker("btc_twd")
```

<details>
<summary>JSON Response</summary>

```js
{
  "data": {
    "pair": "btc_twd",
    "lastPrice": "0.00010800",
    "isBuyer": false,
    "priceChange24hr": "0",
    "volume24hr": "0.00000000",
    "high24hr": "0.00010800",
    "low24hr": "0.00010800"
  }
}
```
</details>

#### GetTickers

[example](https://github.com/bitoex/bitopro-api-go/blob/master/pkg/bitopro/tickers_test.go)

```go
pubClient.GetTickers()
//or
authClient.GetTickers()
```

<details>
<summary>JSON Response</summary>

```js
{
  "data": [
    {
      "pair": "xem_btc",
      "lastPrice": "0.00000098",
      "isBuyer": false,
      "priceChange24hr": "0",
      "volume24hr": "0.00000000",
      "high24hr": "0.00000098",
      "low24hr": "0.00000098"
    },
    {
      "pair": "bch_eth",
      "lastPrice": "0.60010000",
      "isBuyer": false,
      "priceChange24hr": "0",
      "volume24hr": "0.00000000",
      "high24hr": "0.60010000",
      "low24hr": "0.60010000"
    },
    {
      "pair": "eth_usdt",
      "lastPrice": "179.22000000",
      "isBuyer": true,
      "priceChange24hr": "10.85",
      "volume24hr": "925.14654180",
      "high24hr": "182.30000000",
      "low24hr": "159.94000000"
    }
  ]
}
```
</details>

#### GetTrades

[example](https://github.com/bitoex/bitopro-api-go/blob/master/pkg/bitopro/trades_test.go)

```go
pubClient.GetTrades("btc_twd")
//or
authClient.GetTrades("btc_twd")
```

<details>
<summary>JSON Response</summary>

```js
{
  "data": [
    {
      "timestamp": 1557203407,
      "price": "180500.00000000",
      "amount": "0.07717687",
      "isBuyer": false
    },
    {
      "timestamp": 1557203187,
      "price": "180500.00000000",
      "amount": "0.05100000",
      "isBuyer": false
    },
    {
      "timestamp": 1557203053,
      "price": "180500.00000000",
      "amount": "0.01860000",
      "isBuyer": false
    },
    {
      "timestamp": 1557202804,
      "price": "180500.00000000",
      "amount": "0.04781533",
      "isBuyer": false
    },
    {
      "timestamp": 1557202804,
      "price": "180500.00000000",
      "amount": "0.06000000",
      "isBuyer": false
    }
  ]
}
```

</details>

### Authenticated REST Endpoints

#### GetAccountBalance

[example](https://github.com/bitoex/bitopro-api-go/blob/master/pkg/bitopro/account_balance_test.go)

```go
authClient.GetAccountBalance()
```

<details>
<summary>JSON Response</summary>

```js
{
  "data": [
    {
      "amount": "10001",
      "available": "1.0",
      "currency": "bito",
      "stake": "10000"
    },
    {
      "amount": "0.0",
      "available": "1.0",
      "currency": "btc",
      "stake": "0"
    },
    {
      "amount": "3.0",
      "available": "0.01",
      "currency": "eth",
      "stake": "0"
    },
    {
      "amount": "30000",
      "available": "2500",
      "currency": "twd",
      "stake": "0"
    }
  ]
}
```

</details>

#### GetOrderHistory

[example](https://github.com/bitoex/bitopro-api-go/blob/master/pkg/bitopro/order_history_test.go)

```go
authClient.GetOrderHistory()
```

<details>
<summary>JSON Response</summary>

```js
{
  "data": [
    {
      "action": "buy",
      "avgExecutionPrice": "100000.00000000",
      "bitoFee": "0.00000000",
      "executedAmount": "1.00000000",
      "fee": "0.00100000",
      "feeSymbol": "BTC",
      "id": "123",
      "originalAmount": "1.00000000",
      "pair": "btc_twd",
      "price": "100000.00000000",
      "remainingAmount": "0.00000000",
      "status": 2,
      "timestamp": 1508753757000,
      "type": "limit"
    },
    {
      "action": "buy",
      "avgExecutionPrice": "100000.00000000",
      "bitoFee": "0.00000000",
      "executedAmount": "1.00000000",
      "fee": "0.00200000",
      "feeSymbol": "BTC",
      "id": "456",
      "originalAmount": "1.00000000",
      "pair": "btc_twd",
      "price": "100000.00000000",
      "remainingAmount": "0.00000000",
      "status": 2,
      "timestamp": 1508753787000,
      "type": "limit"
    }
  ]
}
```

</details>

#### GetOrderList

[example](https://github.com/bitoex/bitopro-api-go/blob/master/pkg/bitopro/order_list_test.go)

```go
authClient.GetOrderList("btc_twd", false, 1)
```

<details>
<summary>JSON Response</summary>

```js
{
  "data": [
    {
      "action": "buy",
      "avgExecutionPrice": "100000.00000000",
      "bitoFee": "0.00000000",
      "executedAmount": "1.00000000",
      "fee": "0.00100000",
      "feeSymbol": "BTC",
      "id": "123",
      "originalAmount": "1.00000000",
      "pair": "btc_twd",
      "price": "100000.00000000",
      "remainingAmount": "0.00000000",
      "status": 2,
      "timestamp": 1508753757000,
      "type": "limit"
    },
    {
      "action": "buy",
      "avgExecutionPrice": "100000.00000000",
      "bitoFee": "0.00000000",
      "executedAmount": "1.00000000",
      "fee": "0.00200000",
      "feeSymbol": "BTC",
      "id": "456",
      "originalAmount": "1.00000000",
      "pair": "btc_twd",
      "price": "100000.00000000",
      "remainingAmount": "0.00000000",
      "status": 2,
      "timestamp": 1508753787000,
      "type": "limit"
    }
  ],
  "page": 1,
  "totalPages": 10
}
```

</details>

#### CreateOrderLimitBuy/CreateOrderLimitSell/CreateOrderMarketBuy/CreateOrderMarketSell

[example](https://github.com/bitoex/bitopro-api-go/blob/master/pkg/bitopro/order_create_test.go)

```go
// create limit buy order
authClient.CreateOrderLimitBuy("eth_twd", "0.1", "1")
// create limit sell order
authClient.CreateOrderLimitSell("eth_twd", "0.1", "1")
// create market buy order
authClient.CreateOrderMarketBuy("eth_twd", "1")
// create market sell order
authClient.CreateOrderMarketSell("eth_twd", "1")
```

<details>
<summary>JSON Response</summary>

```js
{
  "action": "buy",
  "amount": "0.235",
  "orderId": "11233456",
  "price": "1.0",
  "timestamp": 1504262258000
}
```

</details>

#### CancelOrder

[example](https://github.com/bitoex/bitopro-api-go/blob/master/pkg/bitopro/order_cancel_test.go)

```go
authClient.CancelOrder("eth_twd", 7517762903)
```

<details>
<summary>JSON Response</summary>

```js
{
  "action": "buy",
  "amount": 2.3,
  "orderId": "7517762903",
  "price": 1.2,
  "timestamp": 1504262258000
}
```

</details>

#### GetOrder

[example](https://github.com/bitoex/bitopro-api-go/blob/master/pkg/bitopro/order_get_test.go)

```go
authClient.GetOrder("btc_twd", 2640904509)
```

<details>
<summary>JSON Response</summary>

```js
{
  "action": "sell",
  "avgExecutionPrice": "112000.00000000",
  "bitoFee": "103.70370360",
  "executedAmount": "1.00000000",
  "fee": "0.00000000",
  "feeSymbol": "TWD",
  "id": "2640904509",
  "originalAmount": "1.00000000",
  "pair": "btc_twd",
  "price": "112000.00000000",
  "remainingAmount": "0.00000000",
  "status": 2,
  "timestamp": 1508753757000,
  "type": "limit"
}
```

</details>

### Websocket

```go
publicWs := ws.NewPublicWs()
privateWs := ws.NewPrivateWs("email", "api_key", "api_secret")
```


#### Ticker Stream
[example](https://github.com/bitoex/bitopro-api-go/blob/master/pkg/ws/public_ws_test.go#L12)

```go
tickers, closeCh := publicWs.RunTickerWsConsumer(ctx, []string{"BTC_TWD"})
defer close(closeCh)

for {
  ticker <- tickers
  if ticker.Err != nil {
    fmt.Printf("%+v\n", err)
    return
  }
  fmt.Printf("%+v\n", ticker)
}

```

#### OrderBook Stream
[example](https://github.com/bitoex/bitopro-api-go/blob/master/pkg/ws/public_ws_test.go#L39)

```go
orderBooks, closeCh := publicWs.RunOrderbookWsConsumer(ctx, []string{"BTC_TWD:30"})
defer close(closeCh)

for {
  orderBook <- orderBooks
  if orderBook.Err != nil {
    fmt.Printf("%+v\n", err)
    return
  }
  fmt.Printf("%+v\n", orderBook)
  fmt.Printf("%+v\n", len(orderBook.Bids)) // => 30
  fmt.Printf("%+v\n", len(orderBook.Asks))  // => 30
}

```

#### Trade Stream
[example](https://github.com/bitoex/bitopro-api-go/blob/master/pkg/ws/public_ws_test.go#L62)

```go
trades, closeCh := publicWs.RunTradesWsConsumer(ctx, []string{"BTC_TWD"})
defer close(closeCh)

for {
  trade <- trades
  if trade.Err != nil {
    fmt.Printf("%+v\n", err)
    return
  }
  fmt.Printf("%+v\n", trade)
}
```

#### AccountBalance Stream
[example](https://github.com/bitoex/bitopro-api-go/blob/master/pkg/ws/private_ws_test.go#L9)

```go
accBalances, closeCh := privateWs.RunAccountBalancesWsConsumer(ctx)
defer close(closeCh)

for {
  accBalance <- accBalances
  if accBalance.Err != nil {
    fmt.Printf("%+v\n", err)
    return
  }
  fmt.Printf("%+v\n", accBalance)
}
```

#### UserOrders Stream
[example](https://github.com/bitoex/bitopro-api-go/blob/master/pkg/ws/private_ws_test.go#L26)

```go
ordersList, closeCh := privateWs.RunOrdersWsConsumer(ctx)
defer close(closeCh)

for {
  orders <- ordersList
  if orders.Err != nil {
    fmt.Printf("%+v\n", err)
    return
  }
  fmt.Printf("%+v\n", orders)
}
```

## Contributing

Bug reports and pull requests are welcome on GitHub at [bitopro-api-go](https://github.com/bitoex/bitopro-api-go) and this project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

1. Fork it.
2. Create your feature branch (```git checkout -b my-new-feature```).
3. Commit your changes (```git commit -am 'Added some feature'```).
4. Push to the branch (```git push origin my-new-feature```).
5. Create new Pull Request.

## License

The SDK is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).