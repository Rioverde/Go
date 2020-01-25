package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	cmc "github.com/coincircle/go-coinmarketcap"
)

func main() {
	csvfile, ferr := os.Open("data.csv")
	if ferr != nil {
		log.Fatal(ferr)
	}
	r := csv.NewReader(csvfile)

	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range record {
			convertprice(v, "USD")
		}
	}

}

func convertprice(initial string, converted string) {
	price, err := cmc.Price(&cmc.PriceOptions{
		Symbol:  initial,
		Convert: converted,
	})
	if err != nil {
		log.Fatalf("convertPrice: %v", err)
	}

	fmt.Println(price)
}
