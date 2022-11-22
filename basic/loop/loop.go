package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * @Project GoLearnProject
 * @File    loop.go
 * @Author  Augus Lee
 * @Date    2022/10/28 12:39
 * @GoVersion go1.19.2
 * @Version 1.0
 */

/**
 * @functionName convertToBin
 * @author Augus Lee
 * @description 整数转为二进制表达式
 * @date 2022-10-28 12:42:26
 * @param n int
 * @return string
 **/
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	i := 1
	for {
		fmt.Println("Love You", i, "Times")
		i++
	}

}

func main() {
	fmt.Println(
		convertToBin(5),  //101
		convertToBin(13), //1101
		convertToBin(8941653),
	)

	printFile("../file/abc.txt")

	forever()
}
