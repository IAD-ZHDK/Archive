import DS from 'ember-data';

export default DS.Model.extend({
  slug: DS.attr('string'),
  name: DS.attr('string'),
  projects: DS.hasMany('project')
});
