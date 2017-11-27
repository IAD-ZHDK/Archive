import Ember from 'ember';

export default Ember.Route.extend({
  model(params) {
    return Ember.RSVP.hash({
      year: params.year,
      projects: this.store.query('project', {
        'filter[year]': params.year,
        'filter[published]': true
      })
    });
  }
});
