import { ApolloServer, gql } from 'apollo-server';
import { DataSources } from 'apollo-server-core/dist/requestPipeline';
import EventAPI from './src/datasources/event';
import { IDataSources, IGraphqlContext } from './types';

const typeDefs = gql`
  type Event {
    name: String
    place: String
  }

  type Query {
    events: [Event]
  }
`;

const resolvers = {
  Query: {
    events: (
      _: undefined,
      __: undefined,
      { dataSources: { eventAPI }, userId }: IGraphqlContext,
    ) => eventAPI.getEventsForUser(userId),
  },
};

const server = new ApolloServer({
  dataSources: (): DataSources<IDataSources> => ({
    eventAPI: new EventAPI(),
  }),
  resolvers,
  typeDefs,
});

// The `listen` method launches a web server.
server.listen().then(({ url }: { url: string }) => {
  console.log(`🚀  Server ready at ${url}`); // tslint:disable-line no-console
});
