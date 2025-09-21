package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter uint32

func increment() {
	atomic.AddUint32(&counter, 1)
}

func getCounter() uint32 {
	return atomic.LoadUint32(&counter)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				increment()
			}
			defer wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println(getCounter())
}
