package main

import (
	"fmt"
)

/**
 * @Project GoLearnProject
 * @File    slices.go
 * @Author  Augus Lee
 * @Date    2022/10/28 15:34
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func updateSlices(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 = ", s1)
	s2 := arr[:]
	fmt.Println("s2 = ", s2)

	fmt.Println("After updateSlices(s1)............")
	updateSlices(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updataSlices(s2)............")
	updateSlices(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslices............")
	s2 = s2[:5]
	s2 = s2[2:]
	fmt.Println("s2 = ", s2)

	fmt.Println("Extending slice............")
	//slice可以向后扩展，不可以向前扩展
	//s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6] //s1 =  [2 3 4 5]
	s2 = s1[3:5]  //s2 =  [5 6]
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n", s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Printf("s3 = %v, len(s3) = %d, cap(s3) = %d\n", s3, len(s3), cap(s3))
	//添加元素时如果超越cap，系统会重新分配更大的底层数组
	fmt.Printf("s4 = %v, len(s4) = %d, cap(s4) = %d\n", s4, len(s4), cap(s4))
	fmt.Printf("s5 = %v, len(s5) = %d, cap(s5) = %d\n", s5, len(s5), cap(s5))
	fmt.Printf("arr = %v, len(arr) = %d, cap(arr) = %d\n", arr, len(arr), cap(arr))
}
