package eastmoney

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

// K线
// https://quote.eastmoney.com/concept/sz000001.html

type Kline struct {
	c     *Eastmoney
	url   string
	pz    int
	secid string
	end   string
}

func NewKline(c *Eastmoney) *Kline {
	return &Kline{
		c:   c,
		url: "https://push2his.eastmoney.com/api/qt/stock/kline/get?secid=%s&fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61&klt=101&fqt=1&end=%s&lmt=%d",
		pz:  210,
		end: time.Now().Format("20060102"),
	}
}

// 0.000001
func (k *Kline) SetSecid(i string) *Kline {
	k.secid = i
	return k
}

// 数量
func (k *Kline) SetSize(i int) *Kline {
	k.pz = i
	return k
}

func (k *Kline) Do() (*KlineResponse, error) {
	opt := func(req *fasthttp.Request) error {
		req.Header.SetMethod(fasthttp.MethodGet)
		url := fmt.Sprintf(k.url, k.secid, k.end, k.pz)
		req.SetRequestURI(url)
		return nil
	}

	var resp struct {
		Data *KlineResponse `json:"data"`
	}
	err := k.c.json(opt, &resp)
	if err != nil {
		return nil, err
	}

	if err = resp.Data.format(); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

type KlineResponse struct {
	Name      string       `json:"name"`
	Code      string       `json:"code"`
	Market    int8         `json:"market"`
	Decimal   int8         `json:"decimal"`
	Dktotal   int          `json:"dktotal"`
	PreKPrice float64      `json:"preKPrice"`
	Items     []*KlineItem `json:"items"`
	Klines    []string     `json:"klines"`
}

func (kr *KlineResponse) format() error {
	var items []*KlineItem
	for _, i := range kr.Klines {
		j := strings.Split(i, ",")
		if len(j) != 11 {
			return fmt.Errorf("error in kline data length")
		}

		open, err := strconv.ParseFloat(j[1], 64)
		if err != nil {
			return err
		}

		close, err := strconv.ParseFloat(j[2], 64)
		if err != nil {
			return err
		}

		high, err := strconv.ParseFloat(j[3], 64)
		if err != nil {
			return err
		}

		low, err := strconv.ParseFloat(j[4], 64)
		if err != nil {
			return err
		}

		vol, err := strconv.ParseInt(j[5], 10, 64)
		if err != nil {
			return err
		}

		amount, err := strconv.ParseFloat(j[6], 64)
		if err != nil {
			return err
		}

		vola, err := strconv.ParseFloat(j[7], 64)
		if err != nil {
			return err
		}

		rose, err := strconv.ParseFloat(j[8], 64)
		if err != nil {
			return err
		}

		updown, err := strconv.ParseFloat(j[9], 64)
		if err != nil {
			return err
		}

		turnover, err := strconv.ParseFloat(j[10], 64)
		if err != nil {
			return err
		}

		items = append(items, &KlineItem{
			Date:     j[0],
			Open:     open,
			Close:    close,
			High:     high,
			Low:      low,
			Vol:      vol,
			Amount:   amount,
			Vola:     vola,
			Rose:     rose,
			Updown:   updown,
			Turnover: turnover,
		})
	}
	kr.Items = items
	return nil
}

type KlineItem struct {
	Date     string  `json:"date"`     // 日期
	Open     float64 `json:"open"`     // 开盘价
	Close    float64 `json:"close"`    // 收盘价
	High     float64 `json:"high"`     // 最高价
	Low      float64 `json:"low"`      // 最低价
	Vol      int64   `json:"vol"`      // 总手数
	Amount   float64 `json:"amount"`   // 总金额
	Vola     float64 `json:"vola"`     // 振幅
	Rose     float64 `json:"rose"`     // 涨幅
	Updown   float64 `json:"updown"`   // 涨跌
	Turnover float64 `json:"turnover"` // 换手率
}
