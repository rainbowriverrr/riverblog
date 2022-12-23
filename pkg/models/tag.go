package models

import (
	"github.com/pocketbase/pocketbase"
	pbModels "github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

var _ pbModels.Model = (*Tag)(nil)

type Tag struct {
	pbModels.BaseModel
	Name string `db:"name, notnull, unique"` // Tag name
}

func (t *Tag) TableName() string {
	return "tags"
}

func (t *Tag) InitCollection(app *pocketbase.PocketBase) error {

	exists, _ := CollectionExists(app, "tags")

	if exists {
		return nil
	}

	collection := &pbModels.Collection{
		Name:       t.TableName(),
		Type:       pbModels.CollectionTypeBase,
		ListRule:   nil,
		ViewRule:   nil,
		CreateRule: nil,
		UpdateRule: nil,
		DeleteRule: nil,
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "name",
				Type:     schema.FieldTypeText,
				Required: true,
				Unique:   true,
			},
		),
	}

	return app.Dao().SaveCollection(collection)

}
