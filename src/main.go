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

func read_file(f_path *string) ([]Recipe, error) {
	var content []byte
	var err error

	_, err = os.Stat(*f_path)
	if err != nil {
		return nil, fmt.Errorf("cannot open the file: %w", err)
	}

	content, err = os.ReadFile(*f_path)

	if err != nil {
		return nil, fmt.Errorf("cannot read the file: %w", err)
	}

	var recipesData RecipeData

	switch {
	case strings.HasSuffix(*f_path, ".json"):
		err = json.Unmarshal(content, &recipesData)
		if err != nil {
			return nil, fmt.Errorf("invalid json format: %w", err)
		}
	case strings.HasSuffix(*f_path, ".xml"):
		err = xml.Unmarshal(content, &recipesData)
		if err != nil {
			return nil, fmt.Errorf("invalid xml format: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported format")
	}
	return recipesData.Cakes, nil
}

func print_json(recipes []Recipe) string {
	var result []byte
	var err error
	result, err = json.MarshalIndent(recipes, "", " ")
	if err != nil {
		return fmt.Sprintf("error while printing: %w", err)
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
	var recipes []Recipe
	var err error
	recipes, err = read_file(f_path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(print_json(recipes))
}
