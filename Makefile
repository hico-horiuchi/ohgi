REVISION    := $(shell git log --abbrev-commit --pretty=oneline | head -1 | cut -d' ' -f1)
GO_BUILTOPT := -ldflags '-s -w'

run:
	go run main.go ${ARGS}

fmt:
	go fmt ./...

build: fmt
	sed -i'.bak' -e "s/__OHGI_REVISION__/$(REVISION)/g" main.go
	go build $(GO_BUILDOPT) -o bin/ohgi main.go
	mv main.go{.bak,}

clean:
	rm -f bin/ohgi

install: build
	cp bin/ohgi /usr/local/bin/

uninstall: clean
	rm -f /usr/local/bin/ohgi
