package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// node instances represent single characters discovered so far.
// They also hold pointers to a node representing a character that's
// next lexically less than this character, and a map of structs node
// representing characters that are lexically greater than this character.
type node struct {
	character rune
	parent    *node
	children  map[rune]*node
}

// nodePool holds references to all the characters seen so far
type nodePool map[rune]*node

func main() {
	graphViz := flag.Bool("g", false, "GraphViz dot format on stdout")
	flag.Parse()

	fin := os.Stdin
	if flag.NArg() >= 1 {
		var err error
		if fin, err = os.Open(flag.Arg(0)); err != nil {
			log.Fatal(err)
		}
		defer fin.Close()
	}

	var words [][]rune

	scanner := bufio.NewScanner(fin)

	lineCounter := 0

	for scanner.Scan() {
		lineCounter++
		line := scanner.Text()
		word := strings.ToLower(strings.TrimSpace(line))
		words = append(words, []rune(word))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("problem line %d: %v", lineCounter, err)
	}

	pool := nodePool(make(map[rune]*node))

	root := pool.characterNode(words[0][0])

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
			// add characters to pool here, so as not to leave any out
			n0 := pool.characterNode(c0)
			n1 := pool.characterNode(c1)

			if c0 == c1 {
				continue
			}

			if n1.isAfter(n0) {
				break
			}

			n1.parent.removeChild(n1.character)
			n0.addChild(n1)
			if n1 == root {
				// keep track of the absolutely lexically least character
				root = n0
				root.parent = nil
			}

			break
		}
	}

	if *graphViz {
		fmt.Printf("/* %d words in input\n", len(words))
		fmt.Printf(" * %d characters in language\n", len(pool))
		fmt.Printf(" * first character %c\n*/\n", root.character)
		fmt.Println("digraph g {")
		fmt.Println("rankdir=\"LR\";")
		root.graphOut()
		fmt.Println("}")
		return
	}

	fmt.Printf("%d words in input\n", len(words))
	fmt.Printf("%d characters in language\n", len(pool))
	fmt.Printf("first character %c\n", root.character)
	root.printChildren()
	/*
		for _, n := range pool.pool {
			fmt.Printf("%c children:\n", n.character)
			n.printChildren()
		}
	*/
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
		// c.parent left as is
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

// characterNode is the official way to get a pointer to a struct node.
func (p nodePool) characterNode(character rune) *node {
	if n, ok := p[character]; ok {
		return n
	}
	n := &node{
		character: character,
		children:  make(map[rune]*node),
	}
	p[character] = n
	return n
}
