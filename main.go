package main

import (
	"log"

	"github.com/diazharizky/go-graphql-bootstrap/internal/app"
	"github.com/diazharizky/go-graphql-bootstrap/internal/repositories"
	"github.com/diazharizky/go-graphql-bootstrap/internal/server"
	"github.com/diazharizky/go-graphql-bootstrap/pkg/mongodb"
)

func main() {
	appCtx := app.NewCtx()

	db, err := mongodb.GetDB()
	if err != nil {
		log.Fatalf("Error unable to get database: %v", err)
	}

	appCtx.UserRepository = repositories.NewUserRepository(db)
	appCtx.TodoRepository = repositories.NewTodoRepository(db)

	server.Run(appCtx)
}
