default: build-all

build-all: build-mac build-linux64

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o stand_darwin_amd64 -ldflags="-w -s"

build-linux64:
	GOOS=linux GOARCH=amd64 go build -o stand_linux_amd64 -ldflags="-w -s"

vet:
	@go vet $$(glide novendor)
