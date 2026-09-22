# Waste-Track — Waste Bank Management System (CLI)

**Waste-Track** adalah aplikasi berbasis CLI (*Command Line Interface*) yang dibangun menggunakan bahasa pemrograman **Go (Golang)**. Aplikasi ini dirancang untuk membantu pengelolaan transaksi bank sampah, meliputi manajemen data warga, pengurutan, pencarian data berbasis algoritma, serta analisis statistik transaksi.

Proyek ini disusun sebagai Tugas Besar Mata Kuliah Algoritma dan Pemrograman 2 (Telkom University).

---

## 🚀 Fitur Utama

* **Manajemen Data Transaksi (CRUD)**:
  * **Tambah Data**: Pencatatan data warga (Nama, Jenis Sampah, Berat, Tanggal Transaksi) dilengkapi sistem *Auto-Generated ID*.
  * **Edit & Hapus Data**: Pembaruan atribut data serta penghapusan rekaman berdasarkan ID warga.
  * **Tampil Data**: Visualisasi daftar transaksi warga yang terdaftar.
* **Algoritma Pengurutan (Sorting / Ranking)**:
  * Peringkatan data warga berdasarkan berat sampah (*Ascending* & *Descending*).
  * Mengimplementasikan algoritma **Insertion Sort** dan **Selection Sort**.
* **Algoritma Pencarian (Searching)**:
  * **Binary Search**: Pencarian data cepat berdasarkan ID warga (dengan alur pra-pengurutan data).
  * **Sequential Search**: Pencarian data secara linier berdasarkan ID.
* **Analisis Statistik Data (Analytics)**:
  * Kalkulasi total akumulasi berat sampah (kg) dan total transaksi.
  * Pencarian nilai ekstremum (transaksi berat sampah tertinggi dan terendah beserta pemiliknya).
  * Rata-rata berat sampah per transaksi.

---

## 🛠️ Tech Stack & Konsep Algoritma

* **Bahasa Pemrograman**: Go (Golang)
* **Struktur Data**: `Array` dari `Struct` (`tabDataSampah`)
* **Algoritma**:
  * **Sorting**: Insertion Sort & Selection Sort
  * **Searching**: Binary Search & Sequential Search
* **Konsep Pemrograman**: Pass-by-pointer, Modular Functions & Procedures, Random ID Generation
