package entity

//附件信息结构
type Stock struct {
	Symbol             string  `gorm:"column:symbol" json:"symbol"`
	Name               string  `gorm:"column:name" json:"name"`
	NetProfitCagr      float64 `gorm:"column:net_profit_cagr" json:"net_profit_cagr"`
	NorthNetInflow     string  `gorm:"column:north_net_inflow" json:"north_net_inflow"`
	Ps                 float64 `gorm:"column:ps" json:"ps"`
	Type               int     `gorm:"column:type" gorm:"column:" json:"type"`
	Percent            float64 `gorm:"column:percent" json:"percent"`
	HasFollow          bool    `gorm:"column:has_follow" json:"has_follow"`
	TickSize           float64 `gorm:"column:tick_size" json:"tick_size"`
	PbTtm              float64 `gorm:"column:pb_ttm" json:"pb_ttm"`
	FloatShares        int     `gorm:"column:float_shares" json:"float_shares"`
	Current            float64 `gorm:"column:current" json:"current"`
	Amplitude          float64 `gorm:"column:amplitude" json:"amplitude"`
	Pcf                float64 `gorm:"column:pcf" json:"pcf"`
	CurrentYearPercent float64 `gorm:"column:current_year_percent" json:"current_year_percent"`
	FloatMarketCapital int64   `gorm:"column:float_market_capital" json:"float_market_capital"`
	NorthNetInflowTime string  `gorm:"column:north_net_inflow_time" json:"north_net_inflow_time"`
	MarketCapital      int64   `gorm:"column:market_capital" json:"market_capital"`
	DividendYield      int     `gorm:"column:dividend_yield" json:"dividend_yield"`
	LotSize            int     `gorm:"column:lot_size" json:"lot_size"`
	RoeTtm             float64 `gorm:"column:roe_ttm" json:"roe_ttm"`
	TotalPercent       float64 `gorm:"column:total_percent" json:"total_percent"`
	Percent5M          int     `gorm:"column:percent5m" json:"percent5m"`
	IncomeCagr         float64 `gorm:"column:income_cagr" json:"income_cagr"`
	Amount             int     `gorm:"column:amount" json:"amount"`
	Chg                float64 `gorm:"column:chg" json:"chg"`
	IssueDateTs        int64   `gorm:"column:issue_date_ts" json:"issue_date_ts"`
	Eps                float64 `gorm:"column:eps" json:"eps"`
	MainNetInflows     int     `gorm:"column:main_net_inflows" json:"main_net_inflows"`
	Volume             int     `gorm:"column:volume" json:"volume"`
	VolumeRatio        float64 `gorm:"column:volume_ratio" json:"volume_ratio"`
	Pb                 float64 `gorm:"column:pb" json:"pb"`
	Followers          int     `gorm:"column:followers" json:"followers"`
	TurnoverRate       float64 `gorm:"column:turnover_rate" json:"turnover_rate"`
	FirstPercent       float64 `gorm:"column:first_percent" json:"first_percent"`
	PeTtm              float64 `gorm:"column:pe_ttm" json:"pe_ttm"`
	TotalShares        int     `gorm:"column:total_shares" json:"total_shares"`
	LimitupDays        int     `gorm:"column:limitup_days" json:"limitup_days"`
}

var TableStock = "stock"
