package models

import (
	"github.com/pocketbase/pocketbase"
	pbModels "github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
)

var _ pbModels.Model = (*PostTag)(nil)

// pivot table for many-to-many relationship between posts and tags
type PostTag struct {
	pbModels.BaseModel
	PostID int64 `db:"post_id, reference:id"` // Foreign key, relates to Post.ID
	TagID  int64 `db:"tag_id, reference:id"`  // Foreign key, relates to Tag.ID
}

func (pt *PostTag) TableName() string {
	return "post_tags"
}

// initializes the collection if it does not exist
func (pt *PostTag) InitCollection(app *pocketbase.PocketBase) error {
	exists, _ := CollectionExists(app, "post_tags")
	if exists {
		return nil
	}

	postsCollection, _ := app.Dao().FindCollectionByNameOrId("posts")
	if postsCollection == nil {
		// TODO create posts collection
	}

	tagsCollection, _ := app.Dao().FindCollectionByNameOrId("tags")
	if tagsCollection == nil {
		// TODO create tags collection
	}

	collection := &pbModels.Collection{
		Name:       pt.TableName(),
		Type:       pbModels.CollectionTypeBase,
		ListRule:   nil,
		ViewRule:   nil,
		CreateRule: nil,
		UpdateRule: nil,
		DeleteRule: nil,
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "post_id",
				Type:     schema.FieldTypeRelation,
				Required: true,
				Unique:   false,
				Options: &schema.RelationOptions{
					CollectionId: postsCollection.GetId(),
				},
			},
			&schema.SchemaField{
				Name:     "tag_id",
				Type:     schema.FieldTypeRelation,
				Required: true,
				Unique:   false,
				Options: &schema.RelationOptions{
					CollectionId: tagsCollection.GetId(),
				},
			},
		),
	}

	return app.Dao().SaveCollection(collection)
}
