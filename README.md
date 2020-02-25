# Buzzwords as a Service (BwaaS)

[![pipeline status](https://gitlab.com/koeniglorenz/bwaas/badges/master/pipeline.svg)](https://gitlab.com/koeniglorenz/bwaas/commits/master)

A simple HTTP server that responds to requests with a randomly generated concatenation of buzzwords that (probably) doesn't make any sense.

The server either responds with HTML or with JSON, when setting the `Accept`-Header of the request to `application/json`

### Buzzwords

The buzzwords are stored in a simple json file with the following structure:
```js
{
	"adj": [
      // adjectives like "Hosted or Automated"
	],
	"sub": [
      // nouns like "IoT" or "Big Data"
	],
	"app": [
      // appendices like "as a service"
	]
}
```
Every time a request is processed, a random combination of one of these three categories is generated.  
Feel free to add  buzzwords and open a pull request!

### Running locally

For running locally either golang (1.11) is required or the docker runtime.

_golang_
```sh
go build -o bwaas ./cmd/main.go
./bwaas --port 8080 --buzzwords ./buzzwords.json
```
With the flag `port` the port for the server to listen on can be specified.  
This is optional and defaults to `8080`.  
The flag `buzzwords` is for passing the path to the buzzwords-JSON-file.  
The flag is also optional and defaults to `buzzwords.json`.  

_docker_
```sh
docker build -t koeniglorenz/bwaas .
docker run -d -p 8080:8080 koeniglorenz/bwaas
```
By default the container exposes port `8080`, so this port can be mapped to any port on your local machine.
