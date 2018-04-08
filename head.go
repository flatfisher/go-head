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

func get_args() (int, []string) {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Invalid argument, you must specify a path.")
		os.Exit(1)
	}

	return len(args), args
}

func print_rows(file *os.File) {
	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			os.Exit(1)
			break
		}
		fmt.Printf("%s\n", sc.Text())
		if i == *nOpt {
			break
		}
	}
}

func main() {
	len, args := get_args()
	for i, path := range args {
		file, err := os.Open(path)

		if i >= 1 {
			fmt.Printf("\n")
		}

		if len != 1 {
			fmt.Printf("==> %s <==\n", path)
		}

		if err != nil {
			fmt.Println("Failed to open file.")
			os.Exit(2)
		}
		defer file.Close()

		print_rows(file)
	}
}
