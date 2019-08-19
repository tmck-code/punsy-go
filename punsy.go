package main

import (
	"fmt"
	"github.com/tmck-code/punsy-go/src/struct"
)

func main() {
	t := Struct.NewTrie()
	msg := "car"
	t.Insert(msg, map[string]string{"key": "val"})
	fmt.Println(t.Root.Repr())
	if n, ok := t.Get("car"); ok {
		fmt.Println(n.Repr())
	}
	if n, ok := t.GetDescendents(msg); ok {
		fmt.Println(n.Repr())
	}
}
