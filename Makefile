REVISION = $$(git log --abbrev-commit --pretty=oneline | head -1 | cut -d' ' -f1)

run:
	go run main.go ${ARGS}

fmt:
	go fmt ./...

build: fmt
	sed -i '' -e "s/__OHGI_REVISION__/$(REVISION)/g" main.go
	go build -o bin/ohgi main.go
	sed -i '' -e "s/$(REVISION)/__OHGI_REVISION__/g" main.go

clean:
	rm -f bin/ohgi

install: build
	cp bin/ohgi /usr/local/bin/

uninstall: clean
	rm -f /usr/local/bin/ohgi
