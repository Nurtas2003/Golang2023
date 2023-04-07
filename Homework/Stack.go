package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}
type Stack struct {
	head *Node
}

func (st *Stack) push(val int) {
	newNode := &Node{val: val, next: nil}
	if st.head == nil {
		st.head = newNode
	} else {
		newNode.next = st.head
		st.head = newNode

	}
}
func (st *Stack) pop() {
	st.head = st.head.next
}
func (st *Stack) peek() int {
	return st.head.val
}
func (st *Stack) clear() {
	st.head = nil
}
func (st *Stack) contains(val int) bool {
	cur := st.head
	for cur != nil {
		if cur.val == val {
			return true
		}
		cur = cur.next
	}
	return false
}
func (st *Stack) increment(val int) {
	cur := st.head
	for cur != nil {
		cur.val += val
		cur = cur.next
	}
}

func (st *Stack) print() {
	cur := st.head
	for cur != nil {
		fmt.Printf("%d ", cur.val)
		cur = cur.next
	}
	fmt.Println()
}

func (st *Stack) printReverse() {
	cur := st.head
	arr := make([]int, 0)
	for cur != nil {
		arr = append(arr, cur.val)
		cur = cur.next
	}
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("%d", arr[i])
	}
	fmt.Println()
}

func main() {
	var st Stack
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		st.push(x)
	}
	st.print()
	st.increment(1)
	st.print()
	st.printReverse()
}
