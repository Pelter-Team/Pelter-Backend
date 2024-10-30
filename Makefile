
run: 
	@CGO_ENABLED=0 go build -o bin/app ./cmd/main.go
	@./bin/app


db:
	@docker compose up db -d