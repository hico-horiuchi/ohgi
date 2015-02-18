VERSION     := 0.1.0
GO_BUILDOPT := -ldflags '-s -w -X main.version $(VERSION)'

run:
	go run main.go ${ARGS}

fmt:
	go fmt ./...

build: fmt
	go build $(GO_BUILDOPT) -o bin/ohgi main.go

clean:
	rm -f bin/ohgi

install: build
	cp bin/ohgi /usr/local/bin/

uninstall: clean
	rm -f /usr/local/bin/ohgi
