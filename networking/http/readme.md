# proxy caching server

Expects a destination server to be running on port 8002 that it can forward client traffic to.

To start the proxy: `go run main.go`

Connect a client to the proxy with curl, nc, or your browser. Full example: 

```
$ python3 server.py --port 8000  # start the destination server
$ go run main.go --port 8001 --dstPort 8000
$ curl localhost:8001/
{
    "Host": "localhost:8000",
    "User-Agent": "curl/7.64.1",
    "Accept": "*/*"
}
$ curl localhost:8001/c/thisiscached
{
    "Host": "localhost:8000",
    "User-Agent": "curl/7.64.1",
    "Accept": "*/*"
}
```

## Notes

- struggled with an issue where connections were hanging either to the dst server or the client. This seems to have been resolved by ensuring I'm always closing my file descriptors, but this still seems to happen with the concurrent tester.
- Another issue, occasionally I'd see 

```$ curl localhost:8002/hey --output -
curl: (18) transfer closed with 86 bytes remaining to read
```

which seems to have been solved by setting the MSG_WAITALL bit in recvfrom while waiting for responses. 

- why does my server hang sometimes sending bytes to the destination server (doesn't ever seem to timeout) - also seems non-deterministic because I can run it again and it handles many requests. 
- seems to work perfectly fine with curl, so assuming this has to do with concurrency issues OR something about an incomplete response from my proxy server to the dest server and python is hanging there waiting for more data not acking 