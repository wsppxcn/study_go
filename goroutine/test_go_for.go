package main

import (
	"fmt"
	"time"
)

func hello (){
	fmt.Sprintf("hello")
	time.Sleep(time.Duration(2) * time.Second)
}

func main() {

	go func() {
		for  {
			hello ()
		}
	}()

	fmt.Println("is done")
}
