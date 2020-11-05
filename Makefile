.PHONY: build clean

GO=CGO_ENABLED=1 GO111MODULE=on go

APP_SERVICES=edgex/*

DOCKERS=docker_simple_filter_xml docker_secrets_example

GIT_SHA=$(shell git rev-parse HEAD)

.PHONY: build $(APP_SERVICES) $(DOCKERS)

build: $(APP_SERVICES)

$(APP_SERVICES):
	$(GO) build $(GOFLAGS) -o $@/app-service ./$@

clean:
	rm -f app-services/*/app-service

docker: $(DOCKERS)

d2c:
	docker build \
	    --build-arg http_proxy \
	    --build-arg https_proxy \
		-f edgex/Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t dragonlinux/d2c:$(GIT_SHA) \
		-t dragonlinux/d2c:dev \
		.
