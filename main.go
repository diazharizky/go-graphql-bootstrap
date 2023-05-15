package main

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/app"
	"github.com/diazharizky/go-graphql-bootstrap/internal/repositories"
	"github.com/diazharizky/go-graphql-bootstrap/internal/server"
)

func main() {
	appCtx := app.NewContext()

	appCtx.UserRepository = repositories.NewUserRepository()
	appCtx.TodoRepository = repositories.NewTodoRepository()

	server.Run(appCtx)
}
