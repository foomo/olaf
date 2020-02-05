SHELL = /bin/zsh

TAG ?= master

SERVICE_VERSION?="${TAG}"
BUILD_HASH?=`git rev-parse --short HEAD || "unknown"`
BUILD_DATE?=`date -u +"%Y-%m-%dT%H:%M:%SZ"`
DOMAIN=localhost
GOPROXY?="https://athens.bestbytes.net,https://proxy.golang.org,direct"

DOCKER_BASE=docker-registry.bestbytes.net/globus/service-base:1.0.0

LDFLAGS=-ldflags="-w -X \"github.com/bestbytes/globus/cmd.BuildVersion=${SERVICE_VERSION}\" -X \"github.com/bestbytes/globus/cmd.BuildHash=${BUILD_HASH}\""

DOCKER_OLD_IMAGES=`docker images | grep "docker-registry.bestbytes.net/globus" | grep "<none>" | awk '{print $$3;}'`

GOPATH_MOUNT := $(if $(DOCKER_GOPATH),$(DOCKER_GOPATH),$(GOPATH))

#build:
#	GOPROXY=$(GOPROXY) CGO_ENABLED=0 go build $(LDFLAGS) -o /go/bin/service cmd/${SERVICE_NAME}/${SERVICE_NAME}.go

gotsrpc:
	gotsrpc gotsrpc.yaml