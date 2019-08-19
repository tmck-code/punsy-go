package Struct

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/pretty"
	"log"
	"reflect"
)

type Node struct {
	Value    rune
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

func (n Node) Repr() {
	if b, err := json.Marshal(n); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(pretty.Color(pretty.Pretty(b), nil)))
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

func (t Trie) GetDescendents(s string) (*Node, bool) {
	if current, ok := t.Get(s); ok {
		for {
			for _, child := range current.Children {
				fmt.Printf("----- %+v\n", child)
			}
			fmt.Printf("+++++++ %+v\n", current)
			if current.Final {
				break
			}
		}
	}
	return nil, false
}
