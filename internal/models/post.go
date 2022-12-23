package models

import (
	"github.com/pocketbase/pocketbase"
	pbModels "github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

var _ pbModels.Model = (*Post)(nil)

type Post struct {
	pbModels.BaseModel

	AuthorID    int64          `db:"author_id, reference:id"` // Foreign key, relates to Author.ID
	Title       string         `db:"title"`
	Content     string         `db:"content"`
	Status      string         `db:"status"`
	Slug        string         `db:"slug"`
	PublishedAt types.DateTime `db:"published_at"`
}

func (p *Post) TableName() string {
	return "posts"
}

// initializes the collection if it does not exist
func (p *Post) InitCollection(app *pocketbase.PocketBase) error {
	exists, _ := CollectionExists(app, "posts")
	if exists {
		return nil
	}

	authorsCollection, _ := app.Dao().FindCollectionByNameOrId("authors")
	if authorsCollection == nil {
		// TODO create authors collection
	}

	authorsCollectionId := authorsCollection.GetId()

	collection := &pbModels.Collection{
		Name:       p.TableName(),
		Type:       pbModels.CollectionTypeBase,
		ListRule:   nil,
		ViewRule:   nil,
		CreateRule: nil,
		UpdateRule: nil,
		DeleteRule: nil,
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "author_id",
				Type:     schema.FieldTypeRelation,
				Required: true,
				Unique:   false,
				Options: &schema.RelationOptions{
					CollectionId: authorsCollectionId,
				},
			},
			&schema.SchemaField{
				Name:     "title",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "content",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "status",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "slug",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "published_at",
				Type:     schema.FieldTypeDate,
				Required: false,
				Unique:   false,
			},
		),
	}

	return app.Dao().SaveCollection(collection)

}
