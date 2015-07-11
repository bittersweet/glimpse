import Reflux from 'reflux';
import Actions from '../actions';

var MessageStore = Reflux.createStore({
  listenables: [Actions],

  init: function() {
    this.messages = [];
  },

  getInitialState: function() {
    return this.messages;
  },

  onNewMessage: function(message) {
    message = JSON.parse(message.data);
    this.messages.push(message);
    if (this.messages.length > 50) {
      this.messages.shift();
    }
    this.trigger(this.messages);
  },

  onMessageMatched: function() {
    const audio = new Audio('/src/notification.mp3');
    audio.play();
  },
});

module.exports = MessageStore;
