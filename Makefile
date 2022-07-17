run:
	go run cmd/app/main.go

.PHONY: pub sub

default: pub sub

pub:
	go build -o pub cmd/pub/pub.go
	go build -o sub cmd/sub/sub.go