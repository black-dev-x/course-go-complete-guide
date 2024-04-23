package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000.0
	var expectedReturnRate = 5.5
	var years = 10.0
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100.0, years)
	fmt.Print(futureValue)
}
