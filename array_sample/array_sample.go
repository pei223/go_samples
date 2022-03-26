package array_sample

import (
	"fmt"
)

func DoArraySample() {
	fmt.Println("\n\n\n配列テスト")
	origin := []int8{1, 2, 3}
	fmt.Println("元の配列: ", origin)
	setV := SetAndReturnAppendedArray(origin)
	fmt.Println("Appendした結果の配列戻り値: ", setV)
	fmt.Println("関数で配列に値設定した後の配列: ", origin)
}

/*
配列に値代入して, appendした配列を返却する
appendは新たに配列が生成されるため、元の配列に影響与えない
*/
func SetAndReturnAppendedArray(b []int8) []int8 {
	c := append(b, 100)
	b[0] = 122
	return c
}
