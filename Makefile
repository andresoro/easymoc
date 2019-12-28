run: build
	./easymock

build: build-ui build-server

build-server:
	go build .

build-ui:
	yarn --cwd client build