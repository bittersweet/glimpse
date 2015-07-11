import React from 'react';
var MessageStore = require('../stores/message_store');
var FilterStore = require('../stores/filter_store');

var Message = require('./message');

var MessageList = React.createClass({
  getInitialState: function() {
    return {
      messages: MessageStore.messages,
      highlight: '',
    };
  },

  onNewMessage: function(messages) {
    this.setState({
      messages: messages
    });
    window.scrollTo(0, document.body.scrollHeight);
  },

  onFilter: function(filter) {
    this.setState({
      highlight: filter
    });
  },

  componentDidMount: function() {
    this.unsubscribeMessageStore = MessageStore.listen(this.onNewMessage);
    this.unsubscribeFilterStore = FilterStore.listen(this.onFilter);
  },

  componentWillUnmount: function() {
    this.unsubscribeMessageStore();
    this.unsubscribeFilterStore();
  },

  render: function() {
    var highlight = this.state.highlight;
    var messages = this.state.messages.map(function(message) {
      return (
        <Message key={message.id} timestamp={message.timestamp} line={message.line} highlight={highlight} />
      );
    });

    return (
      <ul id='messages'>
        {messages}
      </ul>
    );
  },
});

module.exports = MessageList;
