// Package weather contains forecasts.
package weather

// CurrentCondition: The current weather.
var CurrentCondition string
// CurrentLocation: The current location.
var CurrentLocation string

// Forecast: forecasts the weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
