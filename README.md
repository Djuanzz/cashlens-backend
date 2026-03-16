# CashLens Backend

CashLens adalah backend service yang digunakan untuk menganalisis laporan transaksi bank dari file PDF dan mengubahnya menjadi data finansial yang mudah dianalisis.

Aplikasi ini dibuat menggunakan **Golang** dengan arsitektur REST API yang modular dan scalable.

Tujuan dari CashLens adalah membantu pengguna memahami pengeluaran mereka secara otomatis tanpa harus mencatat transaksi secara manual.

---

## Features

- Upload laporan transaksi bank dalam bentuk **PDF**
- Extract transaksi secara otomatis dari PDF
- Klasifikasi kategori pengeluaran (food, transport, bills, dll)
- Menyimpan transaksi ke database
- Menyediakan REST API untuk analisis finansial
- Statistik pengeluaran per kategori
- Summary pengeluaran bulanan

---

## Tech Stack

- **Golang** — backend language
- **Gin** — REST API framework
- **PostgreSQL** — database
- **GORM** — ORM
- **Docker** — containerization
- **JWT** — authentication
- **PDF Parser** — membaca laporan transaksi bank
