package main

import "fmt"

type twice_func_type func(input int) int

type myNum struct {
	data  int
	twice twice_func_type
}

func doTwice(input int) int {
	return input * 2
}

func (input myNum) idealTwice() int {
	return input.data * 2
}
func main() {
	x := myNum{twice: doTwice, data: 100}
	fmt.Println(x.twice(x.data))
	fmt.Println(x.idealTwice())
}

