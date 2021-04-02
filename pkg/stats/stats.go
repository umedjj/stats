package stats

import(

	"github.com/umedjj/bank/v2/pkg/types" 
)

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {

	sum := types.Money(0)
	number := 0
		
    for _, payment:= range payments  {
		if payment.Status != "FAIL" {
			number++	
			sum += payment.Amount
		} 
		 
    }
	averagepayment := sum / types.Money(number)
	return averagepayment
}
// TotalnCategory находит сумму покупок в определенной категории
func TotalInCategory(payments []types.Payment, category types.Category)types.Money  {
	
	sum := types.Money(0)
	for _, payment:= range payments {
		if payment.Status == "FAIL" {
			continue
		} 
		if payment.Category == category{
			sum += payment.Amount
		} 		
	}
	return sum 
} 