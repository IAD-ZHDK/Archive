import Ember from 'ember';

import FindByQuery from 'archive-app/mixins/find_by_query';

export default Ember.Route.extend(FindByQuery, {
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
