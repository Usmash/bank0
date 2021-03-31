  
package stats

import "github.com/Usmash/bank/pkg/bank/types"

// Avg calculates average payment
func Avg(payments []types.Payment) types.Money {
	count := types.Money(0)
	average := types.Money(0)
	for _, payment := range payments {
		average += payment.Amount
		count++
	}

	return average / count
}