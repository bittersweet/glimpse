import Reflux from 'reflux';
import Actions from '../actions';

var FilterStore = Reflux.createStore({
  listenables: [Actions],

  init: function() {
    this.filter = '';
  },

  getInitialState: function() {
    return this.filter;
  },

  onFilter: function(input) {
    console.log('onFilter: ', input);
    this.filter = input;
    this.trigger(input); // trigger so the classes listening to this store will get notified
  }
});

module.exports = FilterStore;
