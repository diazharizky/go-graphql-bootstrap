package handlers

import "github.com/diazharizky/go-graphql-bootstrap/internal/types"

const schema types.Schema = `
type Query {
	users: [User!]!
	todos: [Todo!]!
}

type Mutation {
	createUser(input: NewUser!): User
}

type User {
	id: String!
	firstName: String!
	lastName: String!
	email: String!
	age: Int!
}

input NewUser {
	firstName: String!
	lastName: String!
	email: String!
	age: Int!
}

type Todo {
	id: String!
	description: String!
	owner: User!
}
`
