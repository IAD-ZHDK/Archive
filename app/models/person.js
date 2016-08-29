import DS from 'ember-data';
import HasManyQuery from 'ember-data-has-many-query';

export default DS.Model.extend(HasManyQuery.ModelMixin, {
  slug: DS.attr('string'),
  name: DS.attr('string'),
  documentations: DS.hasMany('documentation')
});
