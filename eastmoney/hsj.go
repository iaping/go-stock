package eastmoney

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

// 沪深京A股
// https://quote.eastmoney.com/center/gridlist.html#hs_a_board

type Hsj struct {
	c   *Eastmoney
	url string
	pn  int
	pz  int
}

func NewHsj(c *Eastmoney) *Hsj {
	return &Hsj{
		c:   c,
		url: "https://1.push2.eastmoney.com/api/qt/clist/get?cb=&pn=%d&pz=%d&po=0&np=1&fltt=2&invt=2&dect=1&wbp2u=|0|0|0|web&fid=f12&fs=m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048&fields=f12,f13,f14&_=%d",
		pn:  1,
		pz:  20,
	}
}

// 页码
func (hsj *Hsj) SetPage(i int) *Hsj {
	hsj.pn = i
	return hsj
}

// 数量
func (hsj *Hsj) SetSize(i int) *Hsj {
	hsj.pz = i
	return hsj
}

func (hsj *Hsj) Do() (*HsjResponse, error) {
	opt := func(req *fasthttp.Request) error {
		req.Header.SetMethod(fasthttp.MethodGet)
		url := fmt.Sprintf(hsj.url, hsj.pn, hsj.pz, time.Now().UnixMicro())
		req.SetRequestURI(url)
		return nil
	}

	var resp struct {
		Data *HsjResponse `json:"data"`
	}
	err := hsj.c.json(opt, &resp)
	return resp.Data, err
}

type HsjResponse struct {
	Total int `json:"total"`
	Data  []struct {
		Name   string `json:"f14"`
		Code   string `json:"f12"`
		Market int8   `json:"f13"`
	} `json:"diff"`
}
