package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //1、定义全局的 WaitGrou

func main() {
	wg.Add(1) //2、启动一个goroutine就登记+1
	go Jishu()
	wg.Add(1) //2、启动一个goroutine就登记+1
	go Oushu()
	defer wg.Wait()
}

func Jishu() {
	//打印1-10的奇数
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i, "是基数")
		}
	}
	defer wg.Done() //goroutine结束就登记-1
}

func Oushu() {
	//打印2-10的偶数
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "是偶数")
		}
	}
	defer wg.Done() //goroutine结束就登记-1
}
