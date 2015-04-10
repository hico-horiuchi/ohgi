VERSION     := 0.1.6
GO_BUILDOPT := -ldflags '-s -w -X main.version $(VERSION)'

run:
	go run main.go ${ARGS}

fmt:
	go fmt ./...

build: fmt
	go build $(GO_BUILDOPT) -o bin/ohgi main.go

release: fmt
	GOOS=linux GOARCH=amd64 go build $(GO_BUILDOPT) -o bin/ohgi$(VERSION).linux-amd64 main.go
	GOOS=linux GOARCH=386 go build $(GO_BUILDOPT) -o bin/ohgi$(VERSION).linux-386 main.go
	GOOS=darwin GOARCH=amd64 go build $(GO_BUILDOPT) -o bin/ohgi$(VERSION).darwin-amd64 main.go
	GOOS=darwin GOARCH=386 go build $(GO_BUILDOPT) -o bin/ohgi$(VERSION).darwin-386 main.go

clean:
	rm -f bin/ohgi*

install: build
	cp bin/ohgi /usr/local/bin/

uninstall: clean
	rm -f /usr/local/bin/ohgi
