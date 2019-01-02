# GraphiQL Handler
Simple Golang http.Handler to wrap and serve GraphiQL alongside the native graphql endpoint.
- <b>HTTP GET</b> requests serve the GraphiQL webclient
- <b>HTTP POST</b> requests serve graphql queries

## Usage
```Go
api := graphql.Handler(schema) // handler.New(...)

http.Handle("/graphql", graphiql.Handler(api))
http.ListenAndServe("localhost:3000", nil)
```

## Related Projects
- Playlyfe - ["github.com/playlyfe/go-graphql"](https://github.com/playlyfe/go-graphql)
- Thunder - ["github.com/samsarahq/thunder"](https://github.com/samsarahq/thunder)
