package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    var result float64
    result = float64(productionRate) * successRate / 100
    return result
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    workingPerHour:=CalculateWorkingCarsPerHour(productionRate,successRate)
	var result int 
    result = int(workingPerHour)/60
    return result
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	tens:= carsCount/10
    ones:= carsCount%10

    var cost int
    cost = (tens * 95000) + (ones*10000)

    return uint(cost)
}
