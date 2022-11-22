package main

import (
	"fmt"
	"time"
)

/**
 * @Project GoLearnProject
 * @File    goroutine.go
 * @Author  Augus Lee
 * @Date    2022/11/21 14:56
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from goroutine %d \n", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
