package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/rainbowriverrr/riverblog/pkg/models"
)

type App struct {
	*pocketbase.PocketBase
	templateRegister *TemplateRegistry
}

func main() {

	templates, err := newTemplateCache()
	if err != nil {
		panic(err)
	}

	app := &App{
		PocketBase: pocketbase.New(),
		templateRegister: &TemplateRegistry{
			templates: templates,
		},
	}

	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {

		log.Println("Initializing collections...")

		// initialize collections
		author := &models.Author{}
		err := author.InitCollection(app.PocketBase)
		if err != nil {
			log.Println(err)
			return err
		} else {
			log.Println("Collection authors initialized")
		}

		tag := &models.Tag{}
		err = tag.InitCollection(app.PocketBase)
		if err != nil {
			log.Println(err)
			return err
		} else {
			log.Println("Collection tags initialized")
		}

		post := &models.Post{}
		err = post.InitCollection(app.PocketBase)
		if err != nil {
			log.Println(err)
			return err
		} else {
			log.Println("Collection posts initialized")
		}

		postTag := &models.PostTag{}
		err = postTag.InitCollection(app.PocketBase)
		if err != nil {
			log.Println(err)
			return err
		} else {
			log.Println("Collection post_tags initialized")
		}

		return nil

	})

	app.OnBeforeServe().Add(initRoutes)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Renderer = app.templateRegister
		return nil
	})

	if err := app.Start(); err != nil {
		panic(err)
	}
}
