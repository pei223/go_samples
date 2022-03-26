package goroutine_sample

import (
	"fmt"
	"time"
)

func DoGoroutineSample() {
	fmt.Println("\n\nゴルーチンテスト")
	ch := make(chan int8)
	go receive(ch)
	fmt.Println("待機開始")
	go sendDelay(ch)
	go sendDelay(ch)
	time.Sleep(1 * time.Second)
	<-ch // 受信待機
	fmt.Println("待機終了")

	ch1, ch2, ch3 := make(chan int8), make(chan int8), make(chan int8)

	go selectReceive(ch1, ch2, ch3)
	go selectReceive(ch1, ch2, ch3)
	go selectReceive(ch1, ch2, ch3)
	fmt.Println("select受信待機開始")

	go sendDelay(ch1)
	go sendDelay(ch3)
	go sendDelay(ch1)

	go sendDelay(ch2)
	selectReceive(ch1, ch2, ch3)

	time.Sleep(1000 * time.Millisecond)
	fmt.Println("待機終了")
}

func receive(in <-chan int8) {
	x := <-in
	fmt.Println("受信", x)
}

func sendDelay(out chan<- int8) {
	fmt.Println("送信 - sleep 0.5s")
	time.Sleep(500 * time.Millisecond)
	out <- 8
}

func selectReceive(in1 <-chan int8, in2 <-chan int8, in3 <-chan int8) {
	select {
	case v := <-in1:
		fmt.Println("Reveive1: ", v)
		return
	case v := <-in2:
		fmt.Println("Reveive2: ", v)
		return
	case v := <-in3:
		fmt.Println("Reveive3: ", v)
		return
		// defaultはチャンネル受信待ちせず通過する
		// default:
		// 	fmt.Println("Default")
		// 	return
	}
}
