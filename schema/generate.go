package schema

//go:generate docker compose down
//go:generate docker compose up pg-migrate --build
//go:generate sh -c "docker compose exec -T pg pg_dump -U postgres postgres --schema-only --no-owner --no-tablespaces > schema.sql"
//go:generate sqlc generate --file sqlc.yaml
