# Running locally

`go run main.go`

# API?

This thing really doesn't have an api but does just one thing.

Anything sent in the url is considered a "password" and is hashed with bcrypt. This is set to put some load on the box. Responses should take at least ~1s to generate.

Examples:

```
$ curl http://localhost:8080/hello
{"plain":"hello","hashed":"$2a$15$5mQzwyE40EQaHfBFKH5Ns.M.YDSdlYarXNZzKdW800mTnGzEn5iE."}

$ curl http://localhost:8080/
{"plain":"","hashed":"$2a$15$QuVNFv32QakKmeaYvxcRJ.hud4JE5DTGlUf.3XELEFooh63UBDXhG"}

$ curl http://localhost:8080/hello/i/am/a/password
{"plain":"hello/i/am/a/password","hashed":"$2a$15$8FEa7wkcEthIv1DdGjM8durXa2b9SrHMNy1FufCRWDZ1dHT5b3p5m"}
```


# How to ship a new release?

Publish code to github on master and docker hub's automated build service will publish a docker image @ `jhgaylor/k8s-load-demo-api`. The deployment pulls it from docker's public repository.

# How to crank everything up?

* Make sure that the k8s-load-demo-charts/api chart is installed in your kubernetes cluster

* start the proxy `kubectl proxy`

* load the service through the proxy by browsing to http://localhost:8001/api/v1/proxy/namespaces/default/services/load-demo-api/

