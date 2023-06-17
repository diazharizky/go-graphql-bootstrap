package resolvers

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/app"
	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

type TodoResolver struct {
	Todo models.Todo
}

func (r TodoResolver) ID() string {
	return r.Todo.ID.Hex()
}

func (r TodoResolver) UserID() string {
	return r.Todo.UserID
}

func (r TodoResolver) Description() string {
	return r.Todo.Description
}

func NewTodoList(appCtx *app.Ctx, userID ...string) *[]*TodoResolver {
	filter := bson.M{}

	if len(userID) > 0 {
		filter["userId"] = userID[0]
	}

	todos, err := appCtx.TodoRepository.List(filter)
	if err != nil {
		return nil
	}

	var todoList = make([]*TodoResolver, len(todos))

	for i, todo := range todos {
		todoList[i] = &TodoResolver{todo}
	}

	return &todoList
}
