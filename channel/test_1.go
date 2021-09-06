package main

import (
	"fmt"
	"time"
)



func recv(i chan int){

	j := <- i
	//fmt.Sprintf("chan recv j : %d",j)
	fmt.Println("接收成功", j)
}
func main() {

	cha := make(chan int)
	go recv(cha)
	cha <- 10
	fmt.Println("send ok")

	time.Sleep(time.Second)



}


