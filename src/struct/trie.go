package Struct

import (
	"encoding/json"
	"github.com/tidwall/pretty"
	"log"
)

type Node struct {
	Final    bool
	Data     interface{}
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
		}
	}
	current.Final = true
	current.Data = data
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

func (current *Node) GetDescendents(max_depth int) ([]*Node) {
	results := make([]*Node, 0)
	if current.Final {
		results = append(results, current)
	}
	for range [5]int{} {
		for _, child := range(current.Children) {
			if max_depth > 0 {
				child.GetDescendents(max_depth-1)
			}
			if current.Final {
				break
			}
		}
	}
	return results
}
