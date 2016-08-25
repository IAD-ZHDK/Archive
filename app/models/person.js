import Ember from 'ember';

export default Ember.Object.extend({
  urlName: Ember.computed('name', function(){
    return encodeURIComponent(this.get('name'));
  })
});
