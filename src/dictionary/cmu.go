package dictionary

import (
	"bufio"
	"github.com/tmck-code/punsy-go/src/struct"
	"os"
	"strings"
)

func LoadCMU(ifpath string, t *Struct.Trie) {
	in, _ := os.Open(ifpath)
	defer in.Close()

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		word := strings.SplitN(line, "|", 2)
		pronunciation := strings.Split(word[1], " ")
		t.Insert(word[0], pronunciation)
	}
}
