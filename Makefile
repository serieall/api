.PHONY: all fmt install vendor deb

fmt:
	go fmt api/*.go
#	go fmt api/bootstrap/*.go
	go fmt api/controllers/*.go
	go fmt api/models/*.go

vendor:
	glide update

install: fmt vendor
	mkdir -p ./bin
	GOOS=linux go build -o ./bin/serieall-api ./api/
	chmod +x ./bin/serieall-api
