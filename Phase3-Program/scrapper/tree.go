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
	fmt.Printf("%q\n", keys)
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
	recurPrint(t.root)
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

func print2DUtil(root *Node, space int) {
	// Base case
	if root == nil {
		return
	}
	// Increase distance between levels
	COUNT := 10
	space += COUNT

	for _, c := range root.child {
		print2DUtil(c, space)
	}
	fmt.Println()
	for i := 1; i < COUNT; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%s\n", root.title)
}
