import EmberRouter from '@ember/routing/router';
import config from 'archive/config/environment';

const Router = EmberRouter.extend({
  location: config.locationType,
  rootURL: config.rootURL
});

Router.map(function() {
  this.route('documentation', { path: 'd/:slug' });
  this.route('person', { path: 'p/:slug' });
  this.route('tag', { path: 't/:slug' });
  this.route('year', { path: 'y/:year' });

  this.route('submit');

  this.route('admin', function(){
    this.route('login');

    this.route('documentations', function(){
      this.route('new');
      this.route('edit', { path: ':slug' });
    });

    this.route('people', function(){
      this.route('new');
      this.route('edit', { path: ':slug' });
    });

    this.route('tags', function(){
      this.route('new');
      this.route('edit', { path: ':slug' });
    });
  });
});

export default Router;
