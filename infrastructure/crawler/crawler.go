package crawler

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/yejingxuan/accumulate/domain/entity"
	"github.com/yejingxuan/accumulate/domain/repository"
	"math/rand"
	"reflect"
	"time"
)

var (
	xueQiuAllDataUrl = "https://xueqiu.com/service/v5/stock/screener/quote/list?page=%d&size=%d&order=desc&orderby=percent&order_by=percent&market=CN&type=sh_sz&_=%d"
	pageNo           = 1
	pageSize         = 200
	total            = 200
	totalPage        = 1
)

func ExecXueQiuJob(stockRepo repository.StockRepo) error {
	c := colly.NewCollector()

	/*c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		//配置代理
		c.SetProxyFunc(randomProxySwitcher())
		//爬取子页
		_ = c.Visit(e.Request.AbsoluteURL(link))
	})*/
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Body)
		//配置代理
		c.SetProxyFunc(randomProxySwitcher())
		resp := XueQiuAllDataResp{}
		json.Unmarshal(r.Body, &resp)
		total = resp.Data.Count
		totalPage = total / pageSize

		for _, item := range resp.Data.List {
			stock := entity.Stock{}
			structAssign(&stock, &item)
			stockRepo.CreateStock(&stock)
		}

		if pageNo < totalPage {
			pageNo++
			time.Sleep(5 * time.Second)
			c.Visit(fmt.Sprintf(xueQiuAllDataUrl, pageNo, pageSize, time.Now().UnixNano()))
		}

	})

	//爬取根网页
	baseUrl := fmt.Sprintf(xueQiuAllDataUrl, pageNo, pageSize, time.Now().UnixNano())
	_ = c.Visit(baseUrl)
	return nil
}

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

//binding type interface 要修改的结构体
//value type interace 有数据的结构体
func structAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}

type XueQiuAllDataResp struct {
	Data struct {
		Count int `json:"count"`
		List  []struct {
			Symbol             string  `json:"symbol"`
			NetProfitCagr      float64 `json:"net_profit_cagr"`
			NorthNetInflow     string  `json:"north_net_inflow"`
			Ps                 float64 `json:"ps"`
			Type               int     `json:"type"`
			Percent            float64 `json:"percent"`
			HasFollow          bool    `json:"has_follow"`
			TickSize           float64 `json:"tick_size"`
			PbTtm              float64 `json:"pb_ttm"`
			FloatShares        int     `json:"float_shares"`
			Current            float64 `json:"current"`
			Amplitude          float64 `json:"amplitude"`
			Pcf                float64 `json:"pcf"`
			CurrentYearPercent float64 `json:"current_year_percent"`
			FloatMarketCapital int64   `json:"float_market_capital"`
			NorthNetInflowTime string  `json:"north_net_inflow_time"`
			MarketCapital      int64   `json:"market_capital"`
			DividendYield      int     `json:"dividend_yield"`
			LotSize            int     `json:"lot_size"`
			RoeTtm             float64 `json:"roe_ttm"`
			TotalPercent       float64 `json:"total_percent"`
			Percent5M          int     `json:"percent5m"`
			IncomeCagr         float64 `json:"income_cagr"`
			Amount             int     `json:"amount"`
			Chg                float64 `json:"chg"`
			IssueDateTs        int64   `json:"issue_date_ts"`
			Eps                float64 `json:"eps"`
			MainNetInflows     int     `json:"main_net_inflows"`
			Volume             int     `json:"volume"`
			VolumeRatio        float64 `json:"volume_ratio"`
			Pb                 float64 `json:"pb"`
			Followers          int     `json:"followers"`
			TurnoverRate       float64 `json:"turnover_rate"`
			FirstPercent       float64 `json:"first_percent"`
			Name               string  `json:"name"`
			PeTtm              float64 `json:"pe_ttm"`
			TotalShares        int     `json:"total_shares"`
			LimitupDays        int     `json:"limitup_days"`
		} `json:"list"`
	} `json:"data"`
}
