package models

type Todo struct {
	ID          int32        `json:"id"`
	Description string       `json:"description"`
	Owner       UserResolver `json:"owner"`
}

type TodoResolver struct {
	Todo Todo
}

func (r TodoResolver) ID() int32 {
	return r.Todo.ID
}

func (r TodoResolver) Description() string {
	return r.Todo.Description
}

func (r TodoResolver) Owner() UserResolver {
	return r.Todo.Owner
}
