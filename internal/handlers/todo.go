package handlers

import (
	"log"

	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
	"github.com/diazharizky/go-graphql-bootstrap/internal/resolvers"
)

type createTodoInput struct {
	UserID      string
	Description string
}

func (h handler) Todos() []*resolvers.TodoResolver {
	return resolvers.NewTodoList(h.appCtx)
}

func (h handler) CreateTodo(args struct{ Input createTodoInput }) *resolvers.TodoResolver {
	todo := models.Todo{
		UserID:      args.Input.UserID,
		Description: args.Input.Description,
	}

	if err := h.appCtx.TodoRepository.Create(&todo); err != nil {
		log.Printf("Error unable to create todo record: %s", err.Error())
		return &resolvers.TodoResolver{}
	}

	todoResolver := resolvers.TodoResolver{
		Todo: todo,
	}

	return &todoResolver
}
