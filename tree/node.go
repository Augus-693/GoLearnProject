package tree

import "fmt"

/**
 * @Project GoLearnProject
 * @File    node.go
 * @Author  Augus Lee
 * @Date    2022/11/4 9:07
 * @GoVersion go1.19.2
 * @Version 1.0
 */

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}
