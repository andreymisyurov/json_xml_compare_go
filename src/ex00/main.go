package main

import (
	"flag"
	"fmt"
	"os"

	"src/dbreader"
)

func main() {
	var f_path *string = nil
	f_path = flag.String("f", "", "path to the file")
	flag.Parse()

	if *f_path == "" {
		fmt.Println("err: path is empty")
		os.Exit(1)
	}

	var db_reader dbreader.DBReader
	var err error
	db_reader, err = dbreader.Get_DB_reader(f_path)
	if err != nil || db_reader.Read(f_path) != nil {
		os.Exit(1)
	}

	fmt.Println(db_reader.ToString())
}
