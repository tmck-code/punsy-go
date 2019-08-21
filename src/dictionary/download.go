package dictionary

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"bufio"
	"unicode"
	"strings"
)

func ShouldSkip(line string) bool {
	if !unicode.IsLetter(rune(line[0])) {
		return true
	} else if strings.Contains(line, "(") {
		return true
	}
	return false
}

func DownloadFile(url string, ofpath string) error {
	// Get the file size
	headResp, err := http.Head(url)
	if err != nil { return err }
	size, err := strconv.Atoi(headResp.Header.Get("Content-Length"))
	headResp.Body.Close()

	// Open the output file
	out, err := os.Create(ofpath)
	if err != nil { return err }
	defer out.Close()

	// Download the file
	resp, err := http.Get(url)
	if err != nil { return err }
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(out)

	current := 0
	for scanner.Scan() {
		line := scanner.Text()
		if !ShouldSkip(line) {
			fmt.Fprintln(writer, strings.Join(strings.SplitN(line, "  ", 2), "|"))
		}
		current += len(line)
		fmt.Printf("\rDownloading... %0.2f%% complete", float64(current) / float64(size) * 100)
    }

    err = scanner.Err()
	return err
}

func Fetch(url string, ofpath string) {
	fmt.Println("Download Started")

	start := time.Now()

	err := DownloadFile(url, ofpath)
	if err != nil { panic(err) }

	elapsed := time.Since(start)
	fmt.Printf("\nDownload completed in %s\n", elapsed)
}
