LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-note-api

generate-note-api:
	mkdir -p pkg/note_v1
	protoc --proto_path api/note_v1 \
	--go_out=pkg/note_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/note_v1/note.proto

build:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/grpc_server/main.go

copy-to-server:
	scp service_linux root@5.35.10.133:

# Variables for Docker registry and image information
REGISTRY_URL=cr.selcloud.ru
REPO=doker-hub/test-server
VERSION=v0.0.1
PLATFORM=linux/amd64
DOCKER_USER=token
DOCKER_PASS=CRgAAAAAxJ1YHa8GbBYlqg0wepU2VgmgVV9E2vuh

# Command to build and push the Docker image
docker-build-and-push:
	@echo "Building Docker image..."
	docker buildx build --platform $(PLATFORM) --no-cache -t $(REGISTRY_URL)/$(REPO):$(VERSION) .
	@echo "Logging in to Docker registry..."
	echo $(DOCKER_PASS) | docker login -u $(DOCKER_USER) --password-stdin $(REGISTRY_URL)
	@echo "Pushing Docker image..."
	docker push $(REGISTRY_URL)/$(REPO):$(VERSION)
	@echo "Docker image pushed: $(REGISTRY_URL)/$(REPO):$(VERSION)"

docker-pull:
	@echo "Logging in to Docker registry..."
	@echo $(DOCKER_PASS) | docker login -u $(DOCKER_USER) --password-stdin $(REGISTRY_URL)
	@echo "Pulling Docker image..."
	docker pull $(REGISTRY_URL)/$(REPO):$(VERSION)
	@echo "Docker image pulled: $(REGISTRY_URL)/$(REPO):$(VERSION)"
#docker run docker run cr.selcloud.ru/doker-hub/test-server:v0.0.1
#run with outside access:docker run -p 50051:50051 cr.selcloud.ru/doker-hub/test-server:v0.0.1



connect-to-server:
	ssh root@5.35.10.133

