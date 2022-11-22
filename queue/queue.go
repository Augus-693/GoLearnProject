package queue

/**
 * @Project GoLearnProject
 * @File    queue.go
 * @Author  Augus Lee
 * @Date    2022/11/4 9:10
 * @GoVersion go1.19.2
 * @Version 1.0
 */

// A FIFO queue.
type Queue []int

// Pushes the element into the queue.
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pops element from head.
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
