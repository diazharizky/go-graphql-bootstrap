package models

type User struct {
	ID        int32  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Age       int32  `json:"age"`
}

type UserResolver struct {
	User User
}

func (r UserResolver) ID() int32 {
	return r.User.ID
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
