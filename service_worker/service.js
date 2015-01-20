this.addEventListener('install', function(event){
	console.log('installing');
	console.log(event)
	debugger;
  //event.waitUntil(caches.create('v1').then(function(cache) {
  //											console.log('populating cache');
  	//										return cache.add('/logo11w.png');
  });

this.addEventListener('activate', function(event){
	console.log('service worker activated.')
});

this.addEventListener('fetch', 
	                    function(e) {
  console.log(e.request.url);
  console.log(e.request.headers);
  console.log(e.request.cache);
},                    function(err) {	
});