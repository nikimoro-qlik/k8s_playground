VERSION = dev
DOCKER_REPO = ghcr.io/nikimoro-qlik
BINARY_NAME = k8s_playground
DOCKER_IMAGE_NAME = $(BINARY_NAME)

BUILDDATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
REVISION := $(shell git rev-parse --short HEAD)

.PHONY: all
all: ghcr-push

.PHONY: build-docker
build-docker: build-local
	docker rmi -f $(DOCKER_REPO)/$(DOCKER_IMAGE_NAME):$(VERSION)
	docker build -t $(DOCKER_REPO)/$(DOCKER_IMAGE_NAME):$(VERSION) \
			--build-arg CREATED=$(BUILDDATE) \
			--build-arg REVISION=$(REVISION) \
			--build-arg VERSION=$(VERSION) \
			.

.PHONY: build-local
build-local: clean
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o $(BINARY_NAME) main.go

.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

.PHONY: push-ghcr
push-ghcr: build-docker
	docker push $(DOCKER_REPO)/$(DOCKER_IMAGE_NAME):$(VERSION)

.PHONY: build-run-local
build-run-local: build-docker
	kubectl delete -f resources/deployment.yaml --ignore-not-found=true
	kubectl apply -f resources/deployment.yaml

.PHONY: run-local
run-local:
	kubectl delete -f resources/deployment.yaml --ignore-not-found=true
	kubectl apply -f resources/deployment.yaml

