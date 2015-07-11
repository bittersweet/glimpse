import React from 'react';

var Actions = require('./actions');

var MessageList = require('./components/message_list');
var FilterInput = require('./components/filter_input');

var App = React.createClass({
  render: function() {
    return (
      <div>
        <FilterInput />
        <MessageList />
      </div>
    );
  }
});

var host = new URL(window.location).hostname;
var client = new EventSource('http://' + host + ':8001');
client.onmessage = function(message) {
  Actions.newMessage(message);
};

React.render(<App />, document.body);
