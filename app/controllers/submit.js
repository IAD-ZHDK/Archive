import Ember from 'ember';

export default Ember.Controller.extend({
  formClass: '',
  actions: {
    create() {
      this.set('error', null);
      this.set('formClass', 'loading');

      this.get('model').save().then(() => {
        this.set('formClass', null);
        this.set('successful', true);
      }).catch(failure => {
        this.set('formClass', null);

        if(failure.errors && failure.errors.length > 0) {
          this.set('error', failure.errors[0].title.capitalize());
        } else {
          this.set('error', failure.capitalize());
        }
      });
    },
  }
});
