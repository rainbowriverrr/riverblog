package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/rainbowriverrr/riverblog/pkg/models"
)

func main() {
	app := pocketbase.New()

	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {

		log.Println("Initializing collections...")

		// initialize collections
		author := &models.Author{}
		err := author.InitCollection(app)
		if err != nil {
			log.Println(err)
			return err
		} else {
			log.Println("Collection authors initialized")
		}

		tag := &models.Tag{}
		err = tag.InitCollection(app)
		if err != nil {
			log.Println(err)
			return err
		} else {
			log.Println("Collection tags initialized")
		}

		post := &models.Post{}
		err = post.InitCollection(app)
		if err != nil {
			log.Println(err)
			return err
		} else {
			log.Println("Collection posts initialized")
		}

		postTag := &models.PostTag{}
		err = postTag.InitCollection(app)
		if err != nil {
			log.Println(err)
			return err
		} else {
			log.Println("Collection post_tags initialized")
		}

		return nil

	})

	app.OnBeforeServe().Add(initRoutes)

	if err := app.Start(); err != nil {
		panic(err)
	}
}
