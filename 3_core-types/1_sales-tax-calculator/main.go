package main

import "fmt"

func getSalesTax(cost float64, salesTaxRate float64) float64 {
	return cost * salesTaxRate
}

func main() {
	var salesTaxTotal float64
	salesTaxTotal += getSalesTax(0.99, 0.075)
	salesTaxTotal += getSalesTax(2.75, 0.015)
	salesTaxTotal += getSalesTax(0.87, 0.02)
	fmt.Println("Sales Tax Total:", salesTaxTotal)
}
