// Package weather for weatherforecast.
package weather

var (
    // CurrentCondition , current weather condition.
	CurrentCondition string
    // CurrentLocation, location.
	CurrentLocation  string
)
// Forecast , hhh .
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
