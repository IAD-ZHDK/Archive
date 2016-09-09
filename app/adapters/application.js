import DS from 'ember-data';
import HasManyQuery from 'ember-data-has-many-query';

import config from 'archive-app/config/environment';

export default DS.JSONAPIAdapter.extend(HasManyQuery.RESTAdapterMixin, {
  host: config.apiBaseURL + '/api'
});
