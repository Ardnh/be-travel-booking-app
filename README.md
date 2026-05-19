# Travel Booking App Backend

Backend untuk aplikasi pemesanan layanan travel/transportasi. Aplikasi ini mengelola data vendor travel, titik pool/keberangkatan, tipe layanan, layout kursi kendaraan, jadwal perjalanan, user, autentikasi, dan otorisasi.

Dibangun dengan **Go**, **Fiber**, **GORM**, **PostgreSQL**, **Redis**, **JWT**, dan **Casbin** menggunakan pendekatan **Clean Architecture** serta **Repository Pattern**.

## Gambaran Arsitektur

```text
Interface (HTTP Handlers & Routes)
        |
        v
Application (Use Case Services & DTO)
        |
        v
Domain (Entities, Repository Contracts, Business Rules)
        |
        v
Infrastructure (PostgreSQL, Redis, Repository Implementations)
```

## Fitur Utama

- Manajemen tipe layanan travel, termasuk kebutuhan kursi, alamat pickup, dan alamat dropoff.
- Manajemen pool point atau titik keberangkatan/kedatangan vendor.
- Struktur data vendor travel dan user vendor.
- Struktur jadwal perjalanan dengan asal, tujuan, tanggal, jam berangkat, harga per kursi, dan ketersediaan kursi.
- Layout kendaraan dan posisi kursi.
- Autentikasi berbasis JWT.
- Otorisasi berbasis Casbin.
- Caching menggunakan Redis.
- Database PostgreSQL dengan GORM.
- REST API menggunakan Fiber.
- Docker Compose untuk PostgreSQL dan Redis.

## Endpoint Aktif

Base path API:

```text
/api/v1
```

Endpoint yang saat ini terdaftar:

```text
GET    /service-types
GET    /service-types/:id
POST   /service-types
PUT    /service-types/:id
DELETE /service-types/:id

GET    /pool-points
GET    /pool-points/:id
GET    /vendors/:vendorId/pool-points
POST   /pool-points
PUT    /pool-points/:id
DELETE /pool-points/:id
```

## Teknologi

- Go 1.25
- Fiber v3
- GORM
- PostgreSQL
- Redis
- JWT
- Casbin
- Logrus
- Validator
- Docker Compose

## Quick Start

1. Install dependency:

   ```bash
   make deps
   ```

2. Jalankan PostgreSQL dan Redis:

   ```bash
   make docker-up
   ```

3. Jalankan migrasi database:

   ```bash
   go run cmd/migrate/main.go
   ```

   Atau jalankan migrasi sekaligus seeder:

   ```bash
   go run cmd/migrate/main.go -seed
   ```

4. Jalankan server:

   ```bash
   make run
   ```

Server berjalan pada port default:

```text
http://localhost:8080
```

## Konfigurasi Environment

Aplikasi membaca konfigurasi dari environment variable. Jika tidak diisi, nilai default berikut akan digunakan:

```env
APP_ENV=development
APP_PORT=8080
APP_JWT_SECRET=

DB_HOST=localhost
DB_PORT=5432
DB_USER=root
DB_PASSWORD=root
DB_NAME=travel_booking_app_db
DB_SSLMODE=disable

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
```

## Struktur Direktori

```text
cmd/                  Entry point aplikasi dan command migrasi
internal/config/      Konfigurasi aplikasi
internal/domain/      Entity, kontrak repository, dan domain service
internal/application/ DTO, mapper, dan implementasi use case service
internal/infrastructure/
                      Implementasi database, cache, migration, seeder, repository
internal/interfaces/  HTTP handler, route, middleware, dan response
internal/utils/       Helper untuk JWT, logger, validator, Casbin
api/                  Dokumentasi atau spesifikasi API
docs/                 Dokumentasi tambahan
scripts/              Script pendukung
tests/                Area pengujian
```

## Perintah Makefile

```bash
make build          # Build binary server
make run            # Jalankan server
make deps           # Download dan rapikan dependency
make test           # Jalankan test
make test-coverage  # Jalankan test dengan coverage report
make fmt            # Format kode Go
make lint           # Jalankan golangci-lint
make docker-up      # Jalankan service Docker Compose
make docker-down    # Matikan service Docker Compose
```

## Catatan Pengembangan

Beberapa entity domain seperti vendor, jadwal, layout kursi, user, dan autentikasi sudah tersedia di kode. Route yang aktif di `cmd/server/main.go` saat ini baru me-wire service type dan pool point, sehingga endpoint lain dapat ditambahkan bertahap mengikuti pola handler, service, repository, dan route yang sudah ada.
