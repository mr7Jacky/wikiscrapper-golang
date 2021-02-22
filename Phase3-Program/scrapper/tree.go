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
	if len(keys)-1 == n.level {
		// Child level
		newN := NewNode(link, title)
		n.child = append(n.child, newN)
	} else {
		// Lower Level
		for _, c := range n.child {
			cKeys := strings.Split(c.link, "/")
			if keys[c.level-2] == cKeys[c.level-2] {
				c.insert(link, title)
				break
			}
		}
	}
}

func (t *Tree) PrintTree() {
	LevelOrderTraversal(t.root)
}

func recurPrint(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("layer: %d Name: %s\n", root.level, root.title)
	for _, c := range root.child {
		recurPrint(c)
	}
}

func LevelOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	// Standard level order traversal code
	// using queue
	var q []*Node // Create a queue
	q = append(q, root)
	for len(q) != 0 {
		n := len(q)
		for n > 0 {
			p := q[0]
			q = q[1:]
			fmt.Printf("%s ", p.title)
			for i := 0; i < len(p.child); i++ {
				q = append(q, p.child[i])
			}
			n--
		}
		fmt.Println()
	}

}
