package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type JSONReader struct {
	BaseReader
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

func (j JSONReader) ToString() string {
	var result []byte
	var err error
	result, err = xml.MarshalIndent(j.recipes, "", " ")
	if err != nil {
		return fmt.Sprintf("error while printing: %+v", err)
	}
	return string(result)
}
