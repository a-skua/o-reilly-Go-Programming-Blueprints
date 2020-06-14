package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var tlds = []string{"com", "net"}
var tldsFile = flag.String("file", "", "top-level-domain-list file")

const allowedChar = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
	flag.Parse()
	if tldsFile != nil && *tldsFile != "" {
		content, err := ioutil.ReadFile(*tldsFile)
		if err != nil {
			panic(err)
		}
		if len(content) == 0 {
			panic("file is empty")
		}
		var strs []string
		for _, str := range strings.Split(string(content), "\n") {
			if str == "" {
				continue
			}
			strs = append(strs, str)
		}
		tlds = strs
	}

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if !strings.ContainsRune(allowedChar, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
}
