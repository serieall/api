.PHONY: all fmt install vendor deb

fmt:
	go fmt api/*.go
	go fmt api/bootstrap/*.go
	go fmt api/controllers/*.go
	go fmt api/models/*.go

vendor:
	glide update
	rm -rf vendor/github.com/nats-io/nkeys
	go get -u golang.org/x/crypto/ed25519

install: fmt vendor
	mkdir -p ./bin
	GOOS=linux go build -o ./bin/serieall-api ./api/
	chmod +x ./bin/serieall-api
