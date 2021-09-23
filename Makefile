.all: generate-swagger-schema

help:
	@echo "generate-swagger-schema	: Generate OpenApi Swagger Definition Files"

generate-swagger-schema:
	swag init --parseDependency --output api/ -d ./cmd/