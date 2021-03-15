SHELL                 = /bin/bash

APP_NAME              = pxydot
VERSION				  = 0.0.1

.PHONY: default
default: help

.PHONY: help
help:
	@echo 'Management commands for ${APP_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make all                                equivalent of "make build" and "make package"'
	@echo '    make build                              build binary file inside ./bin directory'
	@echo '    make package                            build docker image'
	@echo '    make run                                run the app with default configuration'
	@echo '    make docker-run                         run an interactive containerized app with default configuration'
	@echo '    make clean                              remove the binary file and the docker image'
	@echo

.PHONY: build
build:
	@echo "Building ${APP_NAME}"
	go build -o bin/${APP_NAME}

.PHONY: package
package:
	docker build --tag ${APP_NAME}:${VERSION} .

.PHONY: all
all: build package

.PHONY: run
run:
	@echo "Running ${APP_NAME} ${VERSION} with default configuration"
	./bin/${APP_NAME}

.PHONY: docker-run
docker-run:
	@echo "Running ${APP_NAME} ${VERSION} with default configuration on Docker"
	docker run -it --rm -p 127.0.0.1:53:53/UDP -p 127.0.0.1:53:53 ${APP_NAME}:${VERSION}

.PHONY: clean
clean:
	@echo "Removing ${APP_NAME} ${VERSION}"
	@test ! -e bin/${APP_NAME} || rm bin/${APP_NAME} 
	@test ! -e bin/ || rm -rf bin/ && \
	docker rmi ${APP_NAME}:${VERSION}