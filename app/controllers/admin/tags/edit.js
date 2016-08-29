import Ember from 'ember';

import BasicOperations from 'archive-app/mixins/basic_operations';

export default Ember.Controller.extend(BasicOperations, {
  afterUpdateRoute: 'admin.tags',
  afterDeleteRoute: 'admin.tags'
});
