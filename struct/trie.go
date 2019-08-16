package Struct

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Node struct {
	Value    rune
	Children map[rune]*Node
	Final    bool
	Data     interface{}
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
		Children: make(map[rune]*Node, 0),
		Final:    false,
	}
}

func (n Node) Repr() {
	if b, err := json.Marshal(n); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(b))
	}
}

func (t Trie) Insert(s string, data interface{}) {
	current := t.Root
	for i, char := range s {
		fmt.Printf("%v, %+v, %v\n", i, char, reflect.TypeOf(char))
		if _, ok := current.Children[char]; ok {
			current = current.Children[char]
		} else {
			current.Children[char] = NewNode()
			current = current.Children[char]
			current.Value = char
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
