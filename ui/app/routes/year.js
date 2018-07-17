import { hash } from 'rsvp';
import Route from '@ember/routing/route';

export default Route.extend({
  model(params) {
    return hash({
      year: params.year,
      projects: this.store.query('project', {
        'filter[year]': params.year,
        'filter[published]': true
      })
    });
  }
});
