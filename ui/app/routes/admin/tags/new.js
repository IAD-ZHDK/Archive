import Ember from 'ember';

import AutomaticRollback from 'archive-app/mixins/automatic_rollback';

export default Ember.Route.extend(AutomaticRollback, {
  model() {
    return this.store.createRecord('tag');
  }
});
