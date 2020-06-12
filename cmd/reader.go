package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) {
	f, err := os.Open(path)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
