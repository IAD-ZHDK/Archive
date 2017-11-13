import Ember from 'ember';

export default Ember.Mixin.create({
  actions: {
    willTransition(transition) {
      if(this.controller.get('model.hasDirtyAttributes') &&
        !confirm('Are you sure you want to abandon progress?')) {
        transition.abort();
      } else {
        if(this.controller.get('model.isNew')) {
          this.controller.get('model').unloadRecord();
        } else {
          this.controller.get('model').rollbackAttributes();
        }
      }
    }
  }
});
