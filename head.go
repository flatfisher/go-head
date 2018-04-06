package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	nOpt = flag.Int("n", 0, "help message for n option")
)

func main() {
	file_path := flag.Args()
	flag.Parse()

	file, err := os.Open(file_path)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			// エラー処理
			break
		}
		fmt.Printf("%s\n", sc.Text())
	}
}
