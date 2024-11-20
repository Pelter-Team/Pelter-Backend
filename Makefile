
run: 
	@CGO_ENABLED=0 go build -o bin/app.exe ./cmd/main.go
	@./bin/app.exe


db:
	@docker compose up db -d