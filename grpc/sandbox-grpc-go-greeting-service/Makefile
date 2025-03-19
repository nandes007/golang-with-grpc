.PHONY: run-server
run-server:
		@go run cmd/server/main.go
 
.PHONY: run-client
NAME ?= "Anonymous"
run-client:
		@go run cmd/client/main.go --name=$(NAME)
