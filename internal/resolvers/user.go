package resolvers

import (
	"log"

	"github.com/diazharizky/go-graphql-bootstrap/internal/app"
	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
)

type UserResolver struct {
	appCtx *app.Ctx
	User   models.User
}

func (r UserResolver) ID() string {
	return r.User.ID.Hex()
}

func (r UserResolver) FirstName() string {
	return r.User.FirstName
}

func (r UserResolver) LastName() string {
	return r.User.LastName
}

func (r UserResolver) Email() string {
	return r.User.Email
}

func (r UserResolver) Age() int32 {
	return r.User.Age
}

func (r UserResolver) Todos() *[]*TodoResolver { // Must returns pointer type since it returns a resolver
	tdl := NewTodoList(r.appCtx, r.ID())

	return &tdl
}

func NewUser(appCtx *app.Ctx, id string) *UserResolver {
	ur := UserResolver{
		appCtx: appCtx,
	}

	user, err := appCtx.UserRepository.Get(id)
	if err != nil {
		return nil
	}

	ur.User = *user

	return &ur
}

func NewUserList(appCtx *app.Ctx) []UserResolver {
	users, err := appCtx.UserRepository.List()
	if err != nil {
		log.Printf("Error unable to retrieve user records: %s", err.Error())
		return nil
	}

	userList := make([]UserResolver, len(users))
	for i, u := range users {
		userList[i] = UserResolver{
			appCtx: appCtx,
			User:   u,
		}
	}

	return userList
}
