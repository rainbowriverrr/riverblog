package models

import (
	"github.com/pocketbase/pocketbase"
	pbModels "github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

var _ pbModels.Model = (*Author)(nil)

type Author struct {
	pbModels.BaseModel

	FirstName string `db:"first_name, notnull"`
	LastName  string `db:"last_name"`
}

func (a *Author) TableName() string {
	return "authors"
}

// initializes the collection if it does not exist
func (a *Author) InitCollection(app *pocketbase.PocketBase) error {

	exists, _ := CollectionExists(app, "authors")

	if exists {
		return nil
	}

	collection := &pbModels.Collection{
		Name:       a.TableName(),
		Type:       pbModels.CollectionTypeBase,
		ListRule:   nil,
		ViewRule:   nil,
		CreateRule: nil,
		UpdateRule: nil,
		DeleteRule: nil,
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "first_name",
				Type:     schema.FieldTypeText,
				Required: true,
			},
			&schema.SchemaField{
				Name:     "last_name",
				Type:     schema.FieldTypeText,
				Required: false,
			},
		),
	}

	return app.Dao().SaveCollection(collection)

}
