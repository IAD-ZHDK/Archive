import Ember from 'ember';

export default Ember.Route.extend({
  model(params) {
    return Ember.RSVP.hash({
      year: params.year,
      documentations: this.store.query('documentation', {
        'filter[year]': params.year,
        'filter[published]': true
      })
    });
  }
});
