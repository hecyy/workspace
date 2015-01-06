var http = require('http');
var WebSocket = require('ws');

var id = 0;
var registry = {};

var registerEvent = function(evt, fn) {
  registry[evt] = fn;
}

var sendRequest = function(ws, method, opt_params) {
  var params = opt_params || {};
  ws.send(JSON.stringify({id: id++, method: method, params: params}));
}

var initWebSocket = function(body) {
  //var resp = JSON.parse(body);
  //var ws = new WebSocket(resp[0].webSocketDebuggerUrl);
  var ws = new WebSocket('ws://localhost:9999/devtools/page/1');
  ws.on('open', function() {
    sendRequest(ws, 'Page.enable');
    sendRequest(ws, 'Page.navigate', {url:'http://www.engadget.com'});
  });
  ws.on('message', handleMessages);
};

var handleMessages = function(msg) {
  var json = JSON.parse(msg);
  if (typeof json.id == 'number') {
    console.log('received response: %d', json.id);
  } else if (json.method) {
    console.log('received notification: %s', json.method);
  } else {
    console.log('received unexpected %s', msg);
  }
}

var init = function(host, port, path) {
  var options = {
    hostname: host,
    port: port,
    path: path,
    method: 'GET'
  };

  var req = http.request(options, function(res) {
    console.log('STATUS: ' + res.statusCode);
    console.log('HEADERS: ' + JSON.stringify(res.headers));
    res.setEncoding('utf8');
    res.on('data', initWebSocket);
  });

  req.on('error', function(e) {
    console.log('problem with request: ' + e.message);
  });

  req.end();
};

//init('localhost', 9222, '/json');
initWebSocket('');
