run: build
	./mocky

build:
	yarn --cwd client build
	go build .