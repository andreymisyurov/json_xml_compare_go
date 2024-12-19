package main

import (
	"flag"
	"fmt"
	"os"

	"src/dbreader"
)

func main() {
	var f_path_old *string = nil
	var f_path_new *string = nil
	f_path_old = flag.String("old", "", "path to the old recipe file")
	f_path_new = flag.String("new", "", "path to the new recipe file")
	flag.Parse()

	if *f_path_old == "" || *f_path_new == "" {
		fmt.Println("err: path is empty")
		os.Exit(1)
	}

	var db_reader_old dbreader.DBReader
	var db_reader_new dbreader.DBReader
	var err error

	db_reader_old, err = dbreader.Get_DB_reader(f_path_old)
	if err != nil {
		os.Exit(1)
	}

	db_reader_new, err = dbreader.Get_DB_reader(f_path_new)
	if err != nil {
		os.Exit(1)
	}

	db_reader_old.Read(f_path_old)
	fmt.Println(db_reader_old.ToString())
	db_reader_new.Read(f_path_new)
	fmt.Println(db_reader_new.ToString())

	dbreader.Compare(db_reader_old, db_reader_new)
}
