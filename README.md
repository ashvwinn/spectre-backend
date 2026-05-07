
# Spectre
Spectre is in progress.

# Installation
1. Clone the repository:
```bash
git clone https://github.com/ashvwinn/spectre-backend.git
cd spectre-backend
```

2. Set environment variables (look at .env.example for reference)

3. Setup database with Docker Compose:
```bash
docker compose up -d
```

4. Run database migrations with the [`migrate`](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) tool:
```bash
migrate -path=./migrations -database=$SPECTRE_DB_DSN up
```

5. Start the application:
With `go` or with [`air`](https://github.com/air-verse/air)
```bash
air
# or
go run ./cmd/api
# or
go build -o bin/main ./cmd/api
./bin/main
```
