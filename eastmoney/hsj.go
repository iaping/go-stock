package eastmoney

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

// 沪深京A股
// https://quote.eastmoney.com/center/gridlist.html#hs_a_board

type Hsj struct {
	client *Eastmoney
	url    string
	page   int
	size   int
}

func NewHsj(client *Eastmoney) *Hsj {
	return &Hsj{
		client: client,
		url:    "https://1.push2.eastmoney.com/api/qt/clist/get?cb=&pn=%d&pz=%d&po=0&np=1&fltt=2&invt=2&dect=1&wbp2u=|0|0|0|web&fid=f12&fs=m:0+t:6,m:0+t:80,m:1+t:2,m:1+t:23,m:0+t:81+s:2048&fields=f12,f13,f14,f100,f102,f103&_=%d",
		page:   1,
		size:   20,
	}
}

// 页码
func (hsj *Hsj) SetPage(i int) *Hsj {
	hsj.page = i
	return hsj
}

// 数量
func (hsj *Hsj) SetSize(i int) *Hsj {
	hsj.size = i
	return hsj
}

func (hsj *Hsj) Do() (*HsjResponse, error) {
	opt := func(req *fasthttp.Request) error {
		req.Header.SetMethod(fasthttp.MethodGet)
		url := fmt.Sprintf(hsj.url, hsj.page, hsj.size, time.Now().UnixMicro())
		req.SetRequestURI(url)
		return nil
	}

	var resp struct {
		Data *HsjResponse `json:"data"`
	}
	err := hsj.client.json(opt, &resp)
	return resp.Data, err
}

type HsjResponse struct {
	Total int `json:"total"`
	Data  []struct {
		Name     string `json:"f14"`  // 名称
		Code     string `json:"f12"`  // 代码
		Market   int8   `json:"f13"`  // 市场
		Industry string `json:"f100"` // 行业
		Zone     string `json:"f102"` // 地域
		Concept  string `json:"f103"` // 概念
	} `json:"diff"`
}
