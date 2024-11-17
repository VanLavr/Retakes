gen:
	@export PATH="$PATH:$HOME/bin:$HOME/go/bin"
	@oapi-codegen -generate gin -package gen -o api/gen/gen.go ./api/swagger/swagger.yaml

run:
	@go run ./cmd/main.go

test:
	@echo 'test'