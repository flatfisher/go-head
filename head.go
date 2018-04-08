package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	nOpt = flag.Int("n", 10, "Number of line to print")
	sOpt = flag.Bool("s", false, "Print lines number of file")
)

func getArgs() (int, []string) {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Invalid argument, you must specify a path.")
		os.Exit(1)
	}

	return len(args), args
}

func lineCounter(file *os.File) int {
	count := 0
	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			os.Exit(1)
			break
		}
		count++
	}
	return count
}

func printRows(file *os.File) {
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
	len, args := getArgs()
	for i, path := range args {
		file, err := os.Open(path)
		if err != nil {
			fmt.Println("Failed to open file.")
			os.Exit(2)
		}
		defer file.Close()

		if i >= 1 {
			fmt.Printf("\n")
		}

		if len != 1 {
			fmt.Printf("==> %s <==\n", path)
		}

		if *sOpt {
			//README: Skip line print
			rows := lineCounter(file)
			switch rows {
			case 0:
				fmt.Printf("%s has no line\n", path)
			case 1:
				fmt.Printf("%s has one line\n", path)
			default:
				fmt.Printf("%s has %d lines\n", path, rows)
			}
		} else {
			printRows(file)
		}
	}
}
