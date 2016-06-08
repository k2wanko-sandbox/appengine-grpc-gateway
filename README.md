# AppEngine gRPC Gateway example

# Setup

require `gcloud` `GoogleAppEngine for Go SDK` `node.js`

```
$ npm install
```

# Run

```
$ goapp serve
```

```
$ curl -v -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"value":"Hi"}'  http://localhost:8080/echo
```

Access: `http://localhost:8080/swagger-ui/index.html?url=http://localhost:8080/echo_service.swagger.json` (require admin)
