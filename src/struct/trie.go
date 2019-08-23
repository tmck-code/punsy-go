package Struct

import (
	"encoding/json"
	"github.com/tidwall/pretty"
	"log"
)

type Node struct {
	Final    bool
	Data     []interface{}
	Children map[string]*Node
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
		Children: make(map[string]*Node, 0),
	}
}

func (n Node) Repr() string {
	b, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}
	return string(pretty.Color(pretty.Pretty(b), nil))
}

func (t *Trie) Insert(s []string, data interface{}) {
	current := t.Root
	for _, char := range s {
		if _, ok := current.Children[char]; ok {
			current = current.Children[char]
		} else {
			current.Children[char] = NewNode()
			current = current.Children[char]
			current.Data = make([]interface{}, 0)
		}
	}
	current.Final = true
	current.Data = append(current.Data, data)
	t.Len += 1
}

func (t Trie) Get(s []string) (*Node, bool) {
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

func GetDescendentsData(current Node, max_depth int, results []interface{}) []interface{} {
	if current.Final {
		results = append(results, current.Data...)
	}
	if max_depth > 0 && len(current.Children) > 0 {
		for _, child := range current.Children {
			results = append(
				results,
				GetDescendentsData(*child, max_depth-1, make([]interface{}, 0))...,
			)
		}
	}
	return results
}
