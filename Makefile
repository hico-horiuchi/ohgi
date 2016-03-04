VERSION     := 0.5.2
GO_BUILDOPT := -ldflags '-s -w -X main.version=$(VERSION)'

gom:
	go get github.com/mattn/gom
	gom install

run:
	gom run main.go ${ARGS}

fmt:
	gom exec goimports -w *.go ohgi/*.go

build: fmt
	gom build $(GO_BUILDOPT) -o bin/ohgi main.go

release: fmt
	if [ -e bin/v$(VERSION) ]; then rm -rf bin/v$(VERSION); fi
	mkdir bin/v$(VERSION)
	GOOS=linux GOARCH=amd64 gom build $(GO_BUILDOPT) -o bin/v$(VERSION)/ohgi$(VERSION).linux-amd64 main.go
	GOOS=linux GOARCH=386 gom build $(GO_BUILDOPT) -o bin/v$(VERSION)/ohgi$(VERSION).linux-386 main.go
	GOOS=darwin GOARCH=amd64 gom build $(GO_BUILDOPT) -o bin/v$(VERSION)/ohgi$(VERSION).darwin-amd64 main.go
	GOOS=darwin GOARCH=386 gom build $(GO_BUILDOPT) -o bin/v$(VERSION)/ohgi$(VERSION).darwin-386 main.go
	zip bin/v$(VERSION)/ohgi$(VERSION).linux-amd64.zip bin/v$(VERSION)/ohgi$(VERSION).linux-amd64
	zip bin/v$(VERSION)/ohgi$(VERSION).linux-386.zip bin/v$(VERSION)/ohgi$(VERSION).linux-386
	zip bin/v$(VERSION)/ohgi$(VERSION).darwin-amd64.zip bin/v$(VERSION)/ohgi$(VERSION).darwin-amd64
	zip bin/v$(VERSION)/ohgi$(VERSION).darwin-386.zip bin/v$(VERSION)/ohgi$(VERSION).darwin-386
	ls -d bin/v$(VERSION)/* | grep -v zip | xargs rm

clean:
	rm -rf bin/*

link:
	mkdir -p $(GOPATH)/src/github.com/hico-horiuchi
	ln -s $(CURDIR) $(GOPATH)/src/github.com/hico-horiuchi/ohgi

unlink:
	rm $(GOPATH)/src/github.com/hico-horiuchi/ohgi
	rmdir $(GOPATH)/src/github.com/hico-horiuchi

install: build
	cp bin/ohgi /usr/local/bin/

uninstall: clean unlink
	rm -f /usr/local/bin/ohgi
