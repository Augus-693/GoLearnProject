package fib

/**
 * @Project GoLearnProject
 * @File    fib.go
 * @Author  Augus Lee
 * @Date    2022/11/9 11:04
 * @GoVersion go1.19.2
 * @Version 1.0
 */

// Fibonacci 1, 1, 2, 3, 5, 8, 13, ...
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
