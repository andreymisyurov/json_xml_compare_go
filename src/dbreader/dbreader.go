package dbreader

import (
	"fmt"
	"strings"
)

type BaseReader struct {
	recipes []Recipe
}

type DBReader interface {
	Read(f_path *string) error
	ToString() string
	GetMap() map[string]Recipe
}

func Get_DB_reader(f_path *string) (DBReader, error) {
	if strings.HasSuffix(*f_path, ".json") {
		return &JSONReader{}, nil
	} else if strings.HasSuffix(*f_path, ".xml") {
		return &XMLReader{}, nil
	} else {
		return nil, fmt.Errorf("unsupported format")
	}
}

func Compare(old DBReader, new DBReader) error {
	var map_old map[string]Recipe
	var map_new map[string]Recipe
	map_old = old.GetMap()
	map_new = new.GetMap()

	for _, recipe := range map_new {
		if _, exists := map_old[recipe.Name]; !exists {
			fmt.Printf("ADDED cake \"%s\"\n", recipe.Name)
			continue
		}
	}

	for _, recipe := range map_old {
		if _, exists := map_new[recipe.Name]; !exists {
			fmt.Printf("REMOVED cake \"%s\"\n", recipe.Name)
			continue
		} else {
			if recipe.Time != map_new[recipe.Name].Time {
				fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", recipe.Name, recipe.Time, map_new[recipe.Name].Time)
			}
			var ing_map_o map[string]Ingredient
			var ing_map_n map[string]Ingredient
			ing_map_o = getIngrMap(recipe.Ingredients)
			ing_map_n = getIngrMap(map_new[recipe.Name].Ingredients)
			for _, unit := range ing_map_o {
				if _, exists := ing_map_n[unit.Name]; !exists {
					fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", unit.Name, recipe.Name)
					continue
				} else {
					// FIXME
				}
			}
		}
	}
	return nil
}

func (b *BaseReader) GetMap() map[string]Recipe {
	var recipe_map map[string]Recipe
	recipe_map = make(map[string]Recipe)
	for _, recipe := range b.recipes {
		recipe_map[recipe.Name] = recipe
	}
	return recipe_map
}

func getIngrMap(ingr []Ingredient) map[string]Ingredient {
	var ingr_map map[string]Ingredient
	ingr_map = make(map[string]Ingredient)
	for _, ingr := range ingr {
		ingr_map[ingr.Name] = ingr
	}
	return ingr_map
}
