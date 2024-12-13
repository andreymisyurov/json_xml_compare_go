package main

import (
	"fmt"
	"strings"
	"flag"
	"os"
)

func main() {

	var f_path *string = nil;
	f_path = flag.String("f", "", "path to the file");
	flag.Parse();

	if *f_path == "" {
		fmt.Println("err: path is empty");
		os.Exit(1);
	}
	
	var err error;
	_, err = os.Stat(*f_path);
	if err != nil {
		fmt.Println("err: open file");
		os.Exit(1)
	}

	switch {
	case strings.HasSuffix(*f_path, ".json"):
		fmt.Println("it's Json");
	case strings.HasSuffix(*f_path, ".xml"):
		fmt.Println("it's XML");
	default:
		fmt.Println("unknown format");

	}
}
