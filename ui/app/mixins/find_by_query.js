import Ember from 'ember';

// FindByQuery is a Route mixin that allows specifying a query when searching
// for models.
export default Ember.Mixin.create({
  findByQuery(model, query) {
    let newQuery = {};
    Object.keys(query).forEach(function(key){
      newQuery['filter[' + key + ']'] = query[key];
    });

    return new Ember.RSVP.Promise((resolve, reject) => {
      this.store.query(model, newQuery).then(result => {
        resolve(result.objectAt(0));
      }, (err) => {
        reject(err);
      });
    });
  }
});
