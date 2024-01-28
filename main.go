package main

import (
	"fmt"

	"example.com/demo_calc/fileManager"
	"example.com/demo_calc/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.15}

	for _, tr := range taxRates {
		fm := fileManager.New("ioFiles/prices.txt", fmt.Sprintf("ioFiles/result_%.0f.json", tr*100))
		// cmd := cmdManager.New()
		pricejob := prices.NewTaxIncludedPriceJob(fm, tr)
		err := pricejob.Process()
		if err != nil {
			return
		}
	}

}
