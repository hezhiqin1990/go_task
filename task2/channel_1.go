package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for v := range ch {
			fmt.Println("recv:", v)
		}
		wg.Done()
	}()

	wg.Wait()

}
