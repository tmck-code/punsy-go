package main

import (
	"fmt"
	"github.com/tmck-code/go-colours"
	"github.com/tmck-code/punsy-go/src/dictionary"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fileUrl := "http://svn.code.sf.net/p/cmusphinx/code/trunk/cmudict/cmudict-0.7b"
	// dictionary.DownloadFile(fileUrl, ofpath)
	args := os.Args[1:]

	msg := args[0]
	offset, _ := strconv.Atoi(args[1])

	c := dictionary.NewCMU()
	ofpath := "cmudict-0.7b.utf8"
	dictionary.LoadCMU(ofpath, c)

	parts := strings.Split(msg, " ")
	Colours.CPrintln(Colours.Green, fmt.Sprint("parts:\t", parts))

	word := parts[len(parts)-1]
	fmt.Println("word:\t", word)

	pron := c.GetPronunciation(word)
	fmt.Println("pronunciation:\t", word, pron, pron[offset:])
	pron = pron[offset:]

	if results, ok := c.GetRhymes(pron); ok {
		fmt.Println("rhymes:\t", len(results), results)
		choice := results[rand.Int()%len(results)].(string)
		fmt.Println("choice:\t", choice)
		parts[len(parts)-1] = choice
	}
	fmt.Println("pun:\t", strings.Join(parts, " "))
}
