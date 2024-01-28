package prices

import (
	"fmt"

	"example.com/demo_calc/conversion"
	"example.com/demo_calc/ioManager"
)

type TaxIncludedPriceJob struct {
	IOManager         ioManager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) Process() error {

	err := job.LoadData()
	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, p := range job.InputPrices {
		taxIncludedPrice := p * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", p)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	return job.IOManager.WriteResult(job)
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}

func NewTaxIncludedPriceJob(iom ioManager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
