package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var i float32
	if balance < 0 {
		i = 3.213
	} else if balance < 1000 {
		i = 0.5
	} else if balance >= 1000 && balance < 5000 {
		i = 1.621
	} else {
		i = 2.475
	}
	return i

}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return (float64(InterestRate(balance) * float32(balance))) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var y int
	for ; balance < targetBalance; y++ {
		balance = AnnualBalanceUpdate(balance)
	}
	return y
}
