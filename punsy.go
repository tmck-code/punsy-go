package main

import (
	"github.com/tmck-code/punsy-go/struct"
	"fmt"
)

func main() {
	t := new(Struct.Trie)
	msg := "car"
	fmt.Println(msg)
	t.Insert(msg)
}
