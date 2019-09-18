import EventAPI from './src/datasources/event';

export interface IDataSources {
  eventAPI: EventAPI;
}

export interface IGraphqlContext {
  userId: string;
  dataSources: IDataSources;
}

export interface IEvent {
  name: string;
  place: string;
}
