import { inject as service } from '@ember/service';
import Controller from '@ember/controller';

export default Controller.extend({
  session: service(),
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
          this.set('error', failure.errors[0].detail || failure.errors[0].title);
        } else {
          this.set('error', failure.capitalize());
        }
      });
    },
  }
});
