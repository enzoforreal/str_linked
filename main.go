package main

import "fmt"

type Node struct {
	data string
	next *Node
}

var head *Node = new(Node)

func initial() {
	head.data = "R"
	head.next = nil

	var node1 *Node = &Node{data: "E", next: nil}
	head.next = node1

	var node2 *Node = &Node{data: "N", next: nil}
	node1.next = node2

	var node3 *Node = &Node{data: "A", next: nil}
	node2.next = node3

	var node4 *Node = &Node{data: "U", next: nil}
	node3.next = node4

	var tail *Node = &Node{data: "D", next: nil}
	node4.next = tail
}

func output(node *Node) {
	var str = node
	for {
		if str == nil {
			break
		}
		fmt.Printf("%s", str.data)
		str = str.next
	}

}

func main() {

	initial()

	output(head)

}
