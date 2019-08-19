package dictionary

import (
	"os"
	"bufio"
	"strings"
	"github.com/tmck-code/punsy-go/src/struct"
	"fmt"
)

func LoadCMU(ifpath string, t *Struct.Trie) {
	in, _ := os.Open(ifpath)
	defer in.Close()


	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(strings.Split(line, "|"))
	}
}
