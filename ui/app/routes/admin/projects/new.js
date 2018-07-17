import Route from '@ember/routing/route';
import AuthenticatedRouteMixin from 'ember-simple-auth/mixins/authenticated-route-mixin';

import AutomaticRollback from 'archive/mixins/automatic_rollback';

export default Route.extend(AuthenticatedRouteMixin, AutomaticRollback, {
  model() {
    return this.store.createRecord('project');
  }
});
