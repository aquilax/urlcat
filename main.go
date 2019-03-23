package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"path"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		url, err := url.Parse(line)
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(os.Stdout, path.Base(url.Path))
	}
}
