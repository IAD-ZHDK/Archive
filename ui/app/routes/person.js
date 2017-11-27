import Ember from 'ember';

import FindByQuery from 'archive/mixins/find_by_query';

export default Ember.Route.extend(FindByQuery, {
  model(params) {
    return new Ember.RSVP.Promise((resolve, reject) => {
      this.findByQuery('person', {
        slug: params.slug,
      }).then((person) => {
        person.query('projects', {
          'filter[published]': true,
        }).then((projects) => {
          resolve({
            person: person,
            projects: projects,
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
