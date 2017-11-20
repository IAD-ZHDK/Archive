import OAuth2PasswordGrant from 'ember-simple-auth/authenticators/oauth2-password-grant';

import config from 'archive/config/environment';

export default OAuth2PasswordGrant.extend({
  clientId: config.clientID,
  serverTokenEndpoint: config.apiBaseURL + '/auth/token',
  serverTokenRevocationEndpoint: config.apiBaseURL + "/auth/revoke"
});
