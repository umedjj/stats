package stats

import (
	"reflect"
	"testing"

	"github.com/umedjj/bank/v2/pkg/types"
)

func TestAvg(t *testing.T) {
	payments:=[]types.Payment{
		{ID:1, Category:"auto", Amount: 1_000_00},
		{ID:2, Category:"fon", Amount: 1_000_00},
		{ID:3, Category:"fon", Amount: 1_000_00},
		{ID:4, Category:"auto", Amount: 1_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 1_000_00,
		"fon": 1_000_00,
	}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}