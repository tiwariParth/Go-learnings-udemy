package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5

	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println(futureValue)
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Value: %.1f\n", futureRealValue)
	// fmt.Printf(`Future value: %.2f
	// Future Real value: %.2f`, futureValue, futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}
