package eastmoney

import "github.com/iaping/go-stock"

type Eastmoney struct {
	c *stock.Client
}

func New(c *stock.Client) *Eastmoney {
	return &Eastmoney{
		c: c,
	}
}

func Default() *Eastmoney {
	return New(stock.Default())
}

func (e *Eastmoney) json(reqopt stock.Request, v interface{}) error {
	return e.c.Json(reqopt, nil, v)
}

func (e *Eastmoney) Hsj() *Hsj {
	return NewHsj(e)
}

func (e *Eastmoney) Kline() *Kline {
	return NewKline(e)
}
