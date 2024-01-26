package cmdmanager

import "fmt"

type CMDManger struct{}

func (fm CMDManger) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (fm CMDManger) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func NewCMDManger() CMDManger {
	return CMDManger{}
}
