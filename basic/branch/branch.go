package main

import (
	"fmt"
	"os"
)

/**
 * @Project GoLearnProject
 * @File    branch.go
 * @Author  Augus Lee
 * @Date    2022/10/28 12:21
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func readfile() {
	const filename = "../file/abc.txt"
	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	if contents1, err1 := os.ReadFile(filename); err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Printf("%s\n", contents1)
	}
}

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score %d", score))
	}
	return g
}

func main() {
	readfile()

	fmt.Println(
		grade(0),
		grade(12),
		grade(64),
		grade(79),
		grade(88),
		grade(99),
		//grade(101),
		//grade(-1),
	)
}
