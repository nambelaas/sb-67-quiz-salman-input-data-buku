# API Buku Golang

Project ini adalah RESTful API untuk manajemen data buku, kategori, dan user, menggunakan Golang, Gin, dan PostgreSQL.

## Fitur

- Registrasi & login user (dengan JWT)
- CRUD kategori buku
- CRUD buku
- Menampilkan buku berdasarkan kategori

## Prasyarat

- Go 1.18+
- PostgreSQL
- Git

## Instalasi

1. **Clone repository**
   ```sh
   git clone https://github.com/nambelaas/sb-67-quiz-salman-input-data-buku.git
   cd nama-repo
   ```

2. **Copy dan edit konfigurasi**
   - Pastikan file konfigurasi database `config.json` sudah sesuai dengan environment kamu.

3. **Install dependency**
   ```sh
   go mod tidy
   ```

4. **Setup database**
   - Buat database PostgreSQL sesuai konfigurasi.
   - Jalankan migrasi:
     ```sh
     go run main.go
     ```
     (atau gunakan tools migrasi lain jika tersedia)

5. **Jalankan aplikasi**
   ```sh
   go run main.go
   ```
   Aplikasi akan berjalan di port default (misal: `localhost:8080`).

## Endpoint API
`{host}/swagger/index.html`

## Autentikasi

- Endpoint yang membutuhkan autentikasi harus mengirim header:
  ```
  Authorization: Bearer <JWT_TOKEN>
  ```

## Struktur Project

```
/database
/model
/router
/controller
/middleware
/structs
/helper
```



---

**Catatan:**  
- Pastikan environment variable/config sudah benar.
- Untuk pengujian, gunakan Postman/Insomnia.

---

Selamat mencoba!