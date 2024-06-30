.PHONY: all clean

# Detect operating system
OS := $(shell uname)

all: quine.go quine1.go 
ifeq ($(OS),Darwin)
	@shasum -a 256 quine.go quine1.go
else
	@sha256sum quine.go quine1.go
endif 

quine.go: generator.go main.go
	@go run generator.go -- main.go > quine.go

quine1.go: quine.go 
	@go run quine.go > quine1.go

clean: 
	@rm -f quine.go quine1.go

