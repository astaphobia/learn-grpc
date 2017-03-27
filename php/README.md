#gRPC - PHP

## Server
PHP can't do gRPC server-y stuff :-( So run the go server instead.
```bash
$ go run client/main.go
```

## Client
Install dependencies:
```bash
$ cd client && composer install;
```

Run the client:
```bash
$ php index.php
```
