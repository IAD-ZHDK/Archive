import Ember from 'ember';

import FindByQuery from 'archive-app/mixins/find_by_query';

import Person from 'archive-app/models/person';

export default Ember.Route.extend(FindByQuery, {
  model(params) {
    return new Person({
      name: params.name,
    });
  },
  serialize(model) {
    return {
      name: model.get('name')
    };
  }
});
