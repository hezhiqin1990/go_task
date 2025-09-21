package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		wg  sync.WaitGroup
		mu  sync.Mutex
		num int = 0
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			for j := 0; j < 1000; j++ {
				num++
			}
			mu.Unlock()
			defer wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println(num)
}
