package helper

import "errors"

func CalculateROP(avgDailyDemand int, leadTimeDays int, safetyStock int) (int, error) {

	if avgDailyDemand < 0 {
		return 0, errors.New("average daily demand cannot be negative")
	}

	if leadTimeDays <= 0 {
		return 0, errors.New("lead time must be greater than zero")
	}

	if safetyStock < 0 {
		return 0, errors.New("safety stock cannot be negative")
	}

	rop := (avgDailyDemand * leadTimeDays) + safetyStock

	return rop, nil
}
