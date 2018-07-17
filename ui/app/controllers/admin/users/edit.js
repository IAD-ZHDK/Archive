import Controller from '@ember/controller';

import BasicOperations from 'archive/mixins/basic_operations';

export default Controller.extend(BasicOperations, {
  afterUpdateRoute: 'admin.users',
  afterDeleteRoute: 'admin.users'
});
