package main

import (
	"encoding/csv"
	"fmt"       // Take the string and processes to allow us to pull out each field and record
	"io/ioutil" // Quickly read content from the invoice file in the memory
	"os"        // Get access to file system to watch the folder
	"runtime"
	"strconv" // Will help handle non-string fields
	"strings"
	"time" // Conert Unix Timestamp int he invoice file to a time object
)

// Folder directory to watch
const watchedPath = "./sourceInvoice"

func main() {
	runtime.GOMAXPROCS(4)

	fmt.Println("Monitoring folder: ", watchedPath)

	// Infinite loop. Monitor given folder path
	for {
		d, _ := os.Open(watchedPath) // Open the folder
		files, _ := d.Readdir(-1)    // Enumarate folders content. - value will allow us to to get as many files as it filds in the folders

		// Parse through every file inside source folder
		for _, fi := range files {
			filePath := watchedPath + "/" + fi.Name() // New path to each file element found
			f, _ := os.Open(filePath)                 // Open it same as before

			// ByteSlice -> NOT a String
			data, _ := ioutil.ReadAll(f) // Read file content and place it in a data variable
			f.Close()                    // NOT using defer because we would only close it when main is done...and we would leave lots of files open
			os.Remove(filePath)          // Usually we handle the data before removing the file

			// Create a GoRoutine for each data handling
			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data)) // Create reader object with all rows in CSV file

				records, _ := reader.ReadAll() // Turn the reader to a records object

				// Go over records (rows)
				for _, r := range records {
					invoice := new(Invoice) // Since we can't use a marshaller, we will do this manually

					invoice.Number = r[0]
					invoice.Amount, _ = strconv.ParseFloat(r[1], 64)    // 32 or 64 String to float
					invoice.PurchaseOrderNumber, _ = strconv.Atoi(r[2]) // string to int

					unixTime, _ := strconv.ParseInt(r[3], 10, 64) // TimeStamp from a string to an integer (<value>, <base of the number>, <bits>)
					invoice.InvoiceDate = time.Unix(unixTime, 0)  // Create Time objects (<numberOfSecondsSinceUnitEpic>, <fraction of a seconds required>)

					fmt.Printf("Received '%v' for $%.2f and submitted for processing\n", invoice.Number, invoice.Amount)
				}
			}(string(data)) // data variable comes from the ioutil as a ByteSlice

		}
	}
}

// Invoice struct
type Invoice struct {
	Number              string
	Amount              float64
	PurchaseOrderNumber int
	InvoiceDate         time.Time
}
