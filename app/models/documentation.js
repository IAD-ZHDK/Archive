import DS from 'ember-data';

export default DS.Model.extend({
  slug: DS.attr('string'),
  madekId: DS.attr('string'),
  madekCoverId: DS.attr('string'),
  published: DS.attr('boolean'),
  title: DS.attr('string'),
  subtitle: DS.attr('string'),
  abstract: DS.attr(),
  cover: DS.attr(),
  videos: DS.attr(),
  images: DS.attr(),
  documents: DS.attr(),
  files: DS.attr(),
  people: DS.hasMany('person'),
  tags: DS.hasMany('tag'),

  peopleNames: Ember.computed.mapBy("people", "name"),
  tagNames: Ember.computed.mapBy("tags", "name"),
  peopleList: Ember.computed("peopleNames", function(){
    return this.get("peopleNames").join(", ");
  }),
  tagsList: Ember.computed("tagNames", function(){
    return this.get("tagNames").join(", ");
  }),
  madekUrl: Ember.computed("madekId", function(){
    return "https://medienarchiv.zhdk.ch/sets/" + this.get('madekId')
  })
});
