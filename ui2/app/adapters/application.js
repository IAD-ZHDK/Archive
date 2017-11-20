import DS from 'ember-data';
import DataAdapterMixin from 'ember-simple-auth/mixins/data-adapter-mixin';

import config from 'archive/config/environment';

export default DS.JSONAPIAdapter.extend(DataAdapterMixin, {
  host: config.apiBaseURL,
  namespace: 'api',
  authorizer: 'authorizer:oauth2'
});
