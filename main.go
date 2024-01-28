package main

import (
	"fmt"

	"example.com/demo_calc/fileManager"
	"example.com/demo_calc/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.15}

	channels := make([]chan bool, len(taxRates))
	errorChannels := make([]chan error, len(taxRates))

	for ix, tr := range taxRates {

		channels[ix] = make(chan bool)

		fm := fileManager.New("ioFiles/prices.txt", fmt.Sprintf("ioFiles/result_%.0f.json", tr*100))
		// cmd := cmdManager.New()
		pricejob := prices.NewTaxIncludedPriceJob(fm, tr)

		go pricejob.Process(channels[ix], errorChannels[ix])

		// if err != nil {
		// 	return
		// }
	}

	for ix := range taxRates {
		select {
		case err := <-errorChannels[ix]:
			if err != nil {
				fmt.Println(err)
			}
		case <-channels[ix]:

		}
	}
}
