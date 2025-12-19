package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmdm CMDManager) ReadLines() ([]string, error) {
	fmt.Println("ReadLines")

	var prices = make([]string, 0)
	for {
		var price string
		fmt.Println("Price column:")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (cmdm CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() *CMDManager {
	return &CMDManager{}
}
