package main

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/app"
	"github.com/diazharizky/go-graphql-bootstrap/internal/repositories"
	"github.com/diazharizky/go-graphql-bootstrap/internal/server"
	"github.com/diazharizky/go-graphql-bootstrap/pkg/db"
)

func main() {
	appCtx := app.NewContext()

	dbConn := db.GetConnection()

	appCtx.UserRepository = repositories.NewUserRepository(dbConn)
	appCtx.TodoRepository = repositories.NewTodoRepository(dbConn)

	server.Run(appCtx)
}
