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
	inputs = ``

	queries = `
		type Query {
			getUserByEmail(email: String!): User!
		}
	`
	mutations = `
		type Mutation {
			
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
