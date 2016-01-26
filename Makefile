MAKEFLAGS += --warn-undefined-variables
IMAGE ?= carsonoid/ec2_tag_fetcher
COMPRESS_BINARY ?= true

BUILDER_IMAGE ?= centurylink/golang-builder
BUILDER_IMAGE_EXTRA-build-cross = -cross
BUILDER_IMAGE_EXTRA-docker-build =

.PHONY: %

all: build
local: build-local

update-deps:
	go get -v -u -f github.com/aws/aws-sdk-go/...

godep:
	godep save

build-local:
	go generate ./...
	go build

docker-build:
	docker run --rm -v $(PWD):/src -e COMPRESS_BINARY=$(COMPRESS_BINARY) $(BUILDER_IMAGE)$(BUILDER_IMAGE_EXTRA-$@)

build build-cross: update-deps godepdocker-build

