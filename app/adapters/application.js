import DS from 'ember-data';

import config from 'archive-app/config/environment';

export default DS.JSONAPIAdapter.extend({
  host: config.apiBaseURL
});
