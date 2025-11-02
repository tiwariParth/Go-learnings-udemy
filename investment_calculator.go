package main

// import (
// 	"fmt"
// 	"math"
// )

// const inflationRate = 2.5

// func main() {
// 	var investmentAmount float64
// 	var years float64
// 	var expectedReturnRate float64

// 	outputText("Enter your investment amount: ")
// 	fmt.Scan(&investmentAmount)

// 	outputText("Years: ")
// 	fmt.Scan(&years)

// 	outputText("Expected Return Rate: ")
// 	fmt.Scan(&expectedReturnRate)

// 	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

// 	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

// 	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

// 	// fmt.Println(futureValue)
// 	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
// 	formattedRFV := fmt.Sprintf("Future Real Value: %.1f\n", futureRealValue)
// 	// fmt.Printf(`Future value: %.2f
// 	// Future Real value: %.2f`, futureValue, futureRealValue)
// 	fmt.Print(formattedFV, formattedRFV)
// }

// func outputText(text string) {
// 	fmt.Print(text)
// }

// func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
// 	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv = fv / math.Pow(1+inflationRate/100, years)
// 	return fv, rfv
// }
