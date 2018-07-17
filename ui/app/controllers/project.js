import Controller from '@ember/controller';
import { computed } from '@ember/object';
import { htmlSafe } from '@ember/string';

export default Controller.extend({
  backgroundImage: computed('model.cover.high-res', function(){
    return htmlSafe(`background-image: url(${this.get('model.cover.high-res')})`)
  })
});
