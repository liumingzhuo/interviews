/*
实现阻塞读且并发安全的map
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type sp interface {
	Out(key string, val interface{})                  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

//Map include map and rwmutex
type Map struct {
	im map[string]*item
	rw *sync.RWMutex
}

//item include chan value flag
type item struct {
	ch    chan struct{}
	value interface{}
	flag  bool
}

func (m *Map) Out(key string, val interface{}) {
	m.rw.Lock()
	defer m.rw.Unlock()
	it, ok := m.im[key]
	// key not found , create entry with value and flag is true
	if !ok {
		it := &item{
			value: val,
			flag:  true,
		}
		m.im[key] = it
		return
	}
	it.value = val
	//flag true , item's value exist
	if it.flag {
		return
	}
	if it.ch != nil {
		close(it.ch)
		it.ch = nil
	}
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	// add read mutex
	m.rw.RLock()
	if it, ok := m.im[key]; ok && it.flag {
		// value exist
		m.rw.RUnlock()
		return it.value
	} else if !ok {
		m.rw.RUnlock()
		//add write Lock
		m.rw.Lock()
		entry := &item{ch: make(chan struct{}),
			flag: false}
		m.im[key] = entry
		//must unlock
		m.rw.Unlock()
		fmt.Println("goroutine block")
		//select
		select {
		case <-entry.ch:
			return entry.value
		case <-time.After(timeout):
			fmt.Println("goroutine timeout key is ", key)
			return nil
		}
	} else {
		//value exist but flag false
		m.rw.RUnlock()
		fmt.Println("goroutine block")
		select {
		case <-it.ch:
			return it.value
		case <-time.After(timeout):
			fmt.Println("goroutine timeout key is ", key)
			return nil
		}
	}
}
func main() {
	mm := Map{
		im: make(map[string]*item),
		rw: &sync.RWMutex{},
	}

	for i := 0; i < 10; i++ {
		go func() {
			val := mm.Rd("k", time.Second*6)
			fmt.Println("val is :", val)
		}()
	}
	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		go func(i int) {
			mm.Out("k", i)
		}(i)
	}
	time.Sleep(time.Second * 30)

}
