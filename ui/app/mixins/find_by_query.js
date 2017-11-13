import Ember from 'ember';

export default Ember.Mixin.create({
  findByQuery(model, query) {
    var newQuery = {};
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
