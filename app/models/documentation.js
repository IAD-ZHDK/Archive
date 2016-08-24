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
  files: DS.attr()
});
