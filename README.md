url-shortener using go and redis

```
- controllers    // all handler stuff
  - analytics.go // for url analytics
  - url.go       // for creating and redirecting urls

- models   // all db related stuff
  - db.go  // initiate db connection
  - url.go // db functions for urls

- static
  - index.html // try out the api

- utils
  - base64.go // encoding the urls
  - url.go   //  url helper function

- ...
```
