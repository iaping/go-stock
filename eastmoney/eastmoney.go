package eastmoney

import "github.com/iaping/go-stock"

type Eastmoney struct {
	client *stock.Client
}

func New(client *stock.Client) *Eastmoney {
	return &Eastmoney{
		client: client,
	}
}

func Default() *Eastmoney {
	return New(stock.Default())
}

func (e *Eastmoney) json(reqopt stock.Request, v interface{}) error {
	return e.client.Json(reqopt, nil, v)
}

func (e *Eastmoney) Hsj() *Hsj {
	return NewHsj(e)
}

func (e *Eastmoney) Kline() *Kline {
	return NewKline(e)
}
