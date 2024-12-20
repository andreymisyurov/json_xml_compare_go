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

func handleUnitCompare(unit_old, unit_new, recipe_name, ingredient_name string) {
	if unit_old != unit_new {
		if unit_new == "" {
			fmt.Printf("REMOVED unit for ingredient \"%s\" for cake \"%s\"\n", unit_old, recipe_name)
		} else if unit_old == "" {
			fmt.Printf("ADDED unit for ingredient \"%s\" for cake \"%s\"\n", unit_new, recipe_name)
		} else {
			fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
				ingredient_name, recipe_name, unit_new, unit_old)
		}
	}
}

func handleIngredientChanges(ing_old, ing_new map[string]Ingredient, recipe_name string) {
	for _, unit := range ing_new {
		if _, exists := ing_old[unit.Name]; !exists {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", unit.Name, recipe_name)
		}
	}
	for _, unit := range ing_old {
		if _, exists := ing_new[unit.Name]; !exists {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", unit.Name, recipe_name)
		} else {
			if unit.Count != ing_new[unit.Name].Count {
				fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", unit.Name, recipe_name, ing_new[unit.Name].Count, unit.Count)
			}
			handleUnitCompare(unit.Unit, ing_new[unit.Name].Unit, recipe_name, unit.Name)
		}
	}
}

func handleModifiedCake(map_old, map_new map[string]Recipe) {
	for _, old_recipe := range map_old {
		new_recipe, exists := map_new[old_recipe.Name]
		if exists {
			if old_recipe.Time != new_recipe.Time {
				fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n",
					old_recipe.Name, new_recipe.Time, old_recipe.Time)
			}
			handleIngredientChanges(getIngrMap(old_recipe.Ingredients), getIngrMap(new_recipe.Ingredients), old_recipe.Name)
		}
	}
}

func handleCakes(map_old, map_new map[string]Recipe) {
	for _, recipe := range map_new {
		if _, exists := map_old[recipe.Name]; !exists {
			fmt.Printf("ADDED cake \"%s\"\n", recipe.Name)
		}
	}
	for _, recipe := range map_old {
		if _, exists := map_new[recipe.Name]; !exists {
			fmt.Printf("REMOVED cake \"%s\"\n", recipe.Name)
		} else {
			handleModifiedCake(map_old, map_new)
		}
	}
}

func Compare(old DBReader, new DBReader) error {
	handleCakes(old.GetMap(), new.GetMap())
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
