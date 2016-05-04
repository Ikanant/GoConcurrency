package main

import (
	"encoding/csv"
	"fmt"       // Take the string and processes to allow us to pull out each field and record
	"io/ioutil" // Quickly read content from the invoice file in the memory
	"os"        // Get access to file system to watch the folder
	"strconv"   // Will help handle non-string fields
	"strings"
	"time" // Conert Unix Timestamp int he invoice file to a time object
)

const watchedPath = "./sourceInvoice"

func main() {
	// Infinite loop
	fmt.Println("Monitoring folder: ", watchedPath)
	for {
		d, _ := os.Open(watchedPath)
		files, _ := d.Readdir(-1) // Enumarate folders content. - value will allow us to to get as many files as it filds in the folders

		for _, fi := range files {
			filePath := watchedPath + "/" + fi.Name()
			f, _ := os.Open(filePath)
			data, _ := ioutil.ReadAll(f)
			f.Close()
			os.Remove(filePath)

			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data))

				records, _ := reader.ReadAll()

				for _, r := range records {
					invoice := new(Invoice)

					invoice.Number = r[0]
					invoice.Amount, _ = strconv.ParseFloat(r[1], 64) // 32 or 64
					invoice.PurchaseOrderNumber, _ = strconv.Atoi(r[2])

					unixTime, _ := strconv.ParseInt(r[3], 10, 64)
					invoice.InvoiceDate = time.Unix(unixTime, 0)

					fmt.Printf("Received")
				}
			}(string(data))

		}
	}
}

type Invoice struct {
	Number              string
	Amount              float64
	PurchaseOrderNumber int
	InvoiceDate         time.Time
}
