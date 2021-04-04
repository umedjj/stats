package stats

import (
	"reflect"
	"testing"

	"github.com/umedjj/bank/v2/pkg/types"
)
func TestPediodsDynamic(t *testing.T) {
	first:= map[types.Category]types.Money{
		"auto":		10,
		"food":		20,
	}
	second:= map[types.Category]types.Money{
		"food":		20,
	}

	expected:=map[types.Category]types.Money{
		"auto":		-10,
		"food":		0,
	}

	result:=PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}