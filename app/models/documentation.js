import DS from 'ember-data';

import Person from 'archive-app/models/person';

export default DS.Model.extend({
  slug: DS.attr('string'),
  madekId: DS.attr('string'),
  madekCoverId: DS.attr('string'),
  title: DS.attr('string'),
  subtitle: DS.attr('string'),
  authors: DS.attr(),
  abstract: DS.attr(),
  cover: DS.attr(),
  videos: DS.attr(),
  images: DS.attr(),
  documents: DS.attr(),
  files: DS.attr(),

  authorsString: Ember.computed("authors", function(){
    return this.get('authors').join(", ");
  }),
  people: Ember.computed.map("authors", function(name){
    return new Person({
      name: name,
    });
  }),
  madekUrl: Ember.computed("madekId", function(){
    return "https://medienarchiv.zhdk.ch/sets/" + this.get('madekId')
  })
});
