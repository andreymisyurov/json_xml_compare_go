package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type XMLReader struct {
	BaseReader
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

func (x XMLReader) ToString() string {
	var result []byte
	var err error
	result, err = json.MarshalIndent(x.recipes, "", " ")
	if err != nil {
		return fmt.Sprintf("error while printing: %+v", err)
	}
	return string(result)
}
