import Ember from 'ember';

export default Ember.Route.extend({
  model() {
    return this.store.createRecord('project');
  },
  actions: {
    willTransition() {
      if(this.controller.get('model.isNew')) {
        this.controller.get('model').unloadRecord();
      }
    }
  }
});
