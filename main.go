package graphiql

import (
    "net/http"
)

var graphiQL = []byte(`
<!DOCTYPE html>
<head>
  <style>body {height: 100vh; margin: 0; width: 100%; overflow: hidden;}</style>
  <link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/graphiql/0.12.0/graphiql.min.css" />

  <script crossorigin src="//unpkg.com/react@16.7.0/umd/react.production.min.js"></script>
  <script crossorigin src="//unpkg.com/react-dom@16.7.0/umd/react-dom.production.min.js"></script>
  <script crossorigin src="//cdn.jsdelivr.net/es6-promise/4.0.5/es6-promise.auto.min.js"></script>
  <script crossorigin src="//cdn.jsdelivr.net/fetch/0.9.0/fetch.min.js"></script>
  <script crossorigin src="//cdn.jsdelivr.net/npm/graphiql@0.12.0/graphiql.min.js"></script>

  <script>
    (function () {
      document.addEventListener('DOMContentLoaded', function () {
          endpoint = window.location.origin + '/graphql';
          function fetcher(params) {
            var options = {
              method: 'post',
              headers: {'Accept': 'application/json', 'Content-Type': 'application/json'},
              body: JSON.stringify(params),
              credentials: 'include',
            };
            return fetch(endpoint, options)
              .then(function (res) { return res.json() });
          }
          var body = React.createElement(GraphiQL, {
			  fetcher: fetcher,
			  query: '', variables: ''
		  });
          ReactDOM.render(body, document.body);
      });
    }());
  </script>
</head>
<body>
</body>
`)

func Handler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			w.Write(graphiQL)
			return
		}

		next.ServeHTTP(w, r)
	})
}
