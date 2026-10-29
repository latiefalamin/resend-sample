# Resend Sample (Go)

Proyek contoh sederhana untuk mengirim email menggunakan [Resend Go SDK](https://github.com/resend/resend-go) dan konfigurasi berbasis `.env` ([godotenv](https://github.com/joho/godotenv)).

## Prasyarat

- [Go](https://go.dev/doc/install) (versi 1.20 ke atas)
- Akun dan API Key dari [Resend](https://resend.com)

## Instalasi

1. Clone repositori ini atau buka direktori proyek:
   ```bash
   cd resend-sample
   ```

2. Unduh dan perbarui dependensi Go:
   ```bash
   go mod tidy
   ```

## Konfigurasi

1. Salin file template `.env.example` menjadi `.env`:
   ```bash
   cp .env.example .env
   ```

2. Buka file `.env` dan isi variabel konfigurasi:
   ```env
   # API Key dari Dashboard Resend
   RESEND_API_KEY=re_123456789

   # Alamat email penerima (pisahkan dengan koma jika lebih dari satu)
   RESEND_TO=penerima1@example.com, penerima2@example.com
   ```

   > **Catatan:** Aplikasi juga mendukung environment variable sistem jika file `.env` tidak tersedia, serta mendukung nama variabel alternatif `API_KEY` dan `TO_EMAIL`.

## Menjalankan Aplikasi

Jalankan program secara langsung menggunakan:

```bash
go run main.go
```

Atau lakukan build binary terlebih dahulu:

```bash
go build -o resend-sample main.go
./resend-sample
```

Jika email berhasil dikirim, terminal akan mencetak ID email seperti berikut:
```text
2026/10/29 10:45:00 email sent: 49a3999c-0ce1-4ea6-ab68-afcd6dc2e794
```

## Struktur Proyek

```text
resend-sample/
├── .env.example     # Template konfigurasi environment
├── .gitignore        # Daftar file/folder yang diabaikan oleh Git
├── README.md         # Dokumentasi proyek
├── go.mod            # Definisi modul Go dan dependensi
├── go.sum            # Checksum dependensi Go
└── main.go           # Kode utama pengiriman email
```
