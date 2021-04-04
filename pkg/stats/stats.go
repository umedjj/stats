package stats

import(

	"github.com/umedjj/bank/v2/pkg/types" 
)

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money  {
	dynamic:= map[types.Category]types.Money{} 
		for scndCategory, scndPay := range second {
			dynamic[scndCategory]=scndPay-first[scndCategory]			
	 	}
	 for frstCategory, frstPay := range first {
		 for _, _ = range dynamic {
			_, ok:=dynamic[frstCategory]
			if !ok{
				dynamic[frstCategory]=0-frstPay
			}
		 }
				
			}	
	return dynamic	 
 }