// Package weather is used to forecast the future weather condition
// given current weather condition.
package weather

// CurrentCondition is used to store the current weather condition.
var CurrentCondition string

// CurrentLocation is used to store the current location of which the
// current weather condition is recorded.
var CurrentLocation string

// Forecast is a function that gives the future weather condition for a location
// given the current weather condition for that location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
