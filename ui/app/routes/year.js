import { hash } from 'rsvp';
import Route from '@ember/routing/route';

import FindByQuery from 'archive/mixins/find_by_query';

export default Route.extend(FindByQuery, {
  model(params) {
    return hash({
      year: params.year,
      projects: this.findByQuery('project', {
        year: params.year,
        published: true
      }, true)
    });
  }
});
