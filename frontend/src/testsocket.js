const WebSocket = require('ws');

let msg = {
  name: 'channel add',
  data: {
    name: 'Some new channel'
  }
}

let ws = new WebSocket('ws://localhost:4000');

ws.onopen = () => {
  ws.send(JSON.stringify(msg))
}