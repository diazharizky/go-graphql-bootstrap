package resolvers

import "github.com/diazharizky/go-graphql-bootstrap/internal/models"

type TodoResolver struct {
	Todo models.Todo
}

func (r TodoResolver) ID() string {
	return r.Todo.ID.String()
}

func (r TodoResolver) Description() string {
	return r.Todo.Description
}

func (r TodoResolver) Owner() UserResolver {
	return UserResolver{r.Todo.Owner}
}
