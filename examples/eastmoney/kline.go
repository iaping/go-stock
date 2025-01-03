package main

import (
	"fmt"

	"github.com/iaping/go-stock/eastmoney"
)

func main() {
	kline := eastmoney.Default().Kline()
	kline.SetSecid("0.000001")
	kline.SetPeriod(eastmoney.KlinePeriodDay)
	fmt.Println(kline.Do())
}
