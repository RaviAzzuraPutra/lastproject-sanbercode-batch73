package helper

func CalculateROP(avgDailyDemand int, leadTimeDays int, safetyStock int) (int, error) {

	if avgDailyDemand < 0 {
		return 0, NewBadRequest("average daily demand cannot be negative")
	}

	if leadTimeDays <= 0 {
		return 0, NewBadRequest("lead time must be greater than zero")
	}

	if safetyStock < 0 {
		return 0, NewBadRequest("safety stock cannot be negative")
	}

	rop := (avgDailyDemand * leadTimeDays) + safetyStock

	return rop, nil
}
