package dictionary

import (
	"bufio"
	"github.com/tmck-code/punsy-go/src/struct"
	"os"
	"strings"
)

type CMU struct {
	Pronunciations map[string][]string
	Rhymes *Struct.Trie
}

func NewCMU() *CMU {
	return &CMU{
		Pronunciations: make(map[string][]string, 0),
		Rhymes: Struct.NewTrie(),
	}
}

func LoadCMU(ifpath string, c *CMU) {
	in, _ := os.Open(ifpath)
	defer in.Close()

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		word := strings.SplitN(line, "|", 2)
		pronunciation := strings.Split(word[1], " ")
		c.Rhymes.Insert(pronunciation, word[0])
		c.Pronunciations[word[0]] = pronunciation
	}
}

func (c *CMU) GetPronunciation(s string) []string {
	s = strings.ToUpper(s)
	if el, ok := c.Pronunciations[s]; ok {
		return el
	} else {
		return []string{""}
	}
}
