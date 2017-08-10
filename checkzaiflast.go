package main

import (
	"fmt"
	"github.com/yanakend/zaif"
	"log"
)

func main() {

	cslist := map[string]float64{
		"btc_jpy":      float64(0.005),
		"mona_jpy":     float64(1),
		"xem_jpy":      float64(800),
		"pepecash_jpy": float64(4000),
	}
	var total float64

	for s, n := range cslist {
		price, _ := zaif.PublicAPI.LastPrice(s)
		log.Println(s, price.LastPrice)
		it := price.LastPrice * n
		fmt.Println(it, " yen")
		total += it
	}
	fmt.Println(total, " yen")

}
