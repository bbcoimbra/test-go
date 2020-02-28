# Simple HTTP REST API in Go

This is a test project. It contains my start point to start to study the go-lang.

## Compile and Run

First get the dependencies:

```go
go get -u github.com/gorilla/mux
go get -u github.com/shopspring/decimal
```

When dependecies are available run:

```go
go build
```

And run the binary:

```bash
$ ./test-go
```

To check if everything is fine, you can run the commands as follows:

```bash
$ curl -v -H "content-type: application/json" -d'{"document_number": "12345678922"}' localhost:8080/accounts
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> POST /accounts HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
> content-type: application/json
> Content-Length: 34
>
* upload completely sent off: 34 out of 34 bytes
< HTTP/1.1 201 Created
< Date: Fri, 28 Feb 2020 12:23:41 GMT
< Content-Length: 41
< Content-Type: text/plain; charset=utf-8
<
{"id":1,"document_number":"12345678922"}
* Connection #0 to host localhost left intact
```

```bash
$ curl -v -H "content-type: application/json" -d'{"account_id": 1, "operation_type_id": 4, "amount": 123.45}' localhost:8080/transactions
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> POST /transactions HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.54.0
> Accept: */*
> content-type: application/json
> Content-Length: 59
>
* upload completely sent off: 59 out of 59 bytes
< HTTP/1.1 201 Created
< Date: Fri, 28 Feb 2020 12:23:58 GMT
< Content-Length: 148
< Content-Type: text/plain; charset=utf-8
<
{"transaction_id":1,"account_id":1,"operation_type_id":4,"amount":123.45,"event_date":"2020-02-28T09:23:58.247599-03:00","decimal_amount":"123.45"}
* Connection #0 to host localhost left intact
```

```bash
$ curl localhost:8080/accounts/1
{"id":1,"document_number":"12345678922"}
```

Actually the port number is hardcoded inside `main.go`, if you need to run in other port,
edit the file and make the appropriated changes.
