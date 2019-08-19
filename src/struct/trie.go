package Struct

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/pretty"
	"log"
)

type Node struct {
	Final    bool
	Data     interface{}
	Children map[rune]*Node
}

type Trie struct {
	Len  int
	Root *Node
}

func NewTrie() *Trie {
	return &Trie{
		Len:  0,
		Root: NewNode(),
	}
}

func NewNode() *Node {
	return &Node{
		Final:    false,
		Children: make(map[rune]*Node, 0),
	}
}

func (n Node) Repr() string {
	b, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}
	return string(pretty.Color(pretty.Pretty(b), nil))
}

func (t Trie) Insert(s string, data interface{}) {
	current := t.Root
	for _, char := range s {
		if _, ok := current.Children[char]; ok {
			current = current.Children[char]
		} else {
			current.Children[char] = NewNode()
			current = current.Children[char]
		}
	}
	current.Final = true
	current.Data = data
}

func (t Trie) Get(s string) (*Node, bool) {
	current := t.Root
	for _, char := range s {
		if _, ok := current.Children[char]; ok {
			current = current.Children[char]
		} else {
			return nil, false
		}
	}
	return current, true
}

func (t Trie) GetDescendents(s string) (*Node, bool) {
	if current, ok := t.Get(s); ok {
		for {
			for _, child := range current.Children {
				fmt.Printf("----- %+v\n", child)
			}
			if current.Final {
				break
			}
		}
	}
	return nil, false
}
