package main

import (
	"bufio"
	"fmt"
	"github.com/GoLearnProject/functional/fib"
	"io"
	"strings"
)

/**
 * @Project GoLearnProject
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2022/11/9 11:04
 * @GoVersion go1.19.2
 * @Version 1.0
 */

type intGen func() int

func (i intGen) Read(p []byte) (n int, err error) {
	next := i()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fib.Fibonacci()
	printFileContents(f)
}
