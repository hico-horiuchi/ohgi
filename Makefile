VERSION     := 0.4.0
GO_BUILDOPT := -ldflags '-s -w -X main.version=$(VERSION)'

gom:
	go get github.com/mattn/gom
	gom install

run:
	gom run main.go ${ARGS}

fmt:
	gom exec goimports -w *.go sensu/*.go ohgi/*.go

build: fmt
	gom build $(GO_BUILDOPT) -o bin/ohgi main.go

release: fmt
	GOOS=linux GOARCH=amd64 gom build $(GO_BUILDOPT) -o bin/ohgi$(VERSION).linux-amd64 main.go
	GOOS=linux GOARCH=386 gom build $(GO_BUILDOPT) -o bin/ohgi$(VERSION).linux-386 main.go
	GOOS=darwin GOARCH=amd64 gom build $(GO_BUILDOPT) -o bin/ohgi$(VERSION).darwin-amd64 main.go
	GOOS=darwin GOARCH=386 gom build $(GO_BUILDOPT) -o bin/ohgi$(VERSION).darwin-386 main.go

clean:
	rm -f bin/ohgi*

install: build
	cp bin/ohgi /usr/local/bin/

uninstall: clean
	rm -f /usr/local/bin/ohgi
