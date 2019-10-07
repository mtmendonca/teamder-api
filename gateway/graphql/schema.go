package graphql

import "fmt"

var (
	typeDefs = `
		type User {
			name: String!
			email: String!
			avatar: String
		}
	`
	inputs = `
		input CreateUserInput {
			name: String!
			email: String!
			avatar: String
		}
	`
	queries = `
		type Query {
			hello: String!
			sup: String!
		}
	`
	mutations = `
		type Mutation {
			createUser(input: CreateUserInput!): User!
		}
	`
)

// GetSchema returns the graphql schema for the public-facing api
func GetSchema() string {
	return fmt.Sprintf(`
	schema {
		query: Query
		mutation: Mutation
	}
	%s
	%s
	%s
	%s
`, typeDefs, inputs, queries, mutations)
}
