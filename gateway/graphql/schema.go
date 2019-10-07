package graphql

import "fmt"

var (
	typeDefs = `
		type User {
			id: String!
			name: String!
			email: String!
			avatar: String!
			profile: Profile!
		}
		type Profile {
			experience: String!
			education: String!
			role: String!
			skills: [Skill]!
		}
		type Skill {
			name: String!
			level: String!
		}
		type Event {
			id: String!
			name: String!
			description: String!
			venue: String!
			date: Int!
			positions: [Position]!
		}
		type Position {
			name: String!
			company: String!
			location: String!
			description: String!
			experience: String!
			education: String!
			skills: [Skill]!
		}
	`
	inputs = `
		input SkillInput {
			name: String!
			level: String!
		}
		input SetProfileInput {
			experience: String!
			education: String!
			role: String!
			skills: [SkillInput]!
		}
		input CreateEventInput {
			name: String!
			description: String!
			venue: String!
			date: Int!
		}
		input CreatePositionInput {
			name: String!
			company: String!
			location: String!
			description: String!
			experience: String!
			education: String!
			skills: [SkillInput]!
		}
	`

	queries = `
		type Query {
			events: [Event]!
			event(id: String!): Event!
		}
	`
	mutations = `
		type Mutation {
			setProfile(input: SetProfileInput!): Boolean!
			createEvent(input: CreateEventInput!): Event!
			createPosition(input: CreatePositionInput!, eventID: String!): Boolean!
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
