package app

type Context struct {
	UserRepository IUserRepository
	TodoRepository ITodoRepository
}

func NewContext() *Context {
	return &Context{}
}
