package main

import "fmt"

type taxLevel struct {
	Key int32
	Val int32
}

func getTaxConfig() []taxLevel{
	var taxConfig []taxLevel
	taxConfig = append(taxConfig,  taxLevel{50000000, 5})
	taxConfig = append(taxConfig,  taxLevel{250000000, 15})
	taxConfig = append(taxConfig,  taxLevel{500000000, 25})
	taxConfig = append(taxConfig,  taxLevel{2147483647, 30})
	return taxConfig
}

func calculateTax(income int32) int32 {
	var taxLevelBefore, taxVariable int32
	restIncome := income
	var totalTax int32
	for _, v := range getTaxConfig(){
		taxGrade := v.Key
		taxMultiplier := v.Val
		if restIncome > taxGrade {
			taxVariable = taxMultiplier * (taxGrade/100)
			fmt.Println("[1]",taxMultiplier, "persen dikali dengan", (taxGrade),"=", taxVariable)
		} else {
			taxVariable = taxMultiplier * (restIncome/100)
			fmt.Println("[2]",taxMultiplier, "persen dikali dengan", (restIncome),"=", taxVariable)
		}
		totalTax += taxVariable
		taxLevelBefore = taxGrade
		restIncome -= taxLevelBefore
		if restIncome <= 0 {
			break
		}
	}
	return totalTax
}

func main() {
	var income int32
	_, err := fmt.Scanln(&income)
	if err != nil {
		fmt.Print(err)
	}
	tax := calculateTax(income)
	fmt.Println(tax)
}
