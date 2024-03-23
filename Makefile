
server:
	go run cmd/api/main.go

test:
	go test ./pkg/service/ ./pkg/repository --cover

doc:
	swag fmt
	swag init -g ./cmd/api/main.go
	rm docs/docs.go docs/swagger.json
