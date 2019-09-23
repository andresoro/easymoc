run: build build-ui
	./easymock

build:
	go build .

build-ui:
	yarn --cwd client build