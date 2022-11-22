package tree

import "fmt"

/**
 * @Project GoLearnProject
 * @File    traversal.go
 * @Author  Augus Lee
 * @Date    2022/11/4 9:08
 * @GoVersion go1.19.2
 * @Version 1.0
 */

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
