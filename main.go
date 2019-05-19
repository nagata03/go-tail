package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if err := readLine("hoge.txt"); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}

func readLine(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil {
		return err
	}
	return nil
}
