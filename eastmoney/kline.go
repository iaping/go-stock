package eastmoney

import (
	"fmt"
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
	return resp.Data, err
}

type KlineResponse struct {
	Name      string   `json:"name"`
	Code      string   `json:"code"`
	Market    int8     `json:"market"`
	Decimal   int8     `json:"decimal"`
	Dktotal   int      `json:"dktotal"`
	PreKPrice float64  `json:"preKPrice"`
	Klines    []string `json:"klines"`
}
