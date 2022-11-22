### 定义变量
- 使用 var 关键字
    - var a, b, c bool
    - var s1, s2 string = "hello", "world"
    - 可放在函数内，或直接放在包内
    - 使用 var() 集中定义变量
- 让编译器自动决定类型
    - var a, b, i, s1, s2 = true, false, 3, "hello", "world"
- 使用 := 定义变量
    - a, b, i, s1, s2 := true, false, 3, "hello", "world"
    - 只能在函数内使用

### 内建变量类型
- bool
- (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
- byte, rune
- float32, float64, complex64, complex128

### 强制类型转换
- 类型转换是强制的
- var a, b int = 3, 4
- var c int = int(math.Sqrt(float64(a * a + b * b)))

### 常量的定义
- comst filename = "abc.txt"
- const 数值可作为各种类型使用
- const a, b  = 3, 4

### 使用常量定义枚举类型
- 普通枚举类型
- 自增值枚举类型

### 变量定义要点回顾
- 变量类型写在变量名之后
- 编译器可推测变量类型
- 没有 char，只有 rune
- 原生支持复数

### if
- if 的条件里面可以赋值
- if 的条件里赋值的变量作用域就在这个if 语句里

### switch
- switch 会自动 break，除非使用 fallthrough
- switch 后可以没有表达式

### for
- for 的条件里不需要括号
- for 的条件里可以省略初始条件，结束条件，递增表达式

### 函数
- 返回值类型写在最后面
- 函数可以返回多个值
- 函数返回多个值时可以起名字（仅用于非常简单的函数，对于调用者而言没有区别）
- 函数可以作为函数的参数（匿名函数）
- 函数的可以接收可变参数列表

### 指针
- 指针不能运算
- Go语言只有值传递一种方式