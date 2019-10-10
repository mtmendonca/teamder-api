package graphql

import "fmt"

var (
	typeDefs = `
		type User {
			name: String!
			email: String!
			avatar: String!
		}
	`
	inputs = `
		input LogInInput {
			token: String!
			provider: String!
		}
	`

	queries = `
		type Query {}
	`
	mutations = `
		type Mutation {
			logIn(input: LogInInput!): String!
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
