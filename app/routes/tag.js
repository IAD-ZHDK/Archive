import Ember from 'ember';

import FindByQuery from 'archive-app/mixins/find_by_query';

export default Ember.Route.extend(FindByQuery, {
  model(params) {
    return new Promise((resolve, reject) => {
      this.findByQuery('tag', {
        slug: params.slug,
      }).then((tag) => {
        tag.query('documentations', {
          'filter[published]': true,
        }).then((documentations) => {
          resolve({
            tag: tag,
            documentations: documentations,
          });
        }).catch(reject);
      }).catch(reject);
    });
  },
  serialize(model) {
    return {
      slug: model.get('slug')
    };
  }
});
