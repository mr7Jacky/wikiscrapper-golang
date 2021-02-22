package main

import (
	"fmt"
	"strings"
)

type Node struct {
	level int
	link  string
	title string
	child []*Node
}

func NewNode(link string, title string) *Node {
	fmt.Printf("%q\n", strings.Split(link, "/"))
	return &Node{level: len(strings.Split(link, "/")), title: title, link: link, child: nil}
}

type Tree struct {
	root *Node
}

func NewTree(link string, title string) *Tree {
	return &Tree{root: NewNode(link, title)}
}

func (t *Tree) insert(link string, title string) {
	t.root.insert(link, title)
}

func (n *Node) insert(link string, title string) {
	keys := strings.Split(link, "/")
	// Same level
	fmt.Printf("%d, %d\n", n.level, len(keys)-1)
	if len(keys)-1 == n.level {
		// Child level
		newN := NewNode(link, title)
		n.child = append(n.child, newN)

	} else {
		// Lower Level
		for _, c := range n.child {
			if keys[c.level-1] == strings.Split(c.link, "/")[c.level-1] {
				c.insert(link, title)
				break
			}
		}
	}
}

func (t *Tree) PrintTree() {
	fmt.Println("Tree")
}
