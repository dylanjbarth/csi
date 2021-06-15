# proxy caching server

Expects a server to be running on port 8002 that it can forward client traffic to.

`go run main.go`

Connect to the proxy on port 8001 with curl, nc, or your browser. 

## TODO 
- works for the initial request, but seems to get hung up on subsequent requests. 
- also sometimes seeing that it's just hanging when sending bytes to the origin server, or sending bytes to the client. Dosen't seem to be deterministic. Maybe I just need to reboot all servers before re-testing? 

NB I've moved around where I create the socket and connection with the destination server a lot to try and sort this out without much progress. I assumed I could re-use the socket but hit all kinds of broken pipe or socket not connected errors. 

Example of failing on subsequent requests. 

```$ go run main.go 
2021/06/14 21:10:27 Accepted conn from &{Port:61577 Addr:[127 0 0 1] raw:{Len:0 Family:0 Port:0 Addr:[0 0 0 0] Zero:[0 0 0 0 0 0 0 0]}}
2021/06/14 21:10:28 Message received from client: GET / HTTP/1.1
Host: localhost:8001
Connection: keep-alive
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Sec-GPC: 1
Sec-Fetch-Site: none
Sec-Fetch-Mode: navigate
Sec-Fetch-User: ?1
Sec-Fetch-Dest: document
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9


2021/06/14 21:10:28 Sending these bytes to our dst server
2021/06/14 21:10:28 Response from dst server: HTTP/1.0 200 OK
Server: BaseHTTP/0.6 Python/3.9.5
Date: Tue, 15 Jun 2021 01:10:28 GMT
Content-Length: 629

{
    "Host": "localhost:8001",
    "Connection": "keep-alive",
    "Upgrade-Insecure-Requests": "1",
    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36",
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
    "Sec-GPC": "1",
    "Sec-Fetch-Site": "none",
    "Sec-Fetch-Mode": "navigate",
    "Sec-Fetch-User": "?1",
    "Sec-Fetch-Dest": "document",
    "Accept-Encoding": "gzip, deflate, br",
    "Accept-Language": "en-US,en;q=0.9"
}
2021/06/14 21:10:28 Sending these bytes to our client
2021/06/14 21:10:28 Message received from client: 
2021/06/14 21:10:28 Got empty response from client, indicating orderly shutdown. Shutting down socket.
2021/06/14 21:10:28 Accepted conn from &{Port:61579 Addr:[127 0 0 1] raw:{Len:0 Family:0 Port:0 Addr:[0 0 0 0] Zero:[0 0 0 0 0 0 0 0]}}
2021/06/14 21:10:28 Message received from client: GET /favicon.ico HTTP/1.1
Host: localhost:8001
Connection: keep-alive
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36
Accept: image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8
Sec-GPC: 1
Sec-Fetch-Site: same-origin
Sec-Fetch-Mode: no-cors
Sec-Fetch-Dest: image
Referer: http://localhost:8001/
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9


2021/06/14 21:10:28 Sending these bytes to our dst server
2021/06/14 21:10:28 Response from dst server: GET /favicon.ico HTTP/1.1
Host: localhost:8001
Connection: keep-alive
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36
Accept: image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8
Sec-GPC: 1
Sec-Fetch-Site: same-origin
Sec-Fetch-Mode: no-cors
Sec-Fetch-Dest: image
Referer: http://localhost:8001/
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9


2021/06/14 21:10:28 Sending these bytes to our client
2021/06/14 21:10:28 Message received from client: 
2021/06/14 21:10:28 Got empty response from client, indicating orderly shutdown. Shutting down socket.
2021/06/14 21:10:33 Accepted conn from &{Port:61581 Addr:[127 0 0 1] raw:{Len:0 Family:0 Port:0 Addr:[0 0 0 0] Zero:[0 0 0 0 0 0 0 0]}}
2021/06/14 21:10:33 Message received from client: GET / HTTP/1.1
Host: localhost:8001
Connection: keep-alive
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Sec-GPC: 1
Sec-Fetch-Site: none
Sec-Fetch-Mode: navigate
Sec-Fetch-User: ?1
Sec-Fetch-Dest: document
Accept-Encoding: gzip, deflate, br
Accept-Language: en-US,en;q=0.9


2021/06/14 21:10:33 Sending these bytes to our dst server
2021/06/14 21:10:33 send failed. broken pipe
exit status 1```