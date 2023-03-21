.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./dist/generate-100-example main.go

.PHONY: image
image: build
	docker build -t "quay.io/chromevoid/map-generate-100:latest" --target generate-100 .

clean:
	-rm -rf ./dist
