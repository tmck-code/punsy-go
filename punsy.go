package main

import (
	"fmt"
	"github.com/tmck-code/punsy-go/src/struct"
	"github.com/tmck-code/punsy-go/src/dictionary"
)

func main() {
	t := Struct.NewTrie()
	// fileUrl := "http://svn.code.sf.net/p/cmusphinx/code/trunk/cmudict/cmudict-0.7b"
	ofpath := "cmudict-0.7b.utf8"
	// dictionary.DownloadFile(fileUrl, ofpath)
	dictionary.LoadCMU(ofpath, t)

	t.Insert("car", map[string]string{"key": "val"})
	t.Insert("cars", map[string]string{"keys": "vals"})
	t.Insert("carousel", map[string]string{"keys": "vals"})
	fmt.Println(t.Root.Repr())
	if n, ok := t.Get("car"); ok {
		fmt.Println(n.Repr())
		result := n.GetDescendents(1)
		for _, v := range(result) {
			fmt.Printf("%v\n", v)
		}
	}
}
