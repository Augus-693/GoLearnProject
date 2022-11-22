package queue

import "fmt"

/**
 * @Project GoLearnProject
 * @File    queue_test.go
 * @Author  Augus Lee
 * @Date    2022/11/4 9:11
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	// Output:
	// 1
	// 2
	// false
	// 3
	// true
}
