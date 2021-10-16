package main

import (
	"fmt"
	"net/http"

	"github.com/f0ntana/go-begin/internal/comment"
	"github.com/f0ntana/go-begin/internal/database"
	transportHTTP "github.com/f0ntana/go-begin/internal/transport/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Settings Up Application")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDb(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setting up server")
		return err
	}

	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting application")
		fmt.Println(err)
	}
}
