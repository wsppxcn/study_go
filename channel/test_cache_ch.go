package main

import (
	"fmt"
	"time"
)

func Recv(ch chan int) {
	res := <-ch
	fmt.Sprintf("chin :%d", res)
}
func main() {
	//ch := make(chan int,10)
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		if data, ok := <-ch; ok {
			//fmt.Sprintf("get chan:%d", data)
			fmt.Println("get chan:",data)
		} else {
			break
		}
	}

	//go Recv(ch)
	//fmt.Println("发送成功")

	time.Sleep(time.Second)
}
