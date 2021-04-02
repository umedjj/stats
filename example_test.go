package stats
import(

	"github.com/umedjj/bank/pkg/types"
	
	"fmt"
	
)

func ExampleAvg()  {
	payments:= [] types.Payment{
	{
		ID:       1,
	    Amount:   10_000_00,
	    Category: "Avto",
	},
	{
		ID:       2,
	    Amount:   6_000_00,
	    Category: "Avto",
	},
	{
		ID:       3,
	    Amount:   5_000_00,
	    Category: "Avto",
	},
	}
	fmt.Println(Avg(payments))

	// Output: 700000
}
func ExampleTotalInCategory()  {
	payments:= [] types.Payment{
	{
		ID:       1,
		Amount:   10_000_00,
		Category: "Avto",
	},
	{
		ID:       2,
		Amount:   6_000_00,
		Category: "Avto",
	},
	{
		ID:       3,
		Amount:   5_000_00,
		Category: "Avto",
	},
		
	{
	    ID:       1,
		Amount:   7_000_00,
		Category: "Phone",
	},
	{
		ID:       2,
		Amount:   4_000_00,
		Category: "Phone",
	},
	{
		ID:       1,
		Amount:   8_000_00,
		Category: "Pool",
	},
	}

	fmt.Println(TotalInCategory(payments,"Phone"))

	// Output: 1100000	
}