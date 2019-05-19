package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	const defaultNumOfLines = 10
	var n int
	flag.IntVar(&n, "n", defaultNumOfLines, "number of lines to display")

	flag.Parse()
	args := flag.Args()

	for _, filename := range args {
		lines, err := readLine(filename)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Printf("===> %s <===\n", filename)
		startLine := 0
		if n < len(lines) {
			startLine = len(lines) - n
		}
		for i := startLine; i < len(lines); i++ {
			fmt.Println(lines[i])
		}
	}
}

func readLine(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
