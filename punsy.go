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
	if result, ok := c.GetRhymes(result); ok {
		fmt.Println(len(result), result)
	}
}
