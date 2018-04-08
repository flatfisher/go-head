package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	nOpt = flag.Int("n", 10, "help message for n option")
)

func main() {
	flag.Parse()
	file_path := flag.Args()

	if len(file_path) < 1 {
		fmt.Println("Invalid argument, you must specify a path.")
		os.Exit(1)
	}

	file, err := os.Open(file_path[0])
	if err != nil {
		fmt.Println("Failed to open file.")
		os.Exit(2)
	}
	defer file.Close()

	fmt.Println(*nOpt)
	sc := bufio.NewScanner(file)

	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			os.Exit(1)
			break
		}
		fmt.Printf("%s\n", sc.Text())
	}
}
