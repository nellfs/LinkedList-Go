//This was done based on Junmin Lee's example of linked lists.

package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	lenght int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.lenght++
}

func (l linkedList) printListData(description string) {
	toPrint := l.head
	for l.lenght != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.lenght--
	}
  fmt.Printf("-> %s\n", description)
}

func (l *linkedList) deleteWithValue(value int) {
	if l.lenght == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.lenght--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
    if previousToDelete.next.next == nil {
      return
    }
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.lenght--
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 48}
	mylist.prepend(node1)
	node2 := &node{data: 18}
	mylist.prepend(node2)
	node3 := &node{data: 16}
	mylist.prepend(node3)
	node4 := &node{data: 20}
	mylist.prepend(node4)

	mylist.printListData("4 nodes")

	node5 := &node{data:33}
	mylist.prepend(node5)

  mylist.printListData("33 was added to the beginning")

	mylist.deleteWithValue(16)

	mylist.printListData("16 was deleted, 20 connected with 18")
}
