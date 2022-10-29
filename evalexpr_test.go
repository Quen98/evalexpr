package evalexpr

import (
	"fmt"
	"strings"
	"testing"
)

type Test struct {
	Condition      string
	Parameters     []string
	ExpectedOutput bool
}

var tests = []Test{
	Test{"22", []string{""}, false},
	Test{"22", []string{"22"}, true},
	Test{"22&31", []string{"22"}, false},
	Test{"22&31", []string{"31"}, false},
	Test{"22&31", []string{"22", "31"}, true},
	Test{"44|(22&31)", []string{"22", "31"}, true},
	Test{"44|(22&31)", []string{"44"}, true},
	Test{"44|(22&31)", []string{"31"}, false},
	Test{"44&((22&31)|(4&7&9))", []string{"22", "31"}, false},
	Test{"44&((22&31)|(4&7&9))", []string{"44", "22", "31"}, true},
	Test{"44&((22&31)|(4&7&9))", []string{"44", "4", "7", "9"}, true},
}

func TestIsFulfillingCondition(t *testing.T) {
	for index, test := range tests {
		if IsFulfillingCondition(test.Condition, test.Parameters) != test.ExpectedOutput {
			parameters := strings.Join(test.Parameters, ",")
			t.Errorf("Test %d failed, expected %t for condition %s with parameters %s", index, test.ExpectedOutput, test.Condition, parameters)
		}
	}
}

func ExampleIsFulfillingCondition() {
	fmt.Println(IsFulfillingCondition("44", []string{"44"}))
	// Output: true
}
