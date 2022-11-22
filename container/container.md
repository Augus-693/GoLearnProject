### 数组
- var array1 [5]int
- array2 := [3]int{1, 2, 3}
- array3 := [...]int{1, 2, 3}
- var grid [4][5]bool
- 数量写在类型前

### 数组的遍历
- 可通过 _ 省略变量
- 不仅 range，任何地方都可通过 _ 省略变量
- 如果只要 i，可以写成 for i := range numbers{}

### 数组是值类型
- [10]int 和 [20]int 是不同类型
- 调用 func f(arr [10]int) 会**拷贝**数组
- 在go语言中一般不直接使用数组

### Slice(切片)
- arr := [...]int[0, 2, 3, 4, 5, 6, 7]
- s := arr[2:6]
- s 的值是 [2, 3, 4, 5]
- Slice 本身没有数据，是对底层 array 的一个 view
- slice 可以向后扩展，不可以向前扩展
- s[i] 不可以超越 len(s)，向后扩展不可以超过底层数组 cap(s)

### 向 slice 添加元素
- 添加元素时如果超过 cap，系统会重新分配更大的底层数组
- 由于值传递的关系，必须接收 append 的返回值
- s = append(s, val)
- 

### map 的操作
- 创建：make(map[string]string)
- 获取元素：m[key]
- key不存在时，获得value类型的初始值
- 用value, ok := m["key"] 来判断是否存在key
- 用 delete 删除一个key

### map 的遍历
- 使用 range 遍历 map，或者遍历 key, value 对
- 不保证遍历顺序，如需顺序，需手动对 key 进行排序
- 使用 len 获得元素个数

### map 的 key
- map 使用哈希表，必须可以比较相等
- 除了 slice, map, function 的内建类型都可以作为 key
- Struct 类型不包含上述字段，也可作为 key

### 寻找最长不含有重复字符的子串
对于每一个字母X
- lastOccurred[x]不存在，或者 < start  ——> 无需操作
- lastOccurred[x] >= start  ——> 更新 start
- 更新 lastOccurred[x]，更新 maxLength

### rune相当于go的char
- 使用 range 遍历 pos，rune 对
- 使用 utf8.RuneCountInString 获得字符数量
- 使用 len 获得字节长度
- 使用 []byte 获得字节 
