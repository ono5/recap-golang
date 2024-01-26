package main

import (
	"example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.NewFilemanger("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.NewCMDManger()
		pricesJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		pricesJob.Process()
	}
}
