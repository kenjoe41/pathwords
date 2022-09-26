package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	paths := getInput()

	// TODO: Add deduplicating logic with a map.
	for _, path := range paths {
		path_parts := strings.Split(path, "/")
		for _, path_part := range path_parts {
			if path_part != "" { //Remove empty spaces
				fmt.Println(path_part)
			}
		}
	}
}

func getInput() []string {
	// Check for stdin input
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		log.Fatal("Something went wrong.")
	}

	sc := bufio.NewScanner(os.Stdin)

	paths := []string{}

	for sc.Scan() {
		paths = append(paths, sc.Text())
	}

	// check there were no errors reading stdin (unlikely)
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	return paths
}
