import { DataSource, DataSourceConfig } from 'apollo-datasource';
import { IDataSources, IEvent } from '../../types';

export default class EventAPI extends DataSource {
  private context = {};

  constructor() {
    super();
  }

  public initialize(config: DataSourceConfig<IDataSources>) {
    this.context = config.context;
  }

  public async getEventsForUser(userId: string): Promise<Array<IEvent | null>> {
    return [
      {
        name: 'Evento 1',
        place: 'Sala 1 - Goiânia',
      },
      {
        name: 'Evento 2',
        place: 'Sala 2 - Goiânia',
      },
      {
        name: 'Evento 3',
        place: 'Sala 3 - Goiânia',
      },
    ];
  }
}
