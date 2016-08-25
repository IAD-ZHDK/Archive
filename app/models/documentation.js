import DS from 'ember-data';

export default DS.Model.extend({
  slug: DS.attr('string'),
  madekId: DS.attr('string'),
  madekCoverId: DS.attr('string'),
  title: DS.attr('string'),
  subtitle: DS.attr('string'),
  abstract: DS.attr(),
  cover: DS.attr(),
  videos: DS.attr(),
  images: DS.attr(),
  documents: DS.attr(),
  files: DS.attr(),
  people: DS.hasMany("person"),

  peopleNames: Ember.computed.mapBy("people", "name"),
  authors: Ember.computed("peopleNames", function(){
    return this.get("peopleNames").join(", ");
  }),
  madekUrl: Ember.computed("madekId", function(){
    return "https://medienarchiv.zhdk.ch/sets/" + this.get('madekId')
  })
});
