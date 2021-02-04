package cexio

// CexURL CEX.IO url
const (
	CexURL          string = "https://cex.io/api/"
	postContentType string = "application/json"
)

// Public Resources
// GET Methods
const (
	MethodCurrencyLimits  string = "currency_limits"
	MethodTicker          string = "ticker"
	MethodTickers         string = "tickers"
	MethodLastPrice       string = "last_price"
	MethodLastPrices      string = "last_prices"
	MethodHistoricalChart string = "ohlcv/hd"
	MethodOrderBook       string = "order_book"
	MethodTradeHistory    string = "trade_history"
)

// Public Resources
// POST Methods
const (
	MethodConvert string = "convert"
	MethodChart   string = "price_stats"
)

// Private Resources
const (
	MethodBalance              string = "balance"
	MethodOpenOrders           string = "open_orders"
	MethodActiveOrdersStatus   string = "active_orders_status"
	MethodArchivedOrders       string = "archived_orders"
	MethodCancelOrder          string = "cancel_order"
	MethodCancelOrders         string = "cancel_orders"
	MethodPlaceOrder           string = "place_order"
	MethodGetOrder             string = "get_order"
	MethodGetOrderTransactions string = "get_order_tx"
	MethodGetAddress           string = "get_address"
	MethodGetMyFee             string = "get_myfee"
	MethodCancelReplaceOrder   string = "cancel_replace_order"
	MethodOpenPosition         string = "open_position"
	MethodOpenPositions        string = "open_positions"
	MethodGetPosition          string = "get_position"
	MethodClosePosition        string = "close_position"
	MethodArchivedPositions    string = "archived_positions"
	MethodGetMarginalFee       string = "get_marginal_fee"
)

const (
	errorHTTPComunication string = "Error with HTTP comunication to CEX.IO: %s\n"
	errorReadingResponse  string = "Error reading CEX.IO response: %s\n"
	errorDecodingResponse string = "Error decoding CEX.IO response: %s\n"
	errorInvalidAPIKeys   string = "Error the CEX.IO Keys are not defined\n"
)
