import Ember from 'ember';
import DS from 'ember-data';
import HasManyQuery from 'ember-data-has-many-query';

import config from 'archive-app/config/environment';

export default DS.JSONAPIAdapter.extend(HasManyQuery.RESTAdapterMixin, {
  session: Ember.inject.service(),
  host: config.apiBaseURL + '/api',
  headers: Ember.computed('session.password', function() {
    return {
      "Authorization": this.get('session.password'),
    };
  })
});
