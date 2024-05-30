# Image Resizer Service
[![Github Actions](https://github.com/breadrock1/image-resizer/actions/workflows/go.yml/badge.svg)](https://github.com/breadrock1/image-resizer/actions/workflows/go.yml) 
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/breadrock1/image-resizer)

## Description
The service is designed for making a preview (creating an image with new dimensions based on a colored image).

## Architecture structure
The service is a web server (proxy) that loads images, scales/crops them to the desired format and returns them to the user.

## Endpoint description

The service must receive the URL of the original image, download it, change it to the required size and return it as an HTTP response.
 - Works only with HTTP requests.
 - Proxy errors from the remote service and respond to the client with 502 Bad Gateway.
 - Support only JPEG image encoding.

Http `GET` request

```http request
http://localhost:2891/fill/{height}/{width}/{img-url}
```

## Cache
The service saves (cache) the received preview on the local disk and, when requested again, send the image from the disk, without a request to the remote HTTP server.

Since the cache space is limited, the "Least Recent Used" algorithm must be used to remove rarely used images.

## Configuration 

There are multiple configurations parameters. It may be grouped by:
 - logger
 - server
 - cache
 - resizer
 - storage

There is example of `config.toml` file

```toml
[logger]
Level = "INFO"
FilePath = "logs/app.log"
EnableFileLog = false

[server]
Host="0.0.0.0"
HostPort=2891

[cache]
ExpireSeconds=20
CapacityValues=5

[resizer]
TargetQuality=90

[storage]
UseFilesystem=true
UploadDirectory="uploads"
```

## Launch and Deploy

The microservice should be deployed using the make run command (within docker compose up) in the project directory.

Build and run:
```shell
make run
```

Build and run docker container:
```shell
make run-img
```

Build and run by docker compose:
```shell
make run-compose
```

Run linter:
```shell
make lint
```

Run tests:
```shell
make test
```

## Testing tasks
It is necessary to check the operation of the server in different scenarios:

 - the image was found in the cache;
 - the remote server does not exist;
 - the remote server exists, but the image was not found (404 Not Found);
 - the remote server exists, but the image is not an image, but, say, an exe file;
 - the remote server returned an error;
 - the remote server returned the image;
 - the image is smaller than the required size; etc.

## Recension
Maximum - 15 points (subject to fulfillment of mandatory requirements):

 - Implemented an HTTP server that proxies requests to a remote server - 2 points.
 - Implemented image slicing - 2 points.
 - Caching sliced images on disk - 1 point.
 - Cache limit by location (LRU cache) - 1 point.
 - The proxy server correctly transmits the request headers - 1 point.
 - Integration tests written - 3 points.
 - The tests are adequate and fully cover the functionality - 1 point.
 - The project can be built using make build, run using make run and tested using make test - 1 point.
 - Code clarity and purity - up to 3 points.
