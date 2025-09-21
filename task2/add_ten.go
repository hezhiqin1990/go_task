package main

import "fmt"

func main() {
	var x = 5
	add(&x)
	fmt.Println(x)
	var y = 10
	add(&y)
	fmt.Println(y)
}

func add(x *int) {
	*x += 10
}
