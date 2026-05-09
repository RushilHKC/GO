package main

import (
	f "fmt"
)

type Node[T int] struct {
	val  T
	next *Node[T]
}

func NewNode[T int](val T) *Node[T] {
	return &Node[T]{val: val}
}

func (node *Node[T]) Append(val T) {
	n1 := NewNode(val)
	if node == nil {
		node = n1
	}
	current := node
	for current.next != nil {
		current = current.next
	}

	current.next = n1

}

func (node *Node[T]) Prepend(val T) *Node[T] {
	n1 := NewNode(val)
	n1.next = node
	return n1
}

func merge[T int](l1, l2 *Node[T]) *Node[T] {
	left := l1
	right := l2

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.val <= right.val {
		left.next = merge(left.next, right)
		return left
	} else {
		right.next = merge(left, right.next)
		return right
	}

}

func mergeSort[T int](l1 *Node[T]) *Node[T] {
	if l1.next == nil {
		return l1
	}

	slow := l1
	fast := l1
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	mid := slow
	l2 := slow.next
	mid.next = nil
	l1 = mergeSort(l1)
	l2 = mergeSort(l2)
	return merge(l1, l2)
}

// func main() {
// 	l1 := NewNode(5)
// 	l1.Append(8)
// 	l1 = l1.Prepend(12)
// 	l1 = l1.Prepend(76)
// 	l1.Append(1)

// 	current := l1
// 	for current != nil {

// 		f.Print(current.val)
// 		f.Print(" ")
// 		current = current.next
// 	}
// 	f.Println()
// 	l1 = mergeSort(l1)
// 	current = l1
// 	for current != nil {
// 		f.Print(current.val)
// 		f.Print(" ")
// 		current = current.next
// 	}
// }
