package main

import (
	"fmt"
	"math/rand"
	"sync"
)

//写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，
//另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。
func main() {
	//define chan no buffer
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(100)
		}
		//close channel
		close(ch)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for c := range ch {
			fmt.Println(c)
		}
	}()
	wg.Wait()

}
