// dup prints the count and text of lines that appear more than once
//  in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dupFiles := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		counts := make(map[string]int)
		countLines(os.Stdin, counts)
	} else {
		for i, arg := range files {
			counts := make(map[string]int)
			filename := files[i]
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			dupFiles[filename] = counts
			f.Close()
		}
	}
	for filename, counts := range dupFiles {
		fmt.Println("File: ", filename)
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}