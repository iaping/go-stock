package main

import (
	"fmt"

	"github.com/iaping/go-stock/eastmoney"
)

func main() {
	hsj := eastmoney.Default().Hsj()
	hsj.SetPage(1)
	hsj.SetSize(100)
	fmt.Println(hsj.Do())
}
