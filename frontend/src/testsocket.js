const WebSocket = require('ws');

let msg = {
  name: 'channel add',
  data: {
    name: 'Some new channel'
  }
}

let subMsg = {
  name: 'channel subscribe'
}

let ws = new WebSocket('ws://localhost:4000');

ws.onopen = () => {
  ws.send(JSON.stringify(subMsg))
  ws.send(JSON.stringify(msg))
}

ws.onmessage = (e) => {
  console.log(e.data);
}
