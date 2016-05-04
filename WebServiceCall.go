package main

import (
	"encoding/xml" // Unmaarshall the response body into a well formatted xml to convert to a go object
	"fmt"
	"io/ioutil" // Will give us the ability to read the response from the call
	"net/http"  // Allow us to make the web request from our app
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)

	start := time.Now()

	stockSymbols := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}

	counter := 0

	for _, val := range stockSymbols {
		go func(val string) {

			resp, _ := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=" + val)
			defer resp.Body.Close() //Close when main function finishes

			body, _ := ioutil.ReadAll(resp.Body)
			quote := new(QuoteResponse)
			xml.Unmarshal(body, &quote)

			fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
			counter++
		}(val)
	}

	for counter < (len(stockSymbols)) {
		time.Sleep(10 * time.Millisecond)
	}

	elapsed := time.Since(start)
	fmt.Printf("\nExecution time: %s", elapsed)
}

type QuoteResponse struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChangeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}
