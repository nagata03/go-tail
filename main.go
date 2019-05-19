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
	text := readLine(args[0])
	for i := len(text) - n; i < len(text); i++ {
		fmt.Println(text[i])
	}

	// if err := readLine("hoge.txt"); err != nil {
	// 	fmt.Println(os.Stderr, err)
	// 	os.Exit(1)
	// }
}

func readLine(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", filename, err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", filename, err)
	}
	return lines
}
