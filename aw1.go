package main

import (
	"fmt"
	"os"
)

type node struct {
	character rune
	parent    *node
	children  map[rune]*node
}

type nodePool struct {
	pool map[rune]*node
}

func main() {

	N := 1

	graphViz := false
	if os.Args[1] == "-g" {
		graphViz = true
		N = 2
	}

	var words [][]rune

	for _, word := range os.Args[N:] {
		words = append(words, []rune(word))
	}

	pool := &nodePool{
		pool: make(map[rune]*node),
	}

	root := pool.characterNode(words[0][0])

	for n := 0; n < len(words)-1; n++ {
		w0 := words[n]
		w1 := words[n+1]
		// fmt.Printf("%q <= %q\n", w0, w1)

		ln := len(w0)
		if len(w1) < ln {
			ln = len(w1)
		}

		for i := 0; i < ln; i++ {
			c0 := w0[i]
			c1 := w1[i]
			// fmt.Printf("%c <= %c\n", c0, c1)
			n0 := pool.characterNode(c0)
			n1 := pool.characterNode(c1)

			if c0 == c1 {
				// fmt.Printf("%c == %c\n", c0, c1)
				continue
			}

			if n1.isAfter(n0) {
				// fmt.Printf("already know %c < %c\n", c0, c1)
				break
			}

			n1.parent.removeChild(n1.character)
			n0.addChild(n1)
			if n1 == root {
				root = n0
			}

			break
		}
	}

	if graphViz {
		fmt.Printf("/* %d characters in language */\n", len(pool.pool))
		fmt.Println("digraph g {")
		fmt.Println("rankdir=\"LR\";")
		root.graphOut()
		fmt.Println("}")
		return
	}

	fmt.Printf("%d characters in language\n", len(pool.pool))
	fmt.Printf("first character %c\n", root.character)
	root.printChildren()
	/*
		for _, n := range pool.pool {
			fmt.Printf("%c children:\n", n.character)
			n.printChildren()
		}
	*/
}

func NewNode(character rune) *node {
	return &node{
		character: character,
		children:  make(map[rune]*node),
	}
}

func (n *node) addChild(c *node) {
	if n == nil {
		return
	}
	c.parent = n
	n.children[c.character] = c
}

func (n *node) removeChild(character rune) {
	if n == nil {
		return
	}
	if c := n.children[character]; c != nil {
		delete(n.children, character)
	}
}

func (n *node) isAfter(n0 *node) bool {
	if n == nil {
		return false
	}

	if n.character == n0.character {
		return true
	}
	return n.parent.isAfter(n0)
}

func (n *node) printChildren() {
	if n == nil {
		return
	}
	fmt.Printf("\t%c has %d children\n", n.character, len(n.children))
	for character, c := range n.children {
		fmt.Printf("%c < %c\n", n.character, character)
		c.printChildren()
	}
}

func (n *node) graphOut() {
	for _, c := range n.children {
		fmt.Printf("%c -> %c;\n", n.character, c.character)
	}
	for _, c := range n.children {
		c.graphOut()
	}
}

func (p *nodePool) characterNode(character rune) *node {
	if n, ok := p.pool[character]; ok {
		return n
	}
	n := NewNode(character)
	p.pool[character] = n
	return n
}
