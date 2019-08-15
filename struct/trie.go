package Struct

import (
	"fmt"
	"reflect"
)

type Node struct {
	Value rune
	Children map[rune]*Node
	Final bool
	Data interface{}
}

type Trie struct {
	Len int
	Root *Node
}

func (n Node) Insert(s string) {

}

func (t Trie) Insert(s string) {
	current := t.Root
	for i, char := range s {
		fmt.Printf("%v, %+v, %v\n", i, char, reflect.TypeOf(char))
		if _, ok := current.Children[char]; !ok {
			current.Children[char] = &Node{
				Value: char,
				Children: make(map[rune]*Node, 0),
				Final: false,
				Data: nil,
			}
		}
		current = current.Children[char]
	}
}

