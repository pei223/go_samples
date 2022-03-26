package struct_sample

import (
	"fmt"
)

// テスト用構造体
type TestStruct struct {
	Title string
	Count int8
}

// 構造体に値をセットする
// 引数はポインタではないため保持されない
func SetToArg(t TestStruct) TestStruct {
	t.Count = 100
	return t
}

/*
構造体をメソッドで値をセットする。
ポインタレシーバではないため代入はできない
*/
func (t TestStruct) IncorrectSetMethod() TestStruct {
	t.Count = 99
	return t
}

func (t *TestStruct) CorrectSetWithMethod() {
	t.Count = 99
}

func DoStructSamples() {
	fmt.Println("\n\n\n構造体テスト")
	origin := TestStruct{
		Title: "aaa",
	}
	fmt.Println("元の構造体: ", origin)
	afterSetToArgFunc := SetToArg(origin)
	fmt.Println("構造体を引数にした関数で導入した結果の戻り値: ", afterSetToArgFunc)
	fmt.Println("構造体を引数にした関数で導入した後の元の構造体: ", origin)

	fmt.Println("構造体を引数にした関数で導入した結果の戻り値 == 構造体を引数にした関数で導入した後の元の構造体: ", origin == afterSetToArgFunc)

	afterIncorrectSet := origin.IncorrectSetMethod()

	// ポインタレシーバではない場合、構造体はコピーされる
	fmt.Println("ポインタレシーバを使わずに構造体の値変更した結果の戻り値: ", afterIncorrectSet)
	fmt.Println("ポインタレシーバを使わずに構造体の値変更した後の元の構造体: ", origin)

	afterCorrectSet := origin
	afterCorrectSet.CorrectSetWithMethod()

	fmt.Println("ポインタレシーバを使って構造体の値変更した後の元の構造体: ", afterCorrectSet)

	fmt.Println("値更新した後のコピー元の構造体の値: ", origin)
	fmt.Println("コピー先の構造体 == コピー元の構造体: ", origin == afterCorrectSet)

	fmt.Println("ポインタレシーバを使わずに構造体の値変更した後の元の構造体 == ポインタレシーバを使わずに構造体の値変更した結果の戻り値: ", origin == afterIncorrectSet)
}
