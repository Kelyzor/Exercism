// Package weather provides info about the weather in current location.
package weather

var (
    // CurrentCondition represents a current condition.
	CurrentCondition string
    // CurrentLocation represents a current city.
	CurrentLocation  string
)

// Forecast returns a string value equal to the weather condition in current city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
