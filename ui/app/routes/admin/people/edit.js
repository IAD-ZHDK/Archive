import Ember from 'ember';
import AuthenticatedRouteMixin from 'ember-simple-auth/mixins/authenticated-route-mixin';

import FindByQuery from 'archive/mixins/find_by_query';

export default Ember.Route.extend(AuthenticatedRouteMixin, FindByQuery, {
  model(params) {
    return this.findByQuery('person', {
      slug: params.slug
    });
  },
  serialize(model) {
    return {
      slug: model.get('slug')
    };
  }
});
