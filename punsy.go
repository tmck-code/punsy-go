package main

import (
	"fmt"
	"github.com/tmck-code/punsy-go/struct"
)

func main() {
	t := Struct.NewTrie()
	msg := "car"
	fmt.Println(msg)
	t.Insert(msg, map[string]string{"key": "val"})
	current := t.Root
	for _, char := range msg {
		current.Repr()
		fmt.Println(char)
		current = current.Children[char]
	}
	if n, ok := t.Get("car"); ok {
		n.Repr()
	}
}
