package main

import "fmt"

func add() func(x int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	x, y := add(), add()
	for i := 0; i < 10; i++ {
		fmt.Println(x(i), y(-2*i))
	}
}
