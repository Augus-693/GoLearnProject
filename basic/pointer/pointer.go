package main

import "fmt"

/**
 * @Project GoLearnProject
 * @File    pointer.go
 * @Author  Augus Lee
 * @Date    2022/10/28 14:11
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap1(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	a, b = swap1(a, b)
	fmt.Println(a, b)

}
