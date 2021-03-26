//交替打印数字和字⺟
//问题描述
//使⽤两个  goroutine 交替打印序列，⼀个  goroutine 打印数字， 另外⼀个  goroutine 打印字⺟， 最终效果如下：
//12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
package altprint

import (
	"fmt"
	"strings"
	"sync"
)

func Altprint() {
	letter, num := make(chan struct{}), make(chan struct{})
	wait := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <-num:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- struct{}{} //通知打印字母的goroutine
				break
			default:
				break
			}

		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				//i>=26的时候打印完毕
				if i >= strings.Count(str, "")-1 {
					wait.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				if i >= strings.Count(str, "") {
					i = 0
				}
				fmt.Print(str[i : i+1])
				i++
				num <- struct{}{} //通知打印数字的goroutine
				break
			default:
				break
			}
		}
	}(&wait)
	num <- struct{}{}
	wait.Wait()

}
