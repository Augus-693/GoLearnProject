package main

import "fmt"

/**
 * @Project GoLearnProject
 * @File    adder.go
 * @Author  Augus Lee
 * @Date    2022/11/9 10:45
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0+...+%d= %d\n", i, s)
	}
}
