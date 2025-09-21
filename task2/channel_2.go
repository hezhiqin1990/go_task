package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup //1、定义全局的 WaitGrou

var ch = make(chan int, 10)

const jobCount = 100

// 生产者消费者模型
func main() {
	w.Add(2)
	go product()
	go costm()
	w.Wait() // 等两边都收工

}

func product() {
	for i := 1; i <= jobCount; i++ {
		ch <- i
	}
	close(ch)
	defer w.Done() //goroutine结束就登记-1
}

func costm() {
	for v := range ch {
		fmt.Println("recv:", v)
	}
	defer w.Done() //goroutine结束就登记-1
}
