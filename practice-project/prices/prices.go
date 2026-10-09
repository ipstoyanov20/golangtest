package prices

import (
	"fmt"

	"example.com/practice-project/conversion"
	"example.com/practice-project/iomanager"
)

type TaxInlcudedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxInlcudedPriceJob) LoadData() error {

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

func (job *TaxInlcudedPriceJob) Process(doneChan chan bool, errorChan chan error) {

	err := job.LoadData()

	// errorChan <- errors.New("Test Error")

	if err != nil {
		errorChan <- err
		return
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)

	doneChan <- true	

}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxInlcudedPriceJob {
	return &TaxInlcudedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
