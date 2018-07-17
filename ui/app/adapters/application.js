import DS from 'ember-data';
import DataAdapterMixin from 'ember-simple-auth/mixins/data-adapter-mixin';

import config from 'archive/config/environment';

export default DS.JSONAPIAdapter.extend(DataAdapterMixin, {
  host: config.apiBaseURL,
  namespace: 'api',

  authorize(xhr) {
    let { access_token } = this.get('session.data.authenticated');
    if (access_token) {
      xhr.setRequestHeader('Authorization', `Bearer ${access_token}`);
    }
  }
});
