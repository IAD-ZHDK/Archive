import Ember from 'ember';

import FindByQuery from 'archive-app/mixins/find_by_query';

export default Ember.Route.extend(FindByQuery, {
  model(params) {
    return new Promise((resolve, reject) => {
      this.findByQuery('person', {
        slug: params.slug,
      }).then((person) => {
        person.query('documentations', {
          'filter[published]': true,
        }).then((documentations) => {
          resolve({
            person: person,
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
