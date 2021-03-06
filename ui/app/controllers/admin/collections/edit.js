import Controller from '@ember/controller';

import BasicOperations from 'archive/mixins/basic_operations';

export default Controller.extend(BasicOperations, {
  afterUpdateRoute: 'admin.collections',
  afterDeleteRoute: 'admin.collections'
});
