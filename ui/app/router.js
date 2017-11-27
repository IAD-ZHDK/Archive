import EmberRouter from '@ember/routing/router';
import config from 'archive/config/environment';

const Router = EmberRouter.extend({
  location: config.locationType,
  rootURL: config.rootURL
});

Router.map(function() {
  this.route('projects');
  this.route('collections');
  this.route('project', { path: 'project/:slug' });
  this.route('collection', { path: 'collection/:slug' });
  this.route('person', { path: 'person/:slug' });
  this.route('tag', { path: 'tag/:slug' });
  this.route('year', { path: 'year/:year' });

  this.route('submit');
  this.route('sign-in');

  this.route('admin', function(){
    this.route('projects', function(){
      this.route('new');
      this.route('show', { path: 'show/:slug' });
      this.route('edit', { path: 'edit/:slug' });
    });

    this.route('people', function(){
      this.route('new');
      this.route('edit', { path: ':slug' });
    });

    this.route('tags', function(){
      this.route('new');
      this.route('edit', { path: ':slug' });
    });

    this.route('users', function(){
      this.route('new');
      this.route('edit', { path: ':id' });
    });
  });
});

export default Router;
