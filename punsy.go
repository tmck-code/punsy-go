package main

import (
	"fmt"
	"github.com/tmck-code/punsy-go/src/dictionary"
	"os"
)

func main() {
	// fileUrl := "http://svn.code.sf.net/p/cmusphinx/code/trunk/cmudict/cmudict-0.7b"
	// dictionary.DownloadFile(fileUrl, ofpath)
	args := os.Args[1:]

	c := dictionary.NewCMU()
	ofpath := "cmudict-0.7b.utf8"
	dictionary.LoadCMU(ofpath, c)

	msg := args[0]
	result := c.GetPronunciation(msg)
	fmt.Println("pronunciation:", result)
	if rhyme, ok := c.Rhymes.Get(result); ok {
		fmt.Println(rhyme)
		data := rhyme.GetDescendentsData(10, make([]interface{}, 0))
		fmt.Println(data)
	}
}
