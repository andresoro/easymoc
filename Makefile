run: build
	./easymock

build: test build-ui build-server

build-server:
	go build .

build-ui:
	yarn --cwd client build

test:
	go test -v .