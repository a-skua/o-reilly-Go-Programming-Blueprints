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
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord,
	otherWord,
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets" + otherWord,
}

var transformsFile = flag.String("file", "", "transform rules file")

func main() {
	flag.Parse()
	if transformsFile != nil && *transformsFile != "" {
		content, err := ioutil.ReadFile(*transformsFile)
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
		transforms = strs
	}

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
