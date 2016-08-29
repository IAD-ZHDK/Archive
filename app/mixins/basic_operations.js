import Ember from 'ember';

export default Ember.Mixin.create({
  transitionWithModel: false,
  setError(failure) {
    if(failure.errors && failure.errors.length > 0) {
      this.set('error', failure.errors[0].title.capitalize());
    } else {
      this.set('error', failure.capitalize());
    }

    setTimeout(() => {
      this.set('error', null);
    }, 5000);
  },
  setAttribute(key, value) {
    this.get('model').set(key, value);
  },
  saveModel(route) {
    this.set('formClass', 'loading');

    this.get('model').save().then(() => {
      this.set('formClass', null);

      if(this.get('transitionWithModel')) {
        this.transitionToRoute(this.get(route), this.get('model'));
      } else {
        this.transitionToRoute(this.get(route));
      }
    }).catch(failure => {
      this.set('formClass', null);
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
