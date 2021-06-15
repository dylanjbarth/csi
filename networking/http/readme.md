# proxy caching server

Expects a server to be running on port 8002 that it can forward client traffic to.

`go run main.go`

Connect to the proxy on port 8001 with curl, nc, or your browser. 


## Notes

- struggled with an issue where connections were hanging either to the dst server or the client. This seems to have been resolved by ensuring I'm always closing my file descriptors, but this still seems to happen with the concurrent tester.
- Another issue, occasionally I'd see 

```$ curl localhost:8002/hey --output -
curl: (18) transfer closed with 86 bytes remaining to read```

which seems to have been solved by setting the MSG_WAITALL bit in recvfrom while waiting for responses. 