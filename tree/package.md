### 面向对象
- go语言仅支持封装，不支持继承和多态
- go语言没有class，只有struct

### 结构的创建
~~~go
root := treeNode{value: 3}
root.left = &treeNode{}
root.right = &treeNode{5, nil, nil}
root.right.left = new(treeNode)
~~~
- 不论是地址还是结构本身，一律使用 . 来访问成员

~~~go
func createNode(value int) *treeNode {
	return &treeNode{Value: value}
}

root.Left.Right = createNode(2)
~~~
- 使用自定义工厂函数
- 注意返回了局部变量的地址

### 为结构体定义方法
~~~go
func (node TerrNode) print() {
    fmt.Printf(node.Value)
}
~~~
- 显式定义和命名方法接收者

### 使用指针作为方法接收者
~~~go
func (node *TerrNode) setValue(value int) {
    node.Value = value
}
~~~
- 只有使用指针才可以改变结构内容
- nil 指针也可以调用方法

### 值接收者 vs 指针接收者
- 要改变内容必须使用指针接收者
- 结构体较大也考虑使用指针接收者
- 一致性：如有指针接收者，最好都是指针接收者
- 值接收者是go语言特有
- 值/指针接收者均可接收值/指针

### 封装
- 名字一般使用CamelCase
- 首字母大写：public
- 首字母小写：privite 

#### 包
- 每个目录一个包
- main 包包含可执行入口
- 为结构体定义的方法必须放在同一个包里
- 可以是不同文件