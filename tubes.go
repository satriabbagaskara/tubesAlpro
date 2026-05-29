package main

import (
	"fmt"
)

const maxPengguna = 1000
type dataSampah struct {
	id int
	nama string
	beratSampah int
	jenisSampah string
	tanggal string
}
type tabDataSampah [maxPengguna]dataSampah

// PROGRAM UTAMA UNTUK MENJALANKAN APLIKASI //
func main(){
	var user tabDataSampah
	halamanUtama(&user)
}

// INTERFACE UTAMA //
func halamanUtama(user *tabDataSampah) {
 	fmt.Println("...")
 	fmt.Println("=============================== [ SELAMAT DATANG ] ==============================")
 	fmt.Println("                                  App Waste-Track              ")
 	fmt.Println("                           Algoritma pemrograman 2 - 2026       ")
 	fmt.Println("                                  Satria X Zikri               ")
 	fmt.Println("----------------------------------------------------------------------------------")
 	fmt.Println("Data aplikasi masih kosong, dimohon untuk mengisi data warga terlebih dahulu...")

 	inputData(user)
}

func halamanUtama_kedua(user *tabDataSampah, banyakOrang *int) {
	var inputPilihan int

 	fmt.Println("...")
 	fmt.Println("=============================== [ SELAMAT DATANG ]==============================")
 	fmt.Println("                                  App Waste-Track              ")
 	fmt.Println("                           Algoritma pemrograman 2 - 2026       ")
 	fmt.Println("                                  Satria X Zikri               ")
 	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Println(" ")
	
	fmt.Println("1. Isi data Sampah")
	fmt.Println("2. Edit data Sampah")
	fmt.Println("3. Hapus data Sampah")
	fmt.Println("4. Lihat Statistik")
	fmt.Println("5. Tampilkan data Sampah")
	fmt.Println("6. Keluar App")
	fmt.Println("7. Cari data warga")
	fmt.Println("-------------------------->")
	
	fmt.Print("Masukkan pilihan kamu yang mau kami jalankan: ")
	fmt.Scan(&inputPilihan)
	fmt.Println("----------------------------------------------------------------------------------")

	if inputPilihan == 1 {
		inputDataTambahan(user, banyakOrang)
	} else if inputPilihan == 2{
		editData(user, banyakOrang)
	} else if inputPilihan == 3{
		hapusData(user, banyakOrang)
	} else if inputPilihan == 4{
		statistikSampah(user, banyakOrang)
	} else if inputPilihan == 5{
		tampilData(user, banyakOrang)
	} else if inputPilihan == 6{
		fmt.Println("Terimakasih telah menggunakan Waste-Track")
	} else if inputPilihan == 7{
		cariDataWarga(user, banyakOrang)
	} else {
		fmt.Println("Fitur belum tersedia atau input tidak valid")
		halamanUtama_kedua(user, banyakOrang)
	}
}

// --- Procedurre Input Data --- //
func inputData(user *tabDataSampah) {
	var i, banyakOrang int
	
	fmt.Println("\n...")
	fmt.Print("Berapa orang yang setor sampah: \n")
	fmt.Scan(&banyakOrang)

	fmt.Println("<>-------------------------------------<<>>-------------------------------------<>")
	fmt.Println("                            ---- ATURAN PENGISIAN ----                            ")
	fmt.Println("<>------------------------------------------------------------------------------<>")
	fmt.Println("1. Jenis sampah diisi antara Organik dan Anorganik")
	fmt.Println("2. Satuan berat sampah adalah Kg")
	fmt.Println("3. Template tanggal pengiriman (DD-MM-YY)")
	fmt.Println("__________________________________________________________________________________")
	
	for i = 1; i <= banyakOrang; i++ {
		fmt.Print("ID: ")
		fmt.Scan(&user[i].id)
		
		fmt.Print("Nama: ")
		fmt.Scan(&user[i].nama)
		
		fmt.Print("Jenis sampah: ")
		fmt.Scan(&user[i].jenisSampah)
		
		fmt.Print("Berat sampah: ")
		fmt.Scan(&user[i].beratSampah)
		
		fmt.Print("Tanggal pengiriman: ")
		fmt.Scan(&user[i].tanggal)
		
		fmt.Println("")
	}
	fmt.Println("..................................................................................")
	fmt.Println("======================================== <<>> ====================================")
	fmt.Println("")
	halamanUtama_kedua(user, &banyakOrang)
}

func inputDataTambahan(user *tabDataSampah, banyakOrang *int) {
	var banyakOrangBaru, i, totalOrangSementara int
	fmt.Print("Banyak orang setor sampah: ")
	fmt.Scan(&banyakOrangBaru)
	totalOrangSementara = *banyakOrang + banyakOrangBaru
	
	fmt.Println("<>-------------------------------------<<>>-------------------------------------<>")
	fmt.Println("                            ---- ATURAN PENGISIAN ----    ")
	fmt.Println("<>------------------------------------------------------------------------------<>")
	fmt.Println("1. Setiap nama diawali huruf kapital")
	fmt.Println("2. Jenis sampah diisi antara Organik dan Anorganik")
	fmt.Println("3. Satuan berat sampah adalah Kg")
	fmt.Println("4. Template tanggal pengiriman (DD-MM-YY)")
	fmt.Println("__________________________________________________________________________________")


	for i = *banyakOrang+1; i <= totalOrangSementara; i++ {
		fmt.Print("ID: ")
		fmt.Scan(&user[i].id)
		fmt.Print("Nama: ")
		fmt.Scan(&user[i].nama)
		fmt.Print("Jenis sampah: ")
		fmt.Scan(&user[i].jenisSampah)
		fmt.Print("Berat sampah: ")
		fmt.Scan(&user[i].beratSampah)
		fmt.Print("Tanggal pengiriman: ")
		fmt.Scan(&user[i].tanggal)
		fmt.Println("")
	}

	*banyakOrang = totalOrangSementara
	halamanUtama_kedua(user, banyakOrang)
}

// --- PROCEDURE EDIT DATA --- // 
func editData(user *tabDataSampah, banyakOrang *int) {
	var gantiID int
	var i, inputPilihan, IDpilihan int
	var keberadaan bool
	
	if *banyakOrang == 0{
		fmt.Println(" ")
		fmt.Println("Maaf, anda tidak bisa edit data karena tidak ada data yang terdaftar")
		fmt.Println("Silahkan isi data terlebih dahulu...")
	} else {
		keberadaan = false
		fmt.Println("\n...")
		fmt.Println("Data apa yang mau diubah: ")
		fmt.Println("1. ID")
		fmt.Println("2. Nama")
		fmt.Println("3. Jenis Sampah")
		fmt.Println("4. Berat Sampah")
		fmt.Println("5. Tanggal Pengiriman")
		fmt.Println("----------------------------------------------------------------------------------")

		fmt.Print("Pilihan: ")
		fmt.Scan(&inputPilihan)
		fmt.Println("----------------------------------------------------------------------------------")
		
		// Ganti ID //
		if inputPilihan == 1 {
			fmt.Print("ID yang mau diganti: ")
			fmt.Scan(&IDpilihan)
			
			fmt.Print("ID baru: ")
			fmt.Scan(&gantiID)
			
			for i = 1; i <= *banyakOrang; i++ {
				if user[i].id == IDpilihan {
					fmt.Println("ID ditemukan!")
					user[i].id = gantiID
					fmt.Printf("ID sekarang: %d\n\n", user[i].id)
					keberadaan = true
				}
			}
			if !keberadaan {
				fmt.Println("Data tidak dapat ditemukan")
			}
			
		}
		
		// Ganti Nama //
		var namaBaru, namaPilihan string
		if inputPilihan == 2 {
			fmt.Print("Nama yang mau diganti: ")
			fmt.Scan(&namaPilihan)
			fmt.Print("Nama baru: ")
			fmt.Scan(&namaBaru)
			
			for i = 1; i <= *banyakOrang; i++ {
				if user[i].nama == namaPilihan {
					fmt.Println("Nama ditemukan!")
					user[i].nama = namaBaru
					fmt.Printf("Nama sekarang: %s\n", user[i].nama)
					keberadaan = true
				}
			}
			if keberadaan == false {
				fmt.Println("Nama tidak dapat ditemukan")
			}
		}
		
		// Ganti jenis sampah //
		var jenisSampahBaru string
		var IDwarga_gantiJenis int
		if inputPilihan == 3 {
			fmt.Print("ID warga: ")
			fmt.Scan(&IDwarga_gantiJenis)
			fmt.Print("Jenis baru: ")
			fmt.Scan(&jenisSampahBaru)
			
			for i = 1; i <= *banyakOrang; i++ {
				if IDwarga_gantiJenis == user[i].id {
					fmt.Println("ID dapat ditemukan!")
					user[i].jenisSampah = jenisSampahBaru
					fmt.Printf("jenis sampah diperbarui menjadi: %s\n", user[i].jenisSampah)
					keberadaan = true
				}
			}
			
			if !keberadaan {
				fmt.Println("ID tidak dapat ditemukan")
			}
		}
		
		// Edit berat sampah // 
		var IDwarga_gantiberat, beratBaru int
		if inputPilihan == 4 {
			fmt.Print("ID warga: ")
			fmt.Scan(&IDwarga_gantiberat)
			fmt.Print("Berat baru: ")
			fmt.Scan(&beratBaru)
			
			for i = 1; i <= *banyakOrang; i++ {
				if IDwarga_gantiberat == user[i].id {
					fmt.Println("ID dapat ditemukan!")
					user[i].beratSampah = beratBaru
					fmt.Printf("Berat diperbarui menjadi: %d\n", user[i].beratSampah)
					keberadaan = true
				}
			}
			
			if !keberadaan {
				fmt.Println("ID tidak dapat ditemukan")
			}
		}
		
		// Ganti tanggal pengiriman //
		var IDwarga_gantiTanggal int
		var tanggalBaru string
		
		if inputPilihan == 5 {
			fmt.Print("ID warga: ")
			fmt.Scan(&IDwarga_gantiTanggal)
			
			fmt.Print("Tanggal baru: ")
			fmt.Scan(&tanggalBaru)
			for i = 1; i <= *banyakOrang; i++ {
				if user[i].id == IDwarga_gantiTanggal {
					fmt.Println("ID dapat ditemukan!")
					user[i].tanggal = tanggalBaru
					fmt.Printf("Tanggal diperbarui menjadi %s\n", user[i].tanggal)
					keberadaan = true
				}
			}
			
			if !keberadaan {
				fmt.Println("ID tidak dapat ditemukan")
			}
		}
		
		if inputPilihan > 5 {
			fmt.Println("Inputan tidak tersedia")
			editData(user, banyakOrang)
		}
	}
	
	fmt.Println("..................................................................................")
	fmt.Println("======================================== <<>> ====================================")
	halamanUtama_kedua(user, banyakOrang)
}

// Procedure tampikan data warga (hybrid)
func tampilData(user *tabDataSampah, banyakOrang *int) {
	var i int
	
 	fmt.Println("\n=============================[ DATA SAMPAH WARGA ]===============================")
	if *banyakOrang == 0 {
		fmt.Println(" ")
		fmt.Println("Maaf, Belum ada data sampah. Silahkan input datanya terlebih dahulu...")
	} else {
		fmt.Println(" ")
		fmt.Println("___________________________________________________________________________")
		fmt.Println("| No. |   ID   |      Nama     |     Jenis    | Berat (kg) |    Tanggal   |")
		fmt.Println("+-------------------------------------------------------------------------+")
		for i = 1; i <= *banyakOrang; i++ {

			fmt.Printf("| %-3d | %-6d | %-13s | %-12s | %-10d | %-12s |\n", i, user[i].id, user[i].nama, user[i].jenisSampah, user[i].beratSampah, user[i].tanggal)
			fmt.Println("+-------------------------------------------------------------------------+")
		}
	}

	fmt.Println("\n\n...............................................................................")
	fmt.Println("======================================== <<>> =====================================")
	halamanUtama_kedua(user, banyakOrang)
}

// Procedure Hapus data warga (Satriaa)
func hapusData(user *tabDataSampah, banyakOrang *int) {
	var hapusWarga, i int
	var keberadaan int
	
	if *banyakOrang == 0{
		fmt.Println(" ")
		fmt.Println("Tidak ada data warga yang terdaftar")
		fmt.Println("Silahkan isi data transaksi terlebih dahulu")
	} else {
		keberadaan = -1
		fmt.Println(" ")
		fmt.Print("Masukkan ID warga yang mau di hapus: ")
		fmt.Scan(&hapusWarga)

		for i = 1; i <= *banyakOrang && keberadaan == -1; i++ {
			if user[i].id == hapusWarga{
				keberadaan = i
			}
		}
		if keberadaan == -1 {
			fmt.Println("ID warga tidak dapat ditemukan")
		} else {
			fmt.Printf("ID ditemukan atas nama %s\n", user[keberadaan].nama)

			for i = keberadaan; i < *banyakOrang; i++{
				user[i] =  user[i+1]
			}

			user[*banyakOrang] = dataSampah{}
			*banyakOrang= *banyakOrang - 1
			fmt.Println("Data berhasil dihapus")
	}
	
	fmt.Println("..................................................................................")
	fmt.Println("======================================== <<>> ====================================")
	halamanUtama_kedua(user, banyakOrang)
	}
}

// --- Procedure mencari data warga --- //
// ---> MENGGUNAKAN BINARY SEARCH <--- //
func cariDataWarga(user *tabDataSampah, banyakOrang *int) {
	var cariIDWarga, mid, kiri, kanan int
	var keberadaan bool
	keberadaan = false
	
	if *banyakOrang == 0{
		fmt.Println(" ")
		fmt.Println("Data tidak ada. Silahkan isi data terlebih dahulu...")
	} else {
		fmt.Print("Masukkan ID warga yang mau dicari: ")
		fmt.Scan(&cariIDWarga)
		
		kanan = *banyakOrang
		mid = (kanan + kiri) / 2
		
		// --- look dari kiri ke mid --- //
		for kiri = 1; kiri <= mid; kiri++ {
			if user[kiri].id == cariIDWarga {
				keberadaan = true
				fmt.Println("\nData ditemukan!")
				fmt.Println("-------------------------------- <<>> ------------------------------------")
				fmt.Printf("Nama: %s\nID: %d\nBerat sampah: %d kg\n", user[kiri].nama, user[kiri].id, user[kiri].beratSampah)
				fmt.Printf("Jenis sampah: %s\nTanggal pengiriman: %s\n", user[kiri].jenisSampah, user[kiri].tanggal)
				fmt.Println("-------------------------------- <<>> ------------------------------------")
			}
		}
		// --- look dari mid ke kanan --- //
		for kiri = mid+1; kiri <= kanan; kiri++ {
			if user[kiri].id == cariIDWarga {
				keberadaan = true
				fmt.Println("\nData ditemukan!")
				fmt.Println("-------------------------------- <<>> ------------------------------------")
				fmt.Printf("Nama: %s\nID: %d\nBerat sampah: %d kg\n", user[kiri].nama, user[kiri].id, user[kiri].beratSampah)
				fmt.Printf("Jenis sampah: %s\nTanggal pengiriman: %s\n", user[kiri].jenisSampah, user[kiri].tanggal)
				fmt.Println("-------------------------------- <<>> ------------------------------------")
			}
		}
		if !keberadaan {
			fmt.Println("ID tidak dapat ditemukan")
		}
	}
	fmt.Println("..................................................................................")
	fmt.Println("======================================== <<>> ====================================")
	halamanUtama_kedua(user, banyakOrang)
}

// Procedure Statistik Data Sampah Mingguan (Satria)
func statistikSampah(user *tabDataSampah, banyakOrang *int) {
	fmt.Println("<<------------------------- [ STATISTIK DATA MINGGUAN ] ------------------------>>")
	var i, total int
	var rata float64
	var idxMax, idxMin int

	if *banyakOrang == 0{
		fmt.Println(" ")
		fmt.Println("Belum ada transaksi sampah.")
		fmt.Println("Silahkan isi data transaksi terlebih dahulu.")
	} else {
		// total sampah
		total = 0
		for i = 1; i <= *banyakOrang; i++{
			total = total + user[i].beratSampah
		}
		fmt.Printf("Total seluruh sampah: %d kg\n", total)

		// jumlah transaksi
		fmt.Printf("Jumlah transaksi: %d\n", *banyakOrang)

		// sampah tertinggi & terendah
		idxMax = 1
		idxMin = 1
		for i = 2; i <= *banyakOrang; i++{
			if user[i].beratSampah > user[idxMax].beratSampah{
				idxMax = i
			}
			if user[i].beratSampah < user[idxMax].beratSampah{
				idxMin = i
			}
			fmt.Printf("Sampah terbanyak: %d kg, atas nama %s\n", user[idxMax].beratSampah, user[idxMax].nama)
			fmt.Printf("Sampah terendah: %d kg, atas nama %s\n", user[idxMin].beratSampah, user[idxMin].nama)
		}

		// rata-rata sampah
		rata = float64(total) / float64(*banyakOrang)
		fmt.Printf("Rata - rata sampah: %.2f kg / transaksi\n", rata)
	}
	
	fmt.Println(" ")
	fmt.Println("..................................................................................")
	fmt.Println("======================================== <<>> ====================================")
	halamanUtama_kedua(user, banyakOrang)
}
