package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {

	payment := flag.Float64("payment", 0, "Is the payment amount")
	principal := flag.Float64("principal", 0, "You can get its value if you know the interest, annuity payment, and number of months.")
	periods := flag.Int64("periods", 0, "Denotes the number of months needed to repay the loan.")
	interest := flag.Float64("interest", 0, "Is specified without a percent sign.")
	typeFlag := flag.String("type", "", "Is indicating the type of payment: \"annuity\" or \"diff\" (differentiated)")

	flag.Parse()

	if *typeFlag != "annuity" && *typeFlag != "diff" {
		fmt.Println("Incorrect parameters.")
		return
	}

	if *periods == 0 {
		if *typeFlag == "annuity" {

			if float64(*payment) <= 0 || float64(*principal) <= 0 || float64(*interest) <= 0 {
				fmt.Println("Incorrect parameters.")
				return
			}

			p := numberOfPayments(*payment, *principal, *interest)
			overPayment := calculateOverPayment(*payment, *principal, p)

			years := p / 12
			months := p % 12

			titleYear := "year"
			titleMonth := "month"

			if years > 1 {
				titleYear += "s"
			}

			if months > 1 {
				titleMonth += "s"
			}

			fmt.Printf("It will take %d %s ", years, titleYear)
			if months > 0 {
				fmt.Printf("and %d %s ", months, titleMonth)
			}
			fmt.Printf("to repay this loan!\n")
			fmt.Printf("Overpayment = %d\n", int(overPayment))

		} else {
			fmt.Println("Incorrect parameters.")
		}
	} else if *payment == 0 {

		if *periods <= 0 || float64(*principal) <= 0 || float64(*interest) <= 0 {
			fmt.Println("Incorrect parameters.")
			return
		}

		if *typeFlag == "diff" {

			payments, overPayment := differentiatedPayment(*principal, *interest, int(*periods))

			for i := 1; i <= len(payments); i++ {
				fmt.Printf("Month %d: payment is %d\n", i, int(payments[i-1]))
			}

			fmt.Printf("\nOverpayment = %d", int(overPayment))
		} else {
			monthlyPayment := annuityPayment(*principal, *interest, int(*periods))

			overPayment := calculateOverPayment(math.Ceil(monthlyPayment), *principal, int(*periods))
			fmt.Printf("Your annuity payment = %d!\n", int(math.Ceil(monthlyPayment)))
			fmt.Printf("Overpayment = %d\n", int(overPayment))

		}
	} else if *principal == 0 {
		if *typeFlag == "annuity" {

			if *payment <= 0 || *periods <= 0 || *interest <= 0 {

				fmt.Println("Incorrect parameters.")
				return
			}

			loanPrincipal := loanPrincipal(*payment, *interest, int(*periods))

			fmt.Printf("Your loan principal = %d!\n", int(math.Floor(loanPrincipal)))

			overPayment := calculateOverPayment(*payment, loanPrincipal, int(*periods))
			fmt.Printf("Overpayment = %d!\n", int(math.Ceil(overPayment)))

		} else {
			fmt.Println("Incorrect parameters.")
		}
	} else {
		fmt.Println("Incorrect parameters.")
	}

}

func annuityPayment(principal, interest float64, periods int) float64 {
	i := interest / 1200
	return principal * i * math.Pow(1+i, float64(periods)) / (math.Pow(1+i, float64(periods)) - 1)
}

func calculateOverPayment(payment, principal float64, periods int) float64 {
	totalPayments := payment * float64(periods)
	return totalPayments - principal
}

func numberOfPayments(payment, principal, interest float64) int {
	i := interest / 1200
	return int(math.Ceil(math.Log(payment/(payment-i*principal)) / math.Log(1+i)))
}

func loanPrincipal(payment, interest float64, periods int) float64 {
	return payment / calculateAmortizationFactor(interest, periods)
}

func calculateAmortizationFactor(interesRate float64, numberOfPayments int) float64 {
	i := interes(interesRate)
	j := math.Pow(1+i, float64(numberOfPayments))

	return (i * j) / (j - 1)
}

func differentiatedPayment(principal, interest float64, periods int) ([]float64, float64) {
	i := interest / 1200

	payments := make([]float64, periods)
	totalPayments := 0.0

	for m := 1; m <= periods; m++ {
		payment := (principal / float64(periods)) + i*(principal-float64(m-1)*(principal/float64(periods)))
		payments[m-1] = math.Ceil(payment)
		totalPayments += payments[m-1]
	}

	overpayment := totalPayments - principal
	return payments, overpayment
}

func interes(interesRate float64) float64 {
	return interesRate / (12 * 100)
}
