# Testing Sony Camera Remote API with golang

I'm learning go and [did some tests](https://twitter.com/capitangolo/status/1264467208979263488) with the Sony Camera Remote API.

My first go code, so this follows no guidelines or best practices. Happy to get your feedback!

## Building

Just:

```
go build sony-api.go
```

## Using

```
./sony-api --endpoint http://your.camera.ip.address:port -action actTakePicture
```
