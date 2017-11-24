import Ember from 'ember';
import AuthenticatedRouteMixin from 'ember-simple-auth/mixins/authenticated-route-mixin';

import AutomaticRollback from 'archive/mixins/automatic_rollback'

export default Ember.Route.extend(AuthenticatedRouteMixin, AutomaticRollback, {
  model(params) {
    return this.store.findRecord('user', params.id);
  }
});
