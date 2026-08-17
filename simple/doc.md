# Penjelasan Singkat: simple.go & injector.go

Folder `simple` ini berisi contoh implementasi **Dependency Injection (DI)** menggunakan **Google Wire**.

---

## 1. [`simple.go`](file:///Users/danimisme/Documents/Learn/belajar-golang-restful-api/simple/simple.go)
Mendefinisikan komponen dan relasi dependensinya:
- **[`SimpleRepository`](file:///Users/danimisme/Documents/Learn/belajar-golang-restful-api/simple/simple.go#L3)**: Komponen database/data access.
- **[`SimpleService`](file:///Users/danimisme/Documents/Learn/belajar-golang-restful-api/simple/simple.go#L7)**: Komponen bisnis logic yang membutuhkan `SimpleRepository`.
- **Constructor / Provider**:
  - [`NewSimpleRepository()`](file:///Users/danimisme/Documents/Learn/belajar-golang-restful-api/simple/simple.go#L11): Membuat instance `SimpleRepository`.
  - [`NewSimpleService(...)`](file:///Users/danimisme/Documents/Learn/belajar-golang-restful-api/simple/simple.go#L15): Membuat instance `SimpleService` sekaligus menyuntikkan (inject) `SimpleRepository`.

---

## 2. [`injector.go`](file:///Users/danimisme/Documents/Learn/belajar-golang-restful-api/simple/injector.go)
File konfigurasi bagi Google Wire untuk meng-generate objek secara otomatis:
- **Build Tag (`//go:build wireinject`)**: Memastikan file ini hanya diproses oleh tool `wire` dan diabaikan saat build aplikasi standar.
- **[`InitializeService()`](file:///Users/danimisme/Documents/Learn/belajar-golang-restful-api/simple/injector.go#L8)**: Fungsi inisialisasi. Memanggil `wire.Build(...)` dengan mendaftarkan semua constructor dari `simple.go`. Wire akan secara otomatis menentukan urutan pembuatan objek.

---

## 3. Cara Menjalankan
Untuk meng-generate file perakitan otomatis (`wire_gen.go`), jalankan perintah ini di terminal:
```bash
# Generate file wire_gen.go
wire gen belajar-golang-restful-api/simple
```

*(Catatan: Jika tool `wire` belum terinstal, jalankan `go install github.com/google/wire/cmd/wire@latest` terlebih dahulu).*
