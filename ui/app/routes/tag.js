import { Promise } from 'rsvp';
import Route from '@ember/routing/route';

import FindByQuery from 'archive/mixins/find_by_query';

export default Route.extend(FindByQuery, {
  model(params) {
    return new Promise((resolve, reject) => {
      this.findByQuery('tag', {
        slug: params.slug,
      }).then((tag) => {
        tag.query('projects', {
          'filter[published]': true,
        }).then((projects) => {
          resolve({
            tag: tag,
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
