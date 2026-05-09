package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func NewList[T any](val T) *List[T] {
	return &List[T]{val: val}
}

func (l *List[T]) Append(val T) {
	newNode := &List[T]{val: val}

	if l == nil {
		return
	}
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *List[T]) Prepend(val T) *List[T] {
	newNode := &List[T]{val: val}
	fmt.Println(newNode)
	newNode.next = l
	l = newNode
	fmt.Println(l)
	return newNode
}

type Foo struct {
	val int
	arr []int
}

// func (f *Foo) set(newValue int){
// 	f = newValue
// }
// func (f *Foo) replace(newFoo *Foo){
// 	f = newFoo
// }

func (f Foo) AppendToArr(val int) {
	f.val = 110000
	f.arr = append(f.arr, val)
}

// func main() {
// 	l1 := NewList(3)
// 	l1.Append(4)

// 	l2 := l1.Prepend(2)
// 	fmt.Println(l1)
// 	current := l2
// 	for current != nil {
// 		fmt.Println(current.val)
// 		current = current.next
// 	}
// 	// f := &Foo{}
// 	// f.set(3)
// 	// f.replace(&Foo{val:6, arr:{}})
// 	// f.AppendToArr(3)
// 	// fmt.Println(f)
// }
