package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, numCows int) (float64, error) {
	totalFodder, errAmount := fc.FodderAmount(numCows)
	fattenFactor, errFactor := fc.FatteningFactor()

	if errAmount != nil {
		return 0, errAmount
	}

	if errFactor != nil {
		return 0, errFactor
	}

	return (totalFodder / float64(numCows)) * fattenFactor, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error) {
	if numCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	foodPerCow, err := DivideFood(fc, numCows)

	if err != nil {
		return 0, err
	}

	return foodPerCow, nil
}

// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
	numCow        int
	customMessage string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numCow, e.customMessage)
}

func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{
			numCow:        numCows,
			customMessage: "there are no negative cows",
		}
	}
	if numCows == 0 {
		return &InvalidCowsError{
			numCow:        0,
			customMessage: "no cows don't need food",
		}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
