package app

type Ctx struct {
	UserRepository IUserRepository
	TodoRepository ITodoRepository
}

func NewCtx() *Ctx {
	return &Ctx{}
}
