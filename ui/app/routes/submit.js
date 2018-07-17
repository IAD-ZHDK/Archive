import Route from '@ember/routing/route';

export default Route.extend({
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
