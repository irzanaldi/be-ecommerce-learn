# be-ecommerce learn

> Project pembelajaran (learning): backend REST API e-commerce sederhana dengan Go (Gin + GORM + PostgreSQL + Redis).

## Cara Setup

Prasyarat:
- Go 1.25+ (`go version`)
- PostgreSQL berjalan (default `localhost:5432`)
- Redis berjalan (default `localhost:6379`)

Langkah:

1. Masuk ke folder project (perhatikan spasi pada nama folder):
   ```bash
   cd "be-ecommerce learn"
   ```

2. Siapkan konfigurasi. Aplikasi membaca dari environment variables (lihat `.env.example`); nilai default sudah aman untuk pengembangan lokal, jadi langkah ini opsional:
   ```bash
   cp .env.example .env
   # lalu export variabel yang diperlukan, contoh:
   export DB_USER=postgres DB_PASSWORD=postgres DB_NAME=ecommerce JWT_SECRET=ubah-ini
   ```

3. Unduh dependency:
   ```bash
   go mod download
   ```

4. Jalankan aplikasi (entrypoint ada di `cmd/api`):
   ```bash
   go run ./cmd/api
   ```
   API berjalan di `http://localhost:8080` (atau sesuai `PORT`).

5. Jalankan test:
   ```bash
   go test ./...
   ```

## Struktur

- `cmd/api/` — entrypoint aplikasi (`main.go`)
- `internal/` — kode inti: `config`, `handler`, `middleware`, `model`, `repository`, `service`, `router`, `cache`
- `pkg/` — paket reusable
- `migrations/` — skema/migrasi database

## Konfigurasi (environment variables)

| Variabel | Default | Keterangan |
|----------|---------|------------|
| `APP_ENV` | `development` | Mode aplikasi |
| `PORT` | `8080` | Port HTTP |
| `DB_HOST` / `DB_PORT` | `localhost` / `5432` | PostgreSQL |
| `DB_USER` / `DB_PASSWORD` / `DB_NAME` | `postgres` / `postgres` / `ecommerce` | Kredensial DB |
| `JWT_SECRET` | `secret` | Kunci penandatangan JWT (WAJIB diganti) |
| `REDIS_HOST` / `REDIS_PORT` | `localhost` / `6379` | Redis |
