package main

import "fmt"

type Node struct {
	data string
	next *Node
}

func reverseList(n Node) Node {
	m := &Node{
		data: n.data,
		next: nil,
	}
	n = *n.next
	for n.next != nil {
		tmp := &Node{
			data: n.data,
			next: m,
		}
		m = tmp
		n = *n.next
	}
	tmp := &Node{
		next: m,
		data: n.data,
	}
	return *tmp
}

func buildList(ss []string) Node {
	headNode := Node{
		data: ss[0],
		next: &Node{},
	}
	n := headNode.next
	for i := 1; i < len(ss); i++ {
		n.data = ss[i]
		n.next = &Node{}
		n = n.next
	}
	n.data = "nil"
	n.next = nil

	return headNode
}

func main() {
	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	in := buildList(letters)
	out := reverseList(in)

	// print original
	fmt.Printf("original = %s -> ", in.data)
	in = *in.next
	for in.next != nil {
		fmt.Printf("%s -> ", in.data)
		in = *in.next
	}
	fmt.Printf("%s\n", in.data)

	// print reversed
	fmt.Printf("reversed = %s", out.data)
	out = *out.next
	for out.next != nil {
		fmt.Printf(" -> %+v", out.data)
		out = *out.next
	}
	fmt.Printf(" -> %s\n", out.data)
}
