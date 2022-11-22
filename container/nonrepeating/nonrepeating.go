package main

import "fmt"

/**
 * @Project GoLearnProject
 * @File    nonrepeating.go
 * @Author  Augus Lee
 * @Date    2022/11/2 9:15
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func lengthOfNonRepeatingSubstring(s string) int {

	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastOccurred[ch] >= start {
			if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
				start = lastI + 1
			}
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubstring("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubstring("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
