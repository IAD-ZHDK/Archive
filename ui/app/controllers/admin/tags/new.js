import Ember from 'ember';

import BasicOperations from 'archive-app/mixins/basic_operations';

export default Ember.Controller.extend(BasicOperations, {
  afterCreateRoute: 'admin.tags'
});
