import Ember from 'ember';
import DS from 'ember-data';

import config from 'archive/config/environment';

export default DS.Model.extend({
  slug: DS.attr('string'),
  madekId: DS.attr('string'),
  madekCoverId: DS.attr('string'),
  published: DS.attr('boolean'),
  title: DS.attr('string'),
  subtitle: DS.attr('string'),
  abstract: DS.attr('string'),
  year: DS.attr('string'),
  cover: DS.attr(),
  videos: DS.attr(),
  images: DS.attr(),
  documents: DS.attr(),
  websites: DS.attr(),
  files: DS.attr(),
  collections: DS.hasMany('collection'),
  people: DS.hasMany('person'),
  tags: DS.hasMany('tag'),

  peopleNames: Ember.computed.mapBy('people', 'name'),
  tagNames: Ember.computed.mapBy('tags', 'name'),
  peopleList: Ember.computed('peopleNames', function(){
    return this.get('peopleNames').join(', ');
  }),
  tagsList: Ember.computed('tagNames', function(){
    return this.get('tagNames').join(', ');
  }),
  madekUrl: Ember.computed('madekId', function(){
    return 'https://medienarchiv.zhdk.ch/sets/' + this.get('madekId');
  }),
  websitesWithLinks: Ember.computed.map('websites', function(website, i){
    return {
      title: website.title,
      link: config.apiBaseURL + '/web/' + this.get('id') + '/' + i.toString() + '/index.html'
    };
  })
});
