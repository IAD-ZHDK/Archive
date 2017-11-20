import Ember from 'ember';

export default Ember.Controller.extend({
  session: Ember.inject.service('session'),
  email: '',
  password: '',
  actions: {
    authenticate() {
      this.get('session')
        .authenticate('authenticator:oauth2', this.get('email'), this.get('password'))
        .catch((err) => { this.setError(err); });
    }
  },
  setError(failure) {
    this.set('error', failure['error_description'] || failure['error'] || 'Unknown Error');

    setTimeout(() => {
      this.set('error', null);
    }, 2000);
  },
});
