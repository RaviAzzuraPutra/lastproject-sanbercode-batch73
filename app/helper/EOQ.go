package helper

import (
	"errors"
	"math"
)

func CalculateEOQ(annualDemand int, orderCost float64, holdingCost float64) (int, error) {

	if annualDemand <= 0 {
		return 0, errors.New("Annual Demand must be greater than zero")
	}

	if orderCost <= 0 {
		return 0, errors.New("order cost must be greater than zero")
	}

	if holdingCost <= 0 {
		return 0, errors.New("holding cost must be greater than zero")
	}

	eoq := math.Sqrt((2 * float64(annualDemand) * orderCost) / holdingCost)

	return int(math.Round(eoq)), nil
}
