var http = require('http'),
    httpProxy = require('http-proxy');
httpProxy.createServer({
  forward: {
    host: 'nano.nyc.corp.google.com',
    port: 9900
  }}).listen(8009);
