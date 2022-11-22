package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
 * @Project GoLearnProject
 * @File    basic.go
 * @Author  Augus Lee
 * @Date    2022/10/28 9:26
 * @GoVersion go1.19.2
 * @Version 1.0
 */

var (
	aa = 55
	bb = true
	ss = "hello"
)

func variableZeroValue() {
	var i int
	var s string
	fmt.Printf("i = %v, s = %q\n", i, s)
}

func variableInitiaValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i*math.Pi)+1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const ( //普通枚举类型
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	const ( //自增枚举类型
		cpp1 = iota
		_
		python1
		golang1
		javascript1
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(cpp1, javascript1, python1, golang1)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello, world!")
	variableZeroValue()     //演示变量的零值
	variableInitiaValue()   //演示变量的定义加赋值
	variableTypeDeduction() //演示变量同时定义并赋值
	variableShorter()       //演示变量定义简短方法
	fmt.Println(aa, bb, ss) //演示全局变量（用于一个包下）
	euler()                 //演示复数的使用
	triangle()              //演示强制类型转换
	consts()                //演示常量的定义
	enums()                 //演示枚举类型的定义
}
