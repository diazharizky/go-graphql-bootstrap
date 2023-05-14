package handlers

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/models"
)

var schema models.Schema = `
type User {
	id: Int!
	firstName: String!
	lastName: String!
	email: String!
	age: Int!
}

type Todo {
	id: Int!
	description: String!
	owner: User!
}

type Query {
	users: [User!]!
	todos: [Todo!]!
}
`
