NAME=stand
REPO=github.com/shinofara/${NAME}
GO_VERSION=1.7

default: clean glide-install

build: clean
	@cd cmd/$(NAME); \
	sh ../../bin/build.sh

build-on-docker: clean-bin
	docker run --rm \
		-w /go/src/$(REPO)/cmd/$(NAME) \
		-v ${PWD}:/go/src/$(REPO) \
		golang:$(GO_VERSION) \
		sh ../../bin/build.sh

vet:
	@go vet $$(glide novendor)
test:
	@go test $$(glide novendor)
lint:
	@for pkg in $$(go list ./... | grep -v /vendor/) ; do \
		golint $$pkg ; \
	done

circleci-test-all: circleci-test circleci-vet
circleci-test:
	cd /home/ubuntu/.go_workspace/src/github.com/shinofara/stand && \
	go test -race -v $$(go list ./...|grep -v vendor) | go-junit-report set-exit-code=true > $(CIRCLE_TEST_REPORTS)/golang/junit.xml

circleci-vet:
	cd /home/ubuntu/.go_workspace/src/github.com/shinofara/stand && \
	go vet $$(go list ./...|grep -v vendor)

glide-install:
	docker run --rm \
	-v ${PWD}:/work \
	shinofara/docker-glide:0.12.2 install

glide-update:
	docker run --rm \
	-v ${PWD}:/work \
	shinofara/docker-glide:0.12.2 up

clean: clean-bin clean-vendor

clean-bin:
	@rm -f stand*

clean-vendor:
	@rm -rf vendor
