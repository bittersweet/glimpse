import Reflux from 'reflux';

var Actions = Reflux.createActions([
  'newMessage',
  'filter',
  'messageMatched',
]);

module.exports = Actions;
