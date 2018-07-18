import { Promise } from 'rsvp';
import Mixin from '@ember/object/mixin';

// FindByQuery is a Route mixin that allows specifying a query when searching
// for models.
export default Mixin.create({
  findByQuery(model, query, multi=false) {
    let newQuery = {};
    Object.keys(query).forEach(function(key){
      newQuery['filter[' + key + ']'] = query[key];
    });

    return new Promise((resolve, reject) => {
      this.store.query(model, newQuery).then(result => {
        if(multi) {
          resolve(result);
        } else {
          resolve(result.objectAt(0));
        }
      }, (err) => {
        reject(err);
      });
    });
  }
});
