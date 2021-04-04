package stats

import(

	"github.com/umedjj/bank/v2/pkg/types" 
)

// Avg рассчитывает среднюю сумму платежа
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	count := map[types.Category]types.Money{}
	result := map[types.Category]types.Money{}
  
	for _, payment := range payments {
	  result[payment.Category] += payment.Amount
	  count[payment.Category]++
	}
  
	for key := range result {
	  result[key] /= count[key]
	}
  
	return result
  }