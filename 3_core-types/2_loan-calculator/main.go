package main

import (
	"errors"
	"fmt"
)

func calculateLoan(applicantNum int, creditScore int, monthlyIncome float64,
	loanAmount float64, loanTerm int) error {

	if applicantNum < 0 || creditScore < 0 || monthlyIncome < 0 ||
		loanAmount < 0 || loanTerm < 0 {
		return errors.New("inputs cannot be negative")
	}
	if loanTerm%12 != 0 {
		return errors.New("loan term must be divisible by 12")
	}

	var interestRate float64 = 0.2
	if creditScore >= 450 {
		interestRate = 0.15
	}

	monthlyPayment := ((loanAmount * interestRate) + float64(loanAmount)) / float64(loanTerm)
	totalCost := (monthlyPayment * float64(loanTerm)) - float64(loanAmount)

	approved := false
	if creditScore >= 450 && monthlyPayment/monthlyIncome <= 0.2 {
		approved = true
	} else if monthlyPayment/monthlyIncome <= 0.1 {
		approved = true
	}

	fmt.Println("Applicant", applicantNum)
	fmt.Println("-----------")
	fmt.Println("Credit Score   :", creditScore)
	fmt.Println("Monthly Income :", monthlyIncome)
	fmt.Println("Loan Amount    :", loanAmount)
	fmt.Println("Loan Term      :", loanTerm)
	fmt.Println("Monthly Payment:", monthlyPayment)
	fmt.Println("Interest Rate  :", interestRate)
	fmt.Println("Total Cost     :", totalCost)
	fmt.Println("Approved       :", approved)
	fmt.Println("")

	return nil
}

func main() {
	err := calculateLoan(1, 500, 1000, 1000, 24)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = calculateLoan(2, 350, 1000, 10000, 12)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
