# Set the Go compiler
GO := go

# Set the name of the executable
BINARY := service_auth

# Set the source file
SRC := cmd/main.go

# Set the flags
FLAGS := --connect_url="mongodb://mongousr:aweiplfjFWEJIFwiepfaojnrpwjvpjd3@195.80.238.9:27017"

# Define the build target
.PHONY: build
build:
	$(GO) build -o $(BINARY) $(SRC)

# Define the run target
.PHONY: run
run: build
	./$(BINARY) $(FLAGS)

# Define the clean target
.PHONY: clean
clean:
	@rm -f $(BINARY)

# Define the default target
.DEFAULT_GOAL := run