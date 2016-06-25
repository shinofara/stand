DOCKER_PKG="github.com/shinofara/stand"

default: build-all

build-all: clean build-mac build-linux64

build-mac:
	@cd ./cmd/stand && \
	GOOS=darwin GOARCH=amd64 go build -o ../../stand_darwin_amd64 -ldflags="-w -s"

build-linux64:
	@cd ./cmd/stand && \
	GOOS=linux GOARCH=amd64 go build -o ../../stand_linux_amd64 -ldflags="-w -s"

clean:
	@rm -rf stand*

vet:
	@go vet $$(glide novendor)


ci-test:
	mkdir $GOPATH/src/github.com/shinofara
	ln -sf ../stand $GOPATH/src/github.com/shinofara/stand
	go test -v $(go list ./...|grep -v vendor) | go-junit-report set-exit-code=true > $CIRCLE_TEST_REPORTS/golang/junit.xml;
	go vet $(go list ./...|grep -v vendor)
