package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"strings"
)

type RecipeData struct {
	Cakes []Recipe `json:"cake" xml:"cake"`
}

type Recipe struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit" xml:"itemunit"`
}

type DBReader interface {
	Read(f_path *string) error
	toString() string
}

type JSONReader struct {
	recipes []Recipe
}

type XMLReader struct {
	recipes []Recipe
}

func (j *JSONReader) Read(f_path *string) error {
	var content []byte
	var err error
	content, err = os.ReadFile(*f_path)
	if err != nil {
		return fmt.Errorf("cannot read file: %w", err)
	}

	var recipesData RecipeData
	err = json.Unmarshal(content, &recipesData)
	if err != nil {
		return fmt.Errorf("invalid JSON format: %w", err)
	}

	j.recipes = recipesData.Cakes
	return nil
}

func (j JSONReader) toString() string {
	var result []byte
	var err error
	result, err = xml.MarshalIndent(j.recipes, "", " ")
	if err != nil {
		return fmt.Sprintf("error while printing: %+v", err)
	}
	return string(result)
}

func (x *XMLReader) Read(f_path *string) error {
	var content []byte
	var err error
	content, err = os.ReadFile(*f_path)
	if err != nil {
		return fmt.Errorf("cannot read file: %w", err)
	}

	var recipesData RecipeData
	err = xml.Unmarshal(content, &recipesData)
	if err != nil {
		return fmt.Errorf("invalid XML format: %w", err)
	}

	x.recipes = recipesData.Cakes
	return nil
}

func (x XMLReader) toString() string {
	var result []byte
	var err error
	result, err = json.MarshalIndent(x.recipes, "", " ")
	if err != nil {
		return fmt.Sprintf("error while printing: %+v", err)
	}
	return string(result)
}

func get_DB_reader(f_path *string) (DBReader, error) {
	if strings.HasSuffix(*f_path, ".json") {
		return &JSONReader{}, nil
	} else if strings.HasSuffix(*f_path, ".xml") {
		return &XMLReader{}, nil
	} else {
		return nil, fmt.Errorf("unsupported format")
	}
}

func print_json(recipes []Recipe) string {
	var result []byte
	var err error
	result, err = json.MarshalIndent(recipes, "", " ")
	if err != nil {
		return fmt.Sprintf("error while printing: %+v", err)
	}
	return string(result)
}

func main() {
	var f_path *string = nil
	f_path = flag.String("f", "", "path to the file")
	flag.Parse()

	if *f_path == "" {
		fmt.Println("err: path is empty")
		os.Exit(1)
	}

	var db_reader DBReader
	var err error
	db_reader, err = get_DB_reader(f_path)
	if err != nil || db_reader.Read(f_path) != nil {
		os.Exit(1)
	}

	fmt.Println(db_reader.toString())
}
