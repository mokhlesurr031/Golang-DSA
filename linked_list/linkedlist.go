package main

import (
	"fmt"
	"strings"
)

type Node struct {
	value int
	next  *Node
}

type linkedList struct {
	head *Node
	len  int
}

func (l *linkedList) add(value int) {
	newNode := new(Node)
	newNode.value = value

	if l.head == nil {
		l.head = newNode
	} else {
		iterator := l.head
		for iterator.next != nil {
			iterator = iterator.next
		}
		iterator.next = newNode
	}
}
func (l *linkedList) remove(value int) {
	//return nil
}
func (l linkedList) get(value int) *Node {
	return nil
}

func (l linkedList) String() string {
	sb := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		//fmt.Println(iterator.value, iterator.next)
		sb.WriteString(fmt.Sprintf("%d ", iterator.value))
	}
	return sb.String()

}

func main() {
	fmt.Println("Hello")
	l := linkedList{}
	l.add(1)
	l.add(2)
	l.add(3)
	fmt.Println(l)
}
