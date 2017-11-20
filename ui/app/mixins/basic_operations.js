import Ember from 'ember';

// BasicOperations is a Controller mixin that takes care of the common model
// actions: create, update and delete.
export default Ember.Mixin.create({
  transitionWithModel: false,
  afterCreateRoute: 'application',
  afterUpdateRoute: 'application',
  afterDeleteRoute: 'application',
  setError(failure) {
    if(failure.errors && failure.errors.length > 0) {
      this.set('error', failure.errors[0].detail);
    } else {
      this.set('error', failure);
    }

    setTimeout(() => {
      this.set('error', null);
    }, 2000);
  },
  setAttribute(key, value) {
    this.get('model').set(key, value);
  },
  saveModel(route) {
    this.get('model').save().then(() => {
      if(this.get('transitionWithModel')) {
        this.transitionToRoute(this.get(route), this.get('model'));
      } else {
        this.transitionToRoute(this.get(route));
      }
    }).catch(failure => {
      this.setError(failure);
    });
  },
  actions: {
    create() {
      this.saveModel('afterCreateRoute');
    },
    update() {
      this.saveModel('afterUpdateRoute');
    },
    delete() {
      if(confirm('Do you really want to delete it?')) {
        this.get('model').destroyRecord().then((model) => {
          model.unloadRecord();
          this.transitionToRoute(this.get('afterDeleteRoute'));
        }).catch(failure => {
          this.setError(failure);
        });
      }
    }
  }
});
