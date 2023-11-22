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

	var words [][]rune

	for _, word := range os.Args[1:] {
		words = append(words, []rune(word))
	}

	pool := &nodePool{
		pool: make(map[rune]*node),
	}

	for n := 0; n < len(words)-1; n++ {
		w0 := words[n]
		w1 := words[n+1]

		ln := len(w0)
		if len(w1) < ln {
			ln = len(w1)
		}

		for i := 0; i < ln; i++ {
			c0 := w0[i]
			c1 := w1[i]
			if c0 == c1 {
				continue
			}
			n0 := pool.characterNode(c0)
			n1 := pool.characterNode(c1)

			n1.parent.removeChild(n1.character)
			n0.addChild(n1)
		}
	}

	for _, n := range pool.pool {
		n.printChildren()
	}
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

func (n *node) printChildren() {
	if n == nil {
		return
	}
	for character, c := range n.children {
		fmt.Printf("%c < %c\n", n.character, character)
		c.printChildren()
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
