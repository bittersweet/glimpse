import React from 'react';

var PureRenderMixin = require('react/addons').addons.PureRenderMixin;

var Actions = require('../actions');

var Message = React.createClass({
  mixins: [PureRenderMixin],

  formattedLine: function() {
    const highlight = this.props.highlight;
    const line = this.props.line;

    if (highlight === '') {
      return line;
    }

    var re = RegExp('(' + highlight + ')', 'i');
    if (re.test(line)) {
      var newstr = line.replace(re, '<span class="highlight">$1</span>');
      Actions.messageMatched();
      return newstr;
    }

    return line;
  },

  render: function() {
    var raw = this.formattedLine();
    return (
      <li>
        <span className='timestamp'>{this.props.timestamp}</span>
        <span className='line' dangerouslySetInnerHTML={{__html: raw}} />
      </li>
    );
  }
});

module.exports = Message;
