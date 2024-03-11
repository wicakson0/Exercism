package cars

import (
	"math"
)

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerHour := float64(productionRate) * (successRate / 100)

	return int(math.Floor(carsPerHour / 60))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	quotient := carsCount / 10
	remainder := carsCount % 10

	return uint(quotient*95000) + uint(remainder*10000)
}
