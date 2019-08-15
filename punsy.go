package main

import (
	"github.com/tmck-code/punsy-go/struct"
	"fmt"
)

func main() {
	t := Struct.NewTrie()
	msg := "car"
	fmt.Println(msg)
	t.Insert(msg)
	current := t.Root
	for _, char := range msg {
		current.Repr()
		fmt.Println(char)
		current = current.Children[char]
	}
}
