package graphql

import "fmt"

var (
	typeDefs = `
		type Event {
			id: String!
			name: String!
			description: String!
			venue: String!
			date: String!
		}
	`
	inputs = `
		input CreateEventInput {
			name: String!
			description: String!
			venue: String!
			date: String!
		}
	`

	queries = `
		type Query {
			events: [Event]!
		}
	`
	mutations = `
		type Mutation {
			createEvent(input: CreateEventInput!): Event!
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
