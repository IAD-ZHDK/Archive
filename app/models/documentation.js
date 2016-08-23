import DS from 'ember-data';

export default DS.Model.extend({
  title: DS.attr('string'),
  slug: DS.attr('string'),
  madekSet: DS.attr('string'),
  madekCover: DS.attr('string'),
  cover: DS.attr(),
  videos: DS.attr(),
  images: DS.attr(),
  documents: DS.attr(),
  files: DS.attr()
});
