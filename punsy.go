package main

import (
	"fmt"
	"github.com/tmck-code/punsy-go/src/dictionary"
)

func main() {
	// fileUrl := "http://svn.code.sf.net/p/cmusphinx/code/trunk/cmudict/cmudict-0.7b"
	// dictionary.DownloadFile(fileUrl, ofpath)

	c := dictionary.NewCMU()
	ofpath := "cmudict-0.7b.utf8"
	dictionary.LoadCMU(ofpath, c)

	msg := "world"
	result := c.GetPronunciation(msg)
	fmt.Println(result)
}
