package main

import (
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		pricesJob := prices.NewTaxIncludedPriceJob(taxRate)
		pricesJob.Process()
	}

}
