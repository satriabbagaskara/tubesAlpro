package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
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
	dummy(&user)
}

func dummy(user *tabDataSampah) {
	var banyakOrang int

	user[1] = dataSampah{id: 345, nama: "Satria", beratSampah: 3, jenisSampah: "organik", tanggal: "12-12-26"} 
	user[2] = dataSampah{id: 563,nama: "Bima", beratSampah: 10, jenisSampah: "organik", tanggal: "13-12-26"} 
	user[3] = dataSampah{id: 100,nama: "Adnan", beratSampah: 6, jenisSampah: "anorganik", tanggal: "13-12-26"} 
	user[4] = dataSampah{id: 832,nama: "Raihan", beratSampah: 11, jenisSampah: "organik", tanggal: "14-12-26"} 
	user[5] = dataSampah{id: 105,nama: "Antontonan", beratSampah: 20, jenisSampah: "anorganik", tanggal: "14-12-26"} 
	user[6] = dataSampah{id: 521,nama: "Yudish", beratSampah: 4, jenisSampah: "anorganik", tanggal: "15-12-26"} 
	user[7] = dataSampah{id: 739,nama: "Error212", beratSampah: 3, jenisSampah: "anorganik", tanggal: "15-12-26"} 
	user[8] = dataSampah{id: 518,nama: "Gacor222", beratSampah: 121, jenisSampah: "anorganik", tanggal: "16-12-26"} 
	user[9] = dataSampah{id: 933,nama: "Elmacho", beratSampah: 41, jenisSampah: "anorganik", tanggal: "16-12-26"}
	user[10] = dataSampah{id: 742,nama: "Kritang", beratSampah: 27, jenisSampah: "organik", tanggal: "16-12-26"}
	banyakOrang = 10

	halamanUtama(user, &banyakOrang)
}

// interface halaman utama
func halamanUtama(user *tabDataSampah, banyakOrang *int) {
	var inputPilihan int
	fmt.Println("__________________________________________________")
	fmt.Println("|            \033[93m--  SELAMAT DATANG  --\033[0m              |")
	fmt.Println("|________________________________________________|")
	fmt.Println("                App  Waste  Track                 ")
	fmt.Println("         Algoritma pemrograman 2 - 2026           ")
	fmt.Println("                 Satria X Zikri                   ")
	fmt.Println("__________________________________________________")
	
	fmt.Println("| 1. Tambah data        | 2. Edit data            |")
	fmt.Println("|_______________________|_________________________|")
	fmt.Println("| 3. Hapus data         | 4. Statistik            |")
	fmt.Println("|_______________________|_________________________|")
	fmt.Println("| 5. Cari data warga    | 6. List data warga      |")
	fmt.Println("|_______________________|_________________________|")
	fmt.Println("| 7. Keluar aplikasi    | 8. Coming Soon          |")
	fmt.Println("|_______________________|_________________________|")
	
	fmt.Print("\nMasukkan pilihan kamu yang mau kami jalankan: ")
	fmt.Scan(&inputPilihan)

	// pilihan input
	if inputPilihan == 1 {
		inputDataTambahan(user, banyakOrang)
	} else if inputPilihan == 2{
		editData(user, banyakOrang)
	} else if inputPilihan == 3{
		hapusData(user, banyakOrang)
	} else if inputPilihan == 4{
		statistikSampah(user, banyakOrang)
	} else if inputPilihan == 6{
		tampilData(user, banyakOrang)
	} else if inputPilihan == 7{
		fmt.Printf("\n\n\n -- Terimakasih telah menggunakan Waste-Track --\n\n\n\n\n")
		os.Exit(0)
	} else if inputPilihan == 5{
		cariDataWarga(user, banyakOrang)
	} else if inputPilihan == 8 || inputPilihan > 8{
		fmt.Printf("\n\n         -- Fitur belum tersedia --\n")
		time.Sleep(2*time.Second)
		halamanUtama(user, banyakOrang)
	} else {
		fmt.Printf("\n\n         -- Pilihan tidak tersedia --\n")
		time.Sleep(2*time.Second)
		halamanUtama(user, banyakOrang)
	}
}
// memeriksa id random nya sudah ada / belum
func idSudahAda(user *tabDataSampah, banyakOrang *int, id int)bool{
	var i int
	for i = 1; i <= *banyakOrang; i++{
		if user[i].id == id{
			return true
		}
	}
	return false
}
// proses mendapatkan id secara random
func dapatkanID(user *tabDataSampah, banyakOrang *int) int{
	var randNum int
	for {
		randNum = rand.Intn(900) + 100
		if !idSudahAda(user, banyakOrang, randNum){
			return randNum
		}
	}
}

// di interface jika pilih input data sampah
func inputDataTambahan(user *tabDataSampah, banyakOrang *int) {
	var banyakOrangBaru, i, totalOrangSementara, ID int
	fmt.Println("__________________________________________________")
	fmt.Println("|                \033[36mINPUT DATA SAMPAH\033[0m               |")
	fmt.Println("|________________________________________________|")
	fmt.Println(" ")
	fmt.Println("Rule pengisian: ")
	fmt.Println("1. Nama tidak boleh lebih panjang dari 10 karakter")
	fmt.Println("2. Jenis sampah hanya terdiri dari ORGANIK atau ANORGANIK")
	fmt.Println("3. Untuk tanggal pengiriman isi dengan format DD-MM-YY ")
	fmt.Println(" ")
	fmt.Println("Masukkan banyak orang untuk di data")
	fmt.Print("Input: ")
	fmt.Scan(&banyakOrangBaru)
	totalOrangSementara = *banyakOrang + banyakOrangBaru

	for i = *banyakOrang + 1; i <= totalOrangSementara; i++ {
		fmt.Println("---- <<>> ----")
		fmt.Print("Nama: ")
		fmt.Scan(&user[i].nama)
		fmt.Print("Jenis sampah: ")
		fmt.Scan(&user[i].jenisSampah)
		fmt.Print("Berat sampah: ")
		fmt.Scan(&user[i].beratSampah)
		fmt.Print("Tanggal pengiriman: ")
		fmt.Scan(&user[i].tanggal)
		
		fmt.Println("                ---------------------                 ")
		fmt.Println("                   \033[3mGenerating ID\033[0m")
		ID = dapatkanID(user, banyakOrang)
		user[i].id = ID
		*banyakOrang = i
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("                  \033[3mID Created\033[0m : %d\n", user[i].id)
		fmt.Println("                ---------------------                 ")
		time.Sleep(300 * time.Millisecond)
	}

	*banyakOrang = totalOrangSementara
	fmt.Println("======================<<>>========================")
	fmt.Println(" ")
	fmt.Println(" ")
	halamanUtama(user, banyakOrang)
}

// di interface jika memilih edit data
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
		fmt.Println("__________________________________________________")
		fmt.Println("|                \033[36mEdit DATA SAMPAH\033[0m                |")
		fmt.Println("|________________________________________________|")
		fmt.Println("| 1. ID                   | 2. Nama              |")
		fmt.Println("|_________________________|______________________|")
		fmt.Println("| 3. Jenis sampah         | 4. Berat sampah      |")
		fmt.Println("|_________________________|______________________|")
		fmt.Println("| 5. Tanggal pengiriman   | 6. Kembali           |")
		fmt.Println("|_________________________|______________________|")
		fmt.Println(" ")
		fmt.Println("Pilih data yang mau diubah: ")
		fmt.Print("Input: ")
		fmt.Scan(&inputPilihan)
		
		// Ganti ID //
		if inputPilihan == 1 {
			var ID int

			fmt.Print("ID yang mau diganti: ")
			fmt.Scan(&IDpilihan)
			
			for i = 1; i <= *banyakOrang; i++ {
				if user[i].id == IDpilihan {
					fmt.Println("ID ditemukan!")
					fmt.Printf("Atas nama %s\n", user[i].nama)
					time.Sleep(2*time.Second)
					fmt.Println("                ---------------------                 ")
					fmt.Println("                   \033[3mGenerating ID\033[0m")
					ID = dapatkanID(user, banyakOrang)
					user[i].id = ID
					time.Sleep(785 * time.Millisecond)
					fmt.Printf("                  \033[3mID Created\033[0m : %d\n", user[i].id)
					fmt.Println("                ---------------------                 ")
					time.Sleep(1700 * time.Millisecond)
					keberadaan = true
				}
			}
			if !keberadaan {
				fmt.Println("Data tidak dapat ditemukan")
			}
			
		}
		
		// Ganti Nama //
		var namaBaru string
		if inputPilihan == 2 {
			fmt.Println("----------------------------------------------")
			fmt.Print("Masukkan ID: ")
			fmt.Scan(&gantiID)
			fmt.Print("Nama baru: ")
			fmt.Scan(&namaBaru)
			
			for i = 1; i <= *banyakOrang; i++ {
				if user[i].id == gantiID {
					time.Sleep(500*time.Millisecond)
					fmt.Println("           -- ID ditemukan! --")
					user[i].nama = namaBaru
					fmt.Printf("           Nama sekarang: %s\n", user[i].nama)
					keberadaan = true
				}
			}
			if keberadaan == false {
				time.Sleep(500*time.Millisecond)
				fmt.Println("           -- ID tidak dapat ditemukan -- ")
			}
		}
		
		// Ganti jenis sampah //
		var jenisSampahBaru string
		if inputPilihan == 3 {
			fmt.Print("ID warga: ")
			fmt.Scan(&gantiID)
			fmt.Print("Jenis baru: ")
			fmt.Scan(&jenisSampahBaru)
			
			for i = 1; i <= *banyakOrang; i++ {
				if gantiID == user[i].id {
					time.Sleep(500 *time.Millisecond)
					fmt.Println("           -- ID ditemukan! --")
					user[i].jenisSampah = jenisSampahBaru
					fmt.Printf("    Jenis sampah sekarang: %s\n", user[i].jenisSampah)
					keberadaan = true
				}
			}
			
			if !keberadaan {
				time.Sleep(500*time.Millisecond)
				fmt.Println("ID tidak dapat ditemukan")
			}
		}
		
		// Edit berat sampah // 
		var beratBaru int
		if inputPilihan == 4 {
			fmt.Print("ID warga: ")
			fmt.Scan(&gantiID)
			fmt.Print("Berat baru: ")
			fmt.Scan(&beratBaru)
			
			for i = 1; i <= *banyakOrang; i++ {
				if gantiID == user[i].id {
					time.Sleep(500*time.Millisecond)
					fmt.Println("           -- ID ditemukan! --")
					user[i].beratSampah = beratBaru
					fmt.Printf("          Berat sampah sekarang: %d\n", user[i].beratSampah)
					keberadaan = true
				}
			}
			
			if !keberadaan {
				time.Sleep(500*time.Millisecond)
				fmt.Println("ID tidak dapat ditemukan")
			}
		}
		
		// Ganti tanggal pengiriman //
		var tanggalBaru string
		
		if inputPilihan == 5 {
			fmt.Print("ID warga: ")
			fmt.Scan(&gantiID)
			
			fmt.Print("Tanggal baru: ")
			fmt.Scan(&tanggalBaru)
			for i = 1; i <= *banyakOrang; i++ {
				if user[i].id == gantiID {
					time.Sleep(500*time.Millisecond)
					fmt.Println("           -- ID ditemukan! --")
					user[i].tanggal = tanggalBaru
					fmt.Printf("   Tanggal transaksi sekarang: %s\n", user[i].tanggal)
					keberadaan = true
				}
			}
			
			if !keberadaan {
				time.Sleep(500*time.Millisecond)
				fmt.Println("ID tidak dapat ditemukan")
			}
		}
		
		if inputPilihan == 6 {
			time.Sleep(600*time.Millisecond)
			fmt.Println(" ")
			fmt.Println(" ")
			halamanUtama(user, banyakOrang)
			return
		}

		if inputPilihan > 6 {
			time.Sleep(500*time.Millisecond)
			fmt.Println("Inputan tidak tersedia")
			editData(user, banyakOrang)
		}
	}
	
	fmt.Println("======================<<>>========================")
	fmt.Println(" ")
	fmt.Println(" ")
	halamanUtama(user, banyakOrang)
}

// Procedure untuk menampilkan data yang dipilih
func tampilData(user *tabDataSampah, banyakOrang *int) {
	var i int
	
 	fmt.Println("__________________________________________________")
	fmt.Println("|                \033[36mDATA LIST WARGA\033[0m                 |")
	fmt.Println("|________________________________________________|")
	if *banyakOrang == 0 {
		fmt.Println("\n---")
		fmt.Println("Maaf, Belum ada data sampah\nSilahkan input datanya terlebih dahulu...")
		halamanUtama(user, banyakOrang)
		return
	} else {
		for i = 1; i <= *banyakOrang; i++ {
			fmt.Printf("- ID:                  | %d                       \n", user[i].id)
			fmt.Printf("- Nama:                | %s                       \n", user[i].nama)
			fmt.Printf("- Jenis sampah:        | %s                       \n", user[i].jenisSampah)
			fmt.Printf("- Berat sampah:        | %d                       \n", user[i].beratSampah)
			fmt.Printf("- Tanggal pengiriman:  | %s                       \n", user[i].tanggal)
			fmt.Println("----------------------<<>>------------------------")
			time.Sleep(180*time.Millisecond)
		}
	}

	fmt.Println("======================<<>>========================")
	fmt.Println(" ")
	fmt.Println(" ")
	halamanUtama(user, banyakOrang)
}

// Procedure Hapus data warga yang dipilih
func hapusData(user *tabDataSampah, banyakOrang *int) {
	var hapusWarga, i int
	var keberadaan int
	
	if *banyakOrang == 0{
		fmt.Println("Tidak ada data warga yang terdaftar")
		fmt.Println("Silahkan isi data transaksi terlebih dahulu")
	} else {
		keberadaan = -1
		fmt.Println("__________________________________________________")
		fmt.Println("|                \033[36mHAPUS DATA WARGA\033[0m                |")
		fmt.Println("|________________________________________________|")
		fmt.Println("Masukkan ID warga yang mau dihapus")
		fmt.Print("Input: ")
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
	}
	fmt.Println("\n======================<<>>========================")
	fmt.Println(" ")
	fmt.Println(" ")
	halamanUtama(user, banyakOrang)
}
// salin data yg ada
func salinData(user *tabDataSampah, temp *tabDataSampah, banyakOrang *int){
	var i int
	for i = 1; i <= *banyakOrang; i++{
		temp[i] = user[i]
	}
}

// insertion sort menggunakan ID
func insertionSortID(user *tabDataSampah, banyakOrang *int){
	var pass, i int
	var temp dataSampah

	for pass = 2; pass <= *banyakOrang; pass++{
		temp = user[pass]
		i = pass
		for i > 1 && temp.id < user[i-1].id{
			user[i] = user[i-1]
			i--
		}
		user[i] = temp
	}
}
// insertion sort menggunakan Nama
func insertionSortNama(user *tabDataSampah, banyakOrang *int){
	var pass, i int
	var temp dataSampah

	for pass = 2; pass <= *banyakOrang; pass++{
		temp = user[pass]
		i = pass
		for i > 1 && temp.nama < user[i-1].nama{
			user[i] = user[i-1]
			i--
		}
		user[i] = temp
	}
}

// binary search menggunakan ID
func binarySearchID(user *tabDataSampah, banyakOrang *int, cariID int)int {
	var kiri, kanan, mid int

	kiri = 1
	kanan = *banyakOrang
	for kiri <= kanan {
		mid = (kiri + kanan) / 2

		if user[mid].id == cariID{
			return mid
		} else if cariID < user[mid].id{
			kanan = mid - 1
		} else {
			kiri = mid + 1
		}
		
	}
	return -1
}
// binary search menggunakan Nama
func binarySearchNama(user *tabDataSampah, banyakOrang *int, cariNama string)int{
	var kiri, kanan, mid int

	kiri = 1
	kanan = *banyakOrang
	for kiri <= kanan{
		mid = (kiri + kanan) / 2
		if user[mid].nama == cariNama{
			return mid
		} else if cariNama < user[mid].nama{
			kanan = mid - 1
		} else {
			kiri = mid + 1
		}
	}
	return -1
}

// sequential search menggunakan ID
func sequentialSearchID(user *tabDataSampah, banyakOrang *int, cariID int)int{
	var i int
	for i = 1; i <= *banyakOrang; i++{
		if user[i].id == cariID{
			return i
		}
	}
	return -1
}

// sequential search menggunakan Nama
func sequentialSearchNama(user *tabDataSampah, banyakOrang *int, cariNama string)int{
	var i int
	for i = 1; i <= *banyakOrang; i++{
		if user[i].nama == cariNama{
			return i
		}
	}
	return -1
}

// Procedure outputnya
func tampilanData(data dataSampah){
	fmt.Println("\nData ditemukan!")
	fmt.Println("======================<<>>========================")
	fmt.Printf("ID: %d\n", data.id)
	fmt.Printf("Nama: %s\n", data.nama)
	fmt.Printf("Berat sampah: %d kg\n", data.beratSampah)
	fmt.Printf("Jenis Sampah: %s\n", data.jenisSampah)
	fmt.Printf("Tanggal Pengiriman: %s\n", data.tanggal)
	fmt.Println("======================<<>>========================")
}

// Procedure cari data warga
func cariDataWarga(user *tabDataSampah, banyakOrang *int) {
	var inputPilihan, cariID, idx int
	var cariNama string
	var temp tabDataSampah

	fmt.Println("____________________________________________________")
	fmt.Println("|              \033[36mMENCARI DATA WARGA\033[0m                  |")
	fmt.Println("|__________________________________________________|")
	fmt.Println(" ")
	fmt.Println("Pilihlah metode pencarian yang anda inginkan: ")
	fmt.Println("+--------------------------------------------------+")
	fmt.Println("| 1. Binary Search (ID)       | 4. Sequential Search (Nama)    |")
	fmt.Println("| 2. Binary Search (Nama)     | 5. Kembali                     |")
	fmt.Println("| 3. Sequential Search (ID)   |")
	fmt.Println("+--------------------------------------------------+")
	fmt.Println(" ")
	fmt.Print("Input: ")
	fmt.Scan(&inputPilihan)

	if inputPilihan == 1 {
		if *banyakOrang == 0{
			fmt.Println("Data tidak ada. Silahkan isi data terlebih dahulu...")
			fmt.Println(" ")
		} else {
			fmt.Println("|-----------------------------------------------|")
			fmt.Println("|          BINARY SEARCH BERDASARKAN ID         |")
			fmt.Println("|-----------------------------------------------|")
			fmt.Print("Masukkan ID warga yang mau dicari: ")
			fmt.Scan(&cariID)
		
			salinData(user, &temp, banyakOrang)
			insertionSortID(&temp, banyakOrang)

			idx = binarySearchID(&temp, banyakOrang, cariID)

			if idx == -1{
				fmt.Println("ID tidak ditemukan....")
			} else {
				tampilanData(temp[idx])
			}
		}

	} else if inputPilihan == 2 {
		fmt.Println("|-----------------------------------------------|")
		fmt.Println("|          BINARY SEARCH BERDASARKAN NAMA       |")
		fmt.Println("|-----------------------------------------------|")
		fmt.Print("Masukkan Nama warga yang mau dicari: ")
		fmt.Scan(&cariNama)
		fmt.Println("")

		salinData(user, &temp, banyakOrang)
		insertionSortNama(&temp, banyakOrang)

		idx = binarySearchNama(&temp, banyakOrang, cariNama)
		if idx == -1{
			fmt.Println("Nama tidak dapat ditemukan")
		} else {
			tampilanData(temp[idx])
		}

	} else if inputPilihan == 3 {
		fmt.Println("|-----------------------------------------------|")
		fmt.Println("|        SEQUENTIAL SEARCH BERDASARKAN ID       |")
		fmt.Println("|-----------------------------------------------|")
		fmt.Print("Masukkan ID warga yang mau dicari: ")
		fmt.Scan(&cariID)

		idx = sequentialSearchID(user, banyakOrang, cariID)
		if idx == -1{
			fmt.Println("ID tidak ditemukan.....")
		} else {
			tampilanData(user[idx])
		}

	} else if inputPilihan == 4 {
		fmt.Println("|-----------------------------------------------|")
		fmt.Println("|        SEQUENTIAL SEARCH BERDASARKAN NAMA     |")
		fmt.Println("|-----------------------------------------------|")
		fmt.Println("Masukkan Nama warga yang mau dicari: ")
		fmt.Scan(&cariNama)

		idx = sequentialSearchNama(user, banyakOrang, cariNama)
		if idx == -1{
			fmt.Println("Nama tidak ditemukan.......")
		} else {
			tampilanData(user[idx])
		}

	} else if inputPilihan == 5{
		fmt.Println("Kembali ke menu utama......")
	} else {
		fmt.Println(" ")
		fmt.Println("          -- Pilihan tidak tersedia --  ")
		fmt.Println(" ")
	}

	fmt.Println("======================<<>>========================")
	fmt.Println(" ")
	fmt.Println(" ")
	halamanUtama(user, banyakOrang)
}

// Procedure Statistik Data Sampah Mingguan
func statistikSampah(user *tabDataSampah, banyakOrang *int) {
	var i, total int
	var rata float64
	var idxMax, idxMin int

	fmt.Println("__________________________________________________")
	fmt.Println("|              \033[36mSTATISTIK DATA WARGA\033[0m              |")
	fmt.Println("|________________________________________________|")

	if *banyakOrang == 0{
		fmt.Println(" ")
		fmt.Println("Belum ada transaksi sampah.")
		fmt.Println("Silahkan isi data terlebih dahulu.")
	} else {
		// total sampah
		total = 0
		for i = 1; i <= *banyakOrang; i++{
			total = total + user[i].beratSampah
		}
	
		// sampah tertinggi & terendah
		idxMax = 1
		idxMin = 1
		for i = 2; i <= *banyakOrang; i++{
			if user[i].beratSampah > user[idxMax].beratSampah{
				idxMax = i
			}
			if user[i].beratSampah < user[idxMin].beratSampah{
				idxMin = i
			}
			
		}
		// rata-rata sampah
		rata = float64(total) / float64(*banyakOrang)

		// cetak hasil
		fmt.Println(" ")
		fmt.Printf("Total seluruh sampah: %d kg\n", total)
		fmt.Println("---------------------<<>>-------------------")
		fmt.Printf("Jumlah transaksi: %d\n", *banyakOrang)
		fmt.Println("---------------------<<>>-------------------")
		fmt.Printf("Sampah terbanyak: %d kg, atas nama %s\n", user[idxMax].beratSampah, user[idxMax].nama)
		fmt.Println("---------------------<<>>-------------------")
		fmt.Printf("Sampah terendah: %d kg, atas nama %s\n", user[idxMin].beratSampah, user[idxMin].nama)
		fmt.Println("---------------------<<>>-------------------")
		fmt.Printf("Rata - rata sampah: %.2f kg / transaksi\n", rata)
	}
	fmt.Println(" ")
	fmt.Println("======================<<>>========================")
	fmt.Println(" ")
	fmt.Println(" ")
	halamanUtama(user, banyakOrang)
}

// any question?
