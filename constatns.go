package main

type PairStatusType string

const (
	InStageOfEstimation               PairStatusType = "IN_STAGE_OF_COMPUTATION"
	InStageOfTrading                  PairStatusType = "IN_STAGE_OF_TRADING"
	InStageOfOrderPlacedOrderTracking PairStatusType = "IN_STAGE_OF_PLACED_ORDER_TRACKING"
)

type SideType string

const (
	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"
)

type ExchangeName string

const (
	Binance        ExchangeName = "BINANCE"
	BinanceFutures ExchangeName = "BINANCE_FUTURES"
	Uniswap        ExchangeName = "UNISWAP"
	ByBit          ExchangeName = "BUYBIT"
)

type Configs struct {
}

type InProcess struct {
}

type OrderType struct {
	Entry  float64  `json:"entry"`
	Volume float64  `json:"volume"`
	Side   SideType `json:"side"`
}

type PairStatistic struct {
	Orders []*OrderType `json:"orders"`
}

type UserOrderStatistics struct {
	InExchanges map[ExchangeName]map[string]PairStatistic
}

type PairProcessingStatus struct {
	Status         PairStatusType `json:"status"`
	Configurations struct {
		Algorithms []string `json:"algorithms"`
	} `json:"configurations"`
}

// Route keywords.
const (
	ExchangeKWD string = "exchange"
	PairKWD     string = "route"
)
