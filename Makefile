# Release infomation
VERSION := "0.0.2"
COMMIT := $(shell git rev-parse HEAD)
DATE := $(shell git --no-pager log -1 --format=%cd)

# Directory
BIN := ./_bin
LOGS := ./_logs

GO := /usr/local/go/bin/go
GOPATH := $(HOME)/go
GOROOT := /usr/local/go

DOCKER_SRC := ./docker

.PHONY: all build docker clean help

#######################################################################
# Build
#######################################################################
build: $(BIN)/pemgr-plat $(BIN)/pemgr-server $(BIN)/generate_cert $(BIN)/info.toml

$(BIN)/pemgr-plat:
	@$(GO) build -v -gcflags "-N -l" -o $(BIN)/pemgr-plat ./cmd/pemgr-plat

$(BIN)/pemgr-server:
	@$(GO) build -v -gcflags "-N -l" -o $(BIN)/pemgr-server ./cmd/pemgr-server

$(BIN)/generate_cert:
	cp $(GOROOT)/src/crypto/tls/generate_cert.go $(BIN)
	$(GO) build -o $@ $(BIN)/generate_cert.go

$(BIN)/info.toml:
	@echo "" > $(BIN)/info.toml
	@echo "[server]" > $(BIN)/info.toml
	@echo "version=\"$(VERSION)\"" >> $(BIN)/info.toml
	@echo "commit=\"$(COMMIT)\"" >> $(BIN)/info.toml
	@echo "date=\"$(DATE)\"" >> $(BIN)/info.toml

all: build

#######################################################################
# Docker
#######################################################################
docker:
	@docker build -f Dockerfile.release --build-arg MODE=0 . -t emgr

docker_debug:
	@docker build -f Dockerfile.debug --build-arg MODE=2 . -t emgr

docker_mockup:
	@docker build -f Dockerfile.release --build-arg MODE=2 . -t emgr

docker_build:
	@docker build -f Dockerfile.build --build-arg MODE=2 . -t emgr

docker_run:
	@docker run --name emgr --privileged --net=host --uts=host -v /var/run/redis:/var/run/redis -itd emgr

docker_runNet:
	@docker run --name emgr --privileged -p 8880:8880 -p 442:442 -v /var/run/redis:/var/run/redis -itd emgr

docker_runbash:
	@docker exec -it emgr bash

docker_runclean:
	@docker stop emgr;docker rm emgr

docker_clean:
	@docker stop emgr;docker rm emgr;docker rmi emgr


#######################################################################
# Common
#######################################################################
clean:
	@rm -rf $(BIN)
	@rm -rf $(LOGS)

help:
	@printf "Usage:\n"
	@printf "  make <target>\n\n"
	@printf ""
	@printf "Targets:\n"
	@printf "================================================\n"
	@printf "  plat\t\t Build pemgr-plat\n"
	@printf "  server\t Build pemgr-server\n"
	@printf "  all\t\t Build all\n"
	@printf "================================================\n"
	@printf "  docker\t Build docker image\n"
	@printf "  docker_mockup\t Build docker image using mockup API (debug)\n"
	@printf "  docker_clean\t Remove docker image\n"
	@printf "================================================\n"
	@printf "  docker_run\t\t Run container\n"
	@printf "  docker_runbash\t Enter container\n"
	@printf "  docker_runclean\t Remove container\n"
	@printf "================================================\n"
	@printf "  help\t\t Show help message.\n"
	@printf "================================================\n"

