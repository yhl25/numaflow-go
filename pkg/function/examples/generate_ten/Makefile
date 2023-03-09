.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./dist/generate-ten-example main.go

.PHONY: image
image: build
	docker build -t "quay.io/chromevoid/map-generate-ten:latest" --target generate-ten .

clean:
	-rm -rf ./dist
