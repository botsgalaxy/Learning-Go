package main

import ( 
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter Years: ")
	fmt.Scan(&years)



	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100,years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100,years)
	fmt.Println("Future Value:",futureValue)
	fmt.Println("Future Real Value:", futureRealValue)


}
