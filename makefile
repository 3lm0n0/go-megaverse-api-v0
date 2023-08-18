##################################################
# Usage:
##################################################
# make          # compiles all binary.
# make init     # creates microservice module.
# make setup    # sets up the microservice.
# make build    # builds the microservice.
# make run      # runs the microservice.
# make clean    # removes ALL binaries and objects.

BINARY_NAME=megaverse

.PHONY:= hello setup build run test
.DEFAULT_GOAL:= setup build run

init:
	@echo "=> Go module fact initializing"
	@go mod init ${BINARY_NAME}

setup:
	@echo "=> Stetting microservice"
	@export GOSUMDB=off
	@go mod tidy
	@go mod download
	@echo "=> Setup completed"

build:
	@echo "=> Building microservice"
	@go build -o ./bin/${BINARY_NAME}
	@echo "=> Microservice built"
	
run:
	./bin/${BINARY_NAME}

clean:
	@echo "Cleaning up all binaries, objects and sum ..."
	@go clean
	@rm -rvf *.o ./bin/${BINARY_NAME} go.sum