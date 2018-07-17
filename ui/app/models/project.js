import { computed } from '@ember/object';
import { mapBy, map } from '@ember/object/computed';
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

  peopleNames: mapBy('people', 'name'),
  tagNames: mapBy('tags', 'name'),
  peopleList: computed('peopleNames', function(){
    return this.get('peopleNames').join(', ');
  }),
  tagsList: computed('tagNames', function(){
    return this.get('tagNames').join(', ');
  }),
  madekUrl: computed('madekId', function(){
    return 'https://medienarchiv.zhdk.ch/sets/' + this.get('madekId');
  }),
  websitesWithLinks: map('websites', function(website, i){
    return {
      title: website.title,
      link: config.apiBaseURL + '/web/' + this.get('id') + '/' + i.toString() + '/index.html'
    };
  })
});
