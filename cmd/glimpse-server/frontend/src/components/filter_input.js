import React from 'react';
var Actions = require('../actions');

var FilterInput = React.createClass({
  filterList: function(event) {
    var value = event.target.value;
    Actions.filter(value);
  },

  render: function() {
    return (
      <input type='search' onChange={this.filterList} name='search' />
    );
  }
});

module.exports = FilterInput;
