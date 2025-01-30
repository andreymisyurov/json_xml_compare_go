package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var f_path_1 *string = nil
	var f_path_2 *string = nil
	f_path_1 = flag.String("old", "", "path to the old recipe file")
	f_path_2 = flag.String("new", "", "path to the new recipe file")
	flag.Parse()

	if *f_path_1 == "" || *f_path_2 == "" {
		fmt.Println("err: path is empty")
		os.Exit(1)
	}
	var file *os.File
	var err error
	file, err = os.Open(*f_path_1)
	if err != nil {
		fmt.Printf("err: can not open %s\n", *f_path_1)
		os.Exit(1)
	}
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(file)

	var base map[string]bool
	base = make(map[string]bool)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		base[line] = true
	}
	err = scanner.Err()
	if err != nil {
		fmt.Printf("err: reading file %+v\n", err)
		os.Exit(1)
	}

	file, err = os.Open(*f_path_1)
	if err != nil {
		fmt.Printf("err: can not open %s\n", *f_path_1)
		os.Exit(1)
	}


}
