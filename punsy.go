package main

import (
	"fmt"
	"github.com/tmck-code/punsy-go/src/dictionary"
)

func main() {
	c := dictionary.NewCMU()
	// fileUrl := "http://svn.code.sf.net/p/cmusphinx/code/trunk/cmudict/cmudict-0.7b"
	ofpath := "cmudict-0.7b.utf8"
	// dictionary.DownloadFile(fileUrl, ofpath)
	dictionary.LoadCMU(ofpath, c)

	// msg := "hello world"

	if n, ok := c.Rhymes.Get([]string{"AE1", "K", "R", "AH0"}); ok {
		fmt.Println("found: ", n.Repr())
		result := n.GetDescendents(1)
		for _, v := range result {
			fmt.Printf("%v\n", v)
		}
	}
}
