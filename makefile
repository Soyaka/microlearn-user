protoc:
	cd api/proto && protoc \
	--go_out=./gen \
	--go_opt=paths=source_relative \
	*.proto



protoco:
	cd api/proto && protoc \
    	    --go-grpc_out=./gen \
    	    --go-grpc_opt=paths=source_relative \
		    *.proto


progatway:
	cd api/proto && protoc \
	--grpc-gateway_out=./gen --grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	*.proto

all: build

build:
	@echo "Building..."
	
	@go build -o main cmd/main.go


run:
	@go run cmd/main.go

watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

commit:
	@read -p "Enter commit message: " message; \
	git add .; \
	git commit -m "$$message"; \
	git push origin master;


.PHONY: all build run test clean

docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi


testJs:
	cd api/proto &&	protoc --proto_path=. --js_out=import_style=es,binary:./js --grpc-web_out=import_style=es,mode=grpcwebtext:./js user.proto
