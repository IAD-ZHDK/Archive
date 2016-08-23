import DS from 'ember-data';

export default DS.Model.extend({
  title: DS.attr('string'),
  slug: DS.attr('string'),
  madekSet: DS.attr('string'),
  videos: DS.attr(),
  images: DS.attr(),
  files: DS.attr()
});
