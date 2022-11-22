package main

import "fmt"

/**
 * @Project GoLearnProject
 * @File    sliceops.go
 * @Author  Augus Lee
 * @Date    2022/10/28 16:40
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	var s []int //Zero value for slice operations is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{1, 2, 3, 4}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice............")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice............")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping elements from slice front............")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping elements from slice back............")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)

}
