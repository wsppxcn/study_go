package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func printInt(i int32)  {
	defer wg.Done()

	fmt.Println(i)
}

func main() {

	var count int32 = 0
	for count < 5 {
		// goroutine 没有执行完，主协程结束了
		wg.Add(1)
		go printInt(count)
		count += 1

	}
	// 也可以通过主协程block 2秒处理，但是不合理，你不知道要休眠多少s
	//time.Sleep(time.Duration(2) * time.Second)
	fmt.Println("is done")
	wg.Wait()
}