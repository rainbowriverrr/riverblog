package models

import (
	"github.com/pocketbase/pocketbase"
)

// function to check if the collection exists
func CollectionExists(app *pocketbase.PocketBase, name string) (bool, error) {

	col, err := app.Dao().FindCollectionByNameOrId(name)

	if err != nil {
		return false, err
	}

	if col == nil {
		return false, nil
	}

	return true, nil

}
