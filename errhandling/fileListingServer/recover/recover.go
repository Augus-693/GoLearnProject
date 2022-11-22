package main

import "fmt"

/**
 * @Project GoLearnProject
 * @File    recover.go
 * @Author  Augus Lee
 * @Date    2022/11/9 14:48
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func tryRecover() {
	defer func() {
		r := recover()
		if r == nil {
			fmt.Println("Nothing to recover. " +
				"Please try uncomment errors " +
				"below.")
			return
		}
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(fmt.Sprintf(
				"I don't know what to do: %v", r))
		}
	}()

	// Uncomment each block to see different panic
	// scenarios.
	// Normal error
	//panic(errors.New("this is an error"))

	// Division by zero
	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	// Causes re-panic
	//panic(123)
}

func main() {
	tryRecover()
}
