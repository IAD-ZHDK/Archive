import Route from '@ember/routing/route';
import AuthenticatedRouteMixin from 'ember-simple-auth/mixins/authenticated-route-mixin';

import FindByQuery from 'archive/mixins/find_by_query';

export default Route.extend(AuthenticatedRouteMixin, FindByQuery, {
  model(params) {
    return this.findByQuery('collection', {
      slug: params.slug
    });
  },
  serialize(model) {
    return {
      slug: model.get('slug')
    };
  }
});
