# Proyek Absensi Karyawan

Proyek ini adalah aplikasi manajemen absensi karyawan yang dibangun menggunakan Go, framework Gin dan GORM library untuk kebutuhan technical test di Family Mart. Aplikasi ini memungkinkan pengguna untuk melihat list karyawan, mendaftarkan karyawan, login dan melakukan absensi.

## Fitur Utama

- List karyawan
- Mendaftarkan karyawan
- Autentikasi pengguna
- Mencatat absensi karyawan

## Prerequisites

Sebelum memulai, pastikan Anda telah menginstal:

- [Go](https://golang.org/dl/) (versi 1.16 atau lebih baru)
- [Docker](https://www.docker.com/get-started) (untuk menjalankan MySQL)
- [Docker Compose](https://docs.docker.com/compose/) (untuk mengelola kontainer)

## Instalasi

Ikuti langkah-langkah berikut untuk menginstal dan menjalankan aplikasi:

1. **Clone Repository**

   Pertama, clone repositori ini ke mesin lokal Anda:
   > 'git clone ..'

2. **Jalankan MySQL dengan Docker**

   Pastikan Anda memiliki file `docker-compose.yml` yang telah dikonfigurasi untuk menjalankan MySQL. Jalankan perintah berikut untuk memulai kontainer MySQL:
   > 'docker-compose up -d'

3. **Instalasi Dependensi Go**

   Setelah MySQL berjalan, instal dependensi Go dengan perintah berikut:
   > 'go mod tidy'

4. **Menjalankan Aplikasi**

   Setelah semua dependensi terinstal, Anda dapat menjalankan aplikasi dengan perintah:
   > 'go run main.go'


5. **Akses Aplikasi**

   Aplikasi akan berjalan di `http://localhost:8000`. Anda dapat menggunakan alat seperti Postman atau curl untuk menguji endpoint yang tersedia.

## Struktur Proyek

Berikut adalah struktur dasar proyek:

repo-name/ 
├── controllers/ # Logika bisnis dan pengendali 
├── routes/ # Definisi rute 
├── main.go # Titik masuk aplikasi 
├── docker-compose.yml # Konfigurasi Docker untuk MySQL 
├── init.sql # Skrip inisialisasi database 
└── README.md # Dokumentasi proyek