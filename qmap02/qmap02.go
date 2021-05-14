/*
高并发下的锁与map的读写
场景：在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。
每个IP三分钟之内只能访问一次
*/
package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
	lock     sync.Mutex
}

func NewBan(ctx context.Context) *Ban {
	o := &Ban{visitIPs: make(map[string]time.Time)}
	//开启goroutine 对ip进行时间控制
	go func() {
		timer := time.NewTimer(time.Minute * 3)
		for {
			select {
			case <-timer.C: //timer向c字段发送当前时间
				o.lock.Lock()
				for k, v := range o.visitIPs {
					if time.Since(v) >= time.Minute*3 {
						delete(o.visitIPs, k)
					}
				}
				o.lock.Unlock()
				timer.Reset(time.Minute * 3)
			case <-ctx.Done():
				return
			}
		}
	}()
	return o
}
func (o *Ban) visit(key string) bool {
	o.lock.Lock()
	defer o.lock.Unlock()
	if _, ok := o.visitIPs[key]; ok {
		//如果key存在则跳过
		return true
	}
	o.visitIPs[key] = time.Now()
	return false
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	b := NewBan(ctx)
	success := int64(0)
	wait := &sync.WaitGroup{}
	wait.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wait.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !b.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}
	}
	wait.Wait()
	fmt.Println("ip success ", success)
}
