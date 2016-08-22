import Ember from 'ember';
import config from './config/environment';

const Router = Ember.Router.extend({
  location: config.locationType,
  rootURL: config.rootURL
});

Router.map(function() {
  this.route('admin', function(){
    this.route('login');
    this.route('documentations', function(){
      this.route('new');
      this.route('edit', { path: ':slug' });
    });
  });
});

export default Router;
