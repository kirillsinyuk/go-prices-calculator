package prices

import (
	"fmt"
	"price-calculator/conversion"
	"price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64           `json:"tax_rate"`
	InputPrices       []float64         `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
	ioManager         filemanager.FileManager
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.ioManager.ReadLines()
	if err != nil {
		panic(err)
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		panic(err)
	}

	job.InputPrices = prices

}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	job.ioManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(fileManager *filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		ioManager:   *fileManager,
	}
}
