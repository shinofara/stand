NAME=stand
REPO=github.com/shinofara/${NAME}
GO_VERSION=1.6.1

default: build

build: clean
	@cd cmd/$(NAME); \
	sh ../../bin/build.sh

build-on-docker: clean
	docker run --rm
		-w /go/src/$(REPO)/cmd/$(NAME) \
		-v ${PWD}:/go/src/$(REPO) \
	golang:$(GO_VERSION)-alpine \
	sh ../../bin/build.sh

clean:
	@rm -f stand*

vet:
	@go vet $$(glide novendor)
test:
	@go test $$(glide novendor)
lint:
	@for pkg in $$(go list ./... | grep -v /vendor/) ; do \
		golint $$pkg ; \
	done
