package main

import (
	"fmt"
	"github.com/GoLearnProject/queue"
)

/**
 * @Project GoLearnProject
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2022/11/4 9:10
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
