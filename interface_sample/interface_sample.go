package interface_sample

import "fmt"

type SampleIF interface {
	Print()
	Add(message string)
}

type SampleIFImpl struct {
	message string
}

func (t1 SampleIFImpl) Print() {
	fmt.Println("T1 Print: ", t1.message)
}
func (t1 *SampleIFImpl) Add(message string) {
	t1.message += message
}

type SampleIFIncorrectImpl struct {
	message string
}

func (t2 SampleIFIncorrectImpl) Print() {
	fmt.Println("T2 Print: ", t2.message)
}

func DoInterfaceSample() {
	fmt.Println("\n\nインターフェース型テスト")
	st1 := SampleIFImpl{message: "t1"}
	// Setterを呼ぶため参照を渡す必要がある
	DoIfFuncs(&st1)
	// これは実行できない
	// t2 := T2{message: "t2"}
	// ifFunc(t2)
}

func DoIfFuncs(testImple SampleIF) {
	testImple.Add("xhoge")
	testImple.Print()
}
