import Ember from 'ember';

import BasicOperations from 'archive/mixins/basic_operations';

export default Ember.Controller.extend(BasicOperations, {
  afterCreateRoute: 'admin.users'
});
