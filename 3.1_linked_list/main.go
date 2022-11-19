package main

import "fmt"

type node struct {
	value string
	next  *node
}
type linkedList struct {
	head   *node
	lenght int
}

// ADD NODE TO THE LIST
func (l *linkedList) appendNode(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.lenght++
}

// PRINT LIST
func (l linkedList) printListNodes() {
	toPrint := l.head
	for l.lenght != 0 {
		fmt.Print(" ", toPrint.value)
		toPrint = toPrint.next
		l.lenght--
	}

}

// DELETE NODE BY VALUE
func (l *linkedList) deleteNodebyValue(v string) {
	if l.lenght == 0 {
		return
	}
	if l.head.value == v {
		l.head = l.head.next
		l.lenght--
	}
	previousToDelete := l.head
	for previousToDelete.next.value != v {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.lenght--
}

func main() {
	list := linkedList{}
	node1 := &node{value: "first"}
	node2 := &node{value: "second"}
	node3 := &node{value: "third"}
	node4 := &node{value: "fourth"}
	node5 := &node{value: "fifth"}

	// ADD NODE
	list.appendNode(node1)
	list.appendNode(node2)
	list.appendNode(node3)
	list.appendNode(node4)
	list.appendNode(node5)
	fmt.Println(list)
	list.printListNodes()
	fmt.Println()

	// DELETE NODE BY VALUE
	list.deleteNodebyValue("fifth")
	list.printListNodes()
}
