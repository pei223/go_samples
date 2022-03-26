package struct_sample

import (
	"fmt"
	"testing"
)

func TestSetToArgFunc(t *testing.T) {
	origin := TestStruct{Title: "a", Count: 1}
	afterSetToArg := SetToArg(origin)

	if origin.Count == afterSetToArg.Count {
		t.Errorf("origin.Count != afterSetToArg.Count:  %d, %d", origin.Count, afterSetToArg.Count)
	}
}

func BenchmarkSetToArgFunc(b *testing.B) {
	origin := TestStruct{Title: "a", Count: 1}
	for i := 0; i < b.N; i++ {
		SetToArg(origin)
	}
}

// サンプル関数
func ExampleSetToArg() {
	origin := TestStruct{Title: "a", Count: 1}
	afterSetToArg := SetToArg(origin)
	fmt.Println(afterSetToArg)
}
