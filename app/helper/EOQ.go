package helper

import (
	"math"
)

func CalculateEOQ(annualDemand int, orderCost float64, holdingCost float64) (int, error) {

	if annualDemand <= 0 {
		return 0, NewBadRequest("Annual Demand must be greater than zero")
	}

	if orderCost <= 0 {
		return 0, NewBadRequest("order cost must be greater than zero")
	}

	if holdingCost <= 0 {
		return 0, NewBadRequest("holding cost must be greater than zero")
	}

	eoq := math.Sqrt((2 * float64(annualDemand) * orderCost) / holdingCost)

	return int(math.Round(eoq)), nil
}
