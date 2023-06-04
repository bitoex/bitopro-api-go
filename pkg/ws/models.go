package ws

// OrderBookData struct
//
//	{
//	  "event": "ORDER_BOOK",
//	  "pair": "BTC_TWD",
//	  "bids": [
//	    {
//	      "price": "1",
//	      "amount": "1",
//	      "count": 1,
//	      "total": "1"
//	    }
//	    ...
//	  ],
//	  "asks": [
//	    {
//	      "price": "0.9",
//	      "amount": "1",
//	      "count": 2,
//	      "total": "2"
//	    }
//	    ...
//	  ],
//	  "timestamp": 1136185445000,
//	  "datetime": "2006-01-02T15:04:05.700Z"
//	}
type OrderBookData struct {
	Event string `json:"event"`
	Pair  string `json:"pair"`
	Bids  []struct {
		Price  string `json:"price"`
		Amount string `json:"amount"`
		Count  int    `json:"count"`
		Total  string `json:"total"`
	} `json:"bids"`
	Asks []struct {
		Price  string `json:"price"`
		Amount string `json:"amount"`
		Count  int    `json:"count"`
		Total  string `json:"total"`
	} `json:"asks"`
	Timestamp int64  `json:"timestamp"`
	DateTime  string `json:"datetime"`
	Err       error
}

// TickerData struct
//
//	{
//		"event": "TICKER",
//		"pair": "BTC_TWD",
//		"lastPrice": "1",
//		"isBuyer": true,
//		"priceChange24hr": "1",
//		"volume24hr": "1",
//		"high24hr": "1",
//		"low24hr": "1",
//		"timestamp": 1136185445000,
//		"datetime": "2006-01-02T15:04:05.700Z"
//	  }
type TickerData struct {
	Event           string `json:"event"`
	Pair            string `json:"pair"`
	LastPrice       string `json:"lastPrice"`
	IsBuyer         bool   `json:"isBuyer"`
	PriceChange24hr string `json:"priceChange24hr"`
	Volume24hr      string `json:"volume24hr"`
	High24hr        string `json:"high24hr"`
	Low24hr         string `json:"low24hr"`
	Timestamp       int64  `json:"timestamp"`
	DateTime        string `json:"datetime"`
	Err             error
}

// TradeData struct
//
//	{
//		"event": "TRADE",
//		"pair": "BTC_TWD",
//		"timestamp": 1136185445000,
//		"datetime": "2006-01-02T15:04:05.700Z",
//		"data": [
//		  {
//			"timestamp": 1136185445,
//			"price": "1",
//			"amount": "1",
//			"isBuyer": false
//		  },
//		  {
//			"timestamp": 1136185445,
//			"price": "1",
//			"amount": "1",
//			"isBuyer": true
//		  }
//		  ...
//		]
//	  }
type TradeData struct {
	Event     string `json:"event"`
	Pair      string `json:"pair"`
	Timestamp int64  `json:"timestamp"`
	DateTime  string `json:"datetime"`
	Data      []struct {
		Timestamp int64  `json:"timestamp"`
		Price     string `json:"price"`
		Amount    string `json:"amount"`
		IsBuyer   bool   `json:"isBuyer"`
	} `json:"data"`
	Err error
}

// AccountBalanceData
//
// {
// "event": "ACCOUNT_BALANCE",
// "timestamp": 1639553303365,
// "datetime": "2021-12-15T07:28:23.365Z",
//
//	"data": {
//	    "ADA": {
//	        "currency": "ADA",
//	        "amount": "999999999999.99999999",
//	        "available": "999999999999.99999999",
//	        "stake": "0",
//	        "tradable": true
//	    },
type AccountBalanceData struct {
	Event     string `json:"event"`
	Timestamp int64  `json:"timestamp"`
	DateTime  string `json:"datetime"`
	Data      map[string]struct {
		Currency  string `json:"currency"`
		Amount    string `json:"amount"`
		Available string `json:"available"`
		Stake     string `json:"stake"`
		Tradable  bool   `json:"tradable"`
	} `json:"data"`
	Err error
}

// OrdersData
//
//	{
//	    "event": "ACTIVE_ORDERS",
//	    "timestamp": 1639552073346,
//	    "datetime": "2021-12-15T07:07:53.346Z",
//	    "data": {
//	        "sol_usdt": [
//	            {
//	                "id": "8917255503",
//	                "pair": "sol_usdt",
//	                "price": "107",
//	                "avgExecutionPrice": "0",
//	                "action": "SELL",
//	                "type": "LIMIT",
//	                "timestamp": 1639386803663,
//	                "updatedTimestamp": 1639386803663,
//	                "createdTimestamp": 1639386803663,
//	                "status": 0,
//	                "originalAmount": "0.02",
//	                "remainingAmount": "0.02",
//	                "executedAmount": "0",
//	                "fee": "0",
//	                "feeSymbol": "usdt",
//	                "bitoFee": "0",
//	                "total": "0",
//	                "seq": "SOLUSDT3273528249",
//	                "timeInForce": "GTC"
//	            }
//	        ],
type OrdersData struct {
	Event     string `json:"event"`
	Timestamp int64  `json:"timestamp"`
	DateTime  string `json:"datetime"`
	Data      map[string][]struct {
		ID                string `json:"id"`
		Pair              string `json:"pair"`
		Price             string `json:"price"`
		AvgExecutionPrice string `json:"avgExecutionPrice"`
		Action            string `json:"action"`
		Type              string `json:"type"`
		Timestamp         int64  `json:"timestamp"`
		UpdatedTimestamp  int64  `json:"updatedTimestamp"`
		CreatedTimestamp  int64  `json:"createdTimestamp"`
		Status            int    `json:"status"`
		OriginalAmount    string `json:"originalAmount"`
		RemainingAmount   string `json:"remainingAmount"`
		ExecutedAmount    string `json:"executedAmount"`
		Fee               string `json:"fee"`
		FeeSymbol         string `json:"feeSymbol"`
		BitoFee           string `json:"bitoFee"`
		Total             string `json:"total"`
		Seq               string `json:"seq"`
		TimeInForce       string `json:"timeInForce"`
	} `json:"data"`
	Err error
}
