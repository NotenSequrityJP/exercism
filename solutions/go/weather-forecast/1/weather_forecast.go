// Package weather provides weather forecasting functionalities for the Goblinocus country.
package weather

var (
	// CurrentCondition holds the current weather condition as a string.
	CurrentCondition string
	// CurrentLocation holds the name of the current location as a string.
	CurrentLocation string
)

// Forecast returns a weather forecast string for the given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
