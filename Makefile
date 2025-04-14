OGEN=go run github.com/ogen-go/ogen/cmd/ogen@latest
SWAGGER=api/openapi/swagger.yaml
GEN_TARGET=internal/generated/api

generate:
	$(OGEN) --target $(GEN_TARGET) --clean $(SWAGGER)

start:
	docker-compose up --build
test:
	go test ./e2e/... -v -tags=integration
