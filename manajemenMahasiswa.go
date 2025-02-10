package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	green  = "\033[32m"
	red    = "\033[31m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	reset  = "\033[0m"
)

type Mahasiswa struct {
	NIM     string
	Nama    string
	Jurusan string
	IPK     float64
}

var dataMahasiswa []Mahasiswa

func loadingAnimasi(pesan string) {
	fmt.Print(pesan)
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println(" ✔")
}

func tambahMahasiswa() {
	reader := bufio.NewReader(os.Stdin)

	var m Mahasiswa
	fmt.Print(yellow + "Masukkan NIM: " + reset)
	fmt.Scanln(&m.NIM)

	fmt.Print(yellow + "Masukkan Nama: " + reset)
	m.Nama, _ = reader.ReadString('\n')
	m.Nama = strings.TrimSpace(m.Nama)

	fmt.Print(yellow + "Masukkan Jurusan: " + reset)
	m.Jurusan, _ = reader.ReadString('\n')
	m.Jurusan = strings.TrimSpace(m.Jurusan)

	fmt.Print(yellow + "Masukkan IPK: " + reset)
	fmt.Scanln(&m.IPK)

	loadingAnimasi("Menyimpan data")
	dataMahasiswa = append(dataMahasiswa, m)
	fmt.Println(green + "Mahasiswa berhasil ditambahkan!" + reset)
}

func tampilkanMahasiswa() {
	if len(dataMahasiswa) == 0 {
		fmt.Println(red + "\nBelum ada mahasiswa yang terdaftar." + reset)
		return
	}

	// Define table width and separators - adjusted based on content
	const (
		noWidth      = 5  // Reduced from 6
		nimWidth     = 16 // Increased from 15
		namaWidth    = 35 // Increased from 30
		jurusanWidth = 25
		ipkWidth     = 8 // Reduced from 10
	)
	totalWidth := noWidth + nimWidth + namaWidth + jurusanWidth + ipkWidth

	// Print header with box drawing characters
	fmt.Print(blue)                                              // Set blue color for the entire table
	fmt.Println("\n╔" + strings.Repeat("═", totalWidth+3) + "╗") // +3 for extra padding
	fmt.Println("║" + center("DAFTAR MAHASISWA", totalWidth+3) + "║")
	fmt.Println("╠" + strings.Repeat("═", noWidth) + "╦" +
		strings.Repeat("═", nimWidth) + "╦" +
		strings.Repeat("═", namaWidth) + "╦" +
		strings.Repeat("═", jurusanWidth) + "╦" +
		strings.Repeat("═", ipkWidth) + "╣")

	// Print column headers
	fmt.Printf("║ %-*s║ %-*s║ %-*s║ %-*s║ %-*s║\n",
		noWidth-1, "NO",
		nimWidth-1, "NIM",
		namaWidth-1, "NAMA",
		jurusanWidth-1, "JURUSAN",
		ipkWidth-1, "IPK")

	fmt.Println("╠" + strings.Repeat("═", noWidth) + "╬" +
		strings.Repeat("═", nimWidth) + "╬" +
		strings.Repeat("═", namaWidth) + "╬" +
		strings.Repeat("═", jurusanWidth) + "╬" +
		strings.Repeat("═", ipkWidth) + "╣")

	// Print data rows
	for i, m := range dataMahasiswa {
		// Add extra padding in the format string
		fmt.Printf("║ %-*d ║ %-*s║ %-*s║ %-*s║ %*.2f ║\n",
			noWidth-2, i+1,
			nimWidth-1, m.NIM,
			namaWidth-1, truncateString(m.Nama, namaWidth-2),
			jurusanWidth-1, truncateString(m.Jurusan, jurusanWidth-2),
			ipkWidth-2, m.IPK)
	}

	// Print footer
	fmt.Println("╚" + strings.Repeat("═", noWidth) + "╩" +
		strings.Repeat("═", nimWidth) + "╩" +
		strings.Repeat("═", namaWidth) + "╩" +
		strings.Repeat("═", jurusanWidth) + "╩" +
		strings.Repeat("═", ipkWidth) + "╝" + reset)
}

func center(s string, width int) string {
	if len(s) >= width {
		return s[:width]
	}
	leftPad := (width - len(s)) / 2
	rightPad := width - len(s) - leftPad
	return strings.Repeat(" ", leftPad) + s + strings.Repeat(" ", rightPad)
}

func truncateString(s string, maxLen int) string {
	s = strings.TrimSpace(s)
	if len(s) <= maxLen {
		return s + strings.Repeat(" ", maxLen-len(s))
	}
	return s[:maxLen-3] + "..."
}

func cariMahasiswa() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(yellow + "Masukkan NIM yang dicari: " + reset)
	nim, _ := reader.ReadString('\n')
	nim = strings.TrimSpace(nim)

	for _, m := range dataMahasiswa {
		if m.NIM == nim {
			fmt.Printf(green+"Ditemukan - NIM: %s, Nama: %s, Jurusan: %s, IPK: %.2f\n"+reset, m.NIM, m.Nama, m.Jurusan, m.IPK)
			return
		}
	}
	fmt.Println(red + "Mahasiswa tidak ditemukan." + reset)
}

func updateMahasiswa() {
	if len(dataMahasiswa) == 0 {
		fmt.Println(red + "\nBelum ada mahasiswa yang terdaftar." + reset)
		return
	}

	tampilkanMahasiswa()
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(yellow + "\nMasukkan nomor mahasiswa yang ingin diperbarui: " + reset)
	var nomor int
	fmt.Scanln(&nomor)

	if nomor < 1 || nomor > len(dataMahasiswa) {
		fmt.Println(red + "Nomor mahasiswa tidak valid." + reset)
		return
	}

	index := nomor - 1
	fmt.Println(blue + "\nPilih data yang ingin diupdate:" + reset)
	fmt.Println("1. NIM")
	fmt.Println("2. Nama")
	fmt.Println("3. Jurusan")
	fmt.Println("4. IPK")
	fmt.Println("5. Semua Data")

	fmt.Print(yellow + "Pilih menu: " + reset)
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		fmt.Print(yellow + "Masukkan NIM baru: " + reset)
		fmt.Scanln(&dataMahasiswa[index].NIM)
	case 2:
		fmt.Print(yellow + "Masukkan Nama baru: " + reset)
		dataMahasiswa[index].Nama, _ = reader.ReadString('\n')
		dataMahasiswa[index].Nama = strings.TrimSpace(dataMahasiswa[index].Nama)
	case 3:
		fmt.Print(yellow + "Masukkan Jurusan baru: " + reset)
		dataMahasiswa[index].Jurusan, _ = reader.ReadString('\n')
		dataMahasiswa[index].Jurusan = strings.TrimSpace(dataMahasiswa[index].Jurusan)
	case 4:
		fmt.Print(yellow + "Masukkan IPK baru: " + reset)
		fmt.Scanln(&dataMahasiswa[index].IPK)
	case 5:
		fmt.Print(yellow + "Masukkan NIM baru: " + reset)
		fmt.Scanln(&dataMahasiswa[index].NIM)

		fmt.Print(yellow + "Masukkan Nama baru: " + reset)
		dataMahasiswa[index].Nama, _ = reader.ReadString('\n')
		dataMahasiswa[index].Nama = strings.TrimSpace(dataMahasiswa[index].Nama)

		fmt.Print(yellow + "Masukkan Jurusan baru: " + reset)
		dataMahasiswa[index].Jurusan, _ = reader.ReadString('\n')
		dataMahasiswa[index].Jurusan = strings.TrimSpace(dataMahasiswa[index].Jurusan)

		fmt.Print(yellow + "Masukkan IPK baru: " + reset)
		fmt.Scanln(&dataMahasiswa[index].IPK)
	default:
		fmt.Println(red + "Pilihan tidak valid!" + reset)
		return
	}

	loadingAnimasi("Memperbarui data")
	fmt.Println(green + "Data mahasiswa berhasil diperbarui!" + reset)
}

func hapusMahasiswa() {
	if len(dataMahasiswa) == 0 {
		fmt.Println(red + "\nBelum ada mahasiswa yang terdaftar." + reset)
		return
	}

	tampilkanMahasiswa()
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(yellow + "\nMasukkan nomor mahasiswa yang ingin dihapus: " + reset)
	var nomor int
	fmt.Scanln(&nomor)

	if nomor < 1 || nomor > len(dataMahasiswa) {
		fmt.Println(red + "Nomor mahasiswa tidak valid." + reset)
		return
	}

	index := nomor - 1
	fmt.Printf(yellow+"Anda akan menghapus mahasiswa dengan NIM: %s, Nama: %s\n"+reset,
		dataMahasiswa[index].NIM, dataMahasiswa[index].Nama)

	fmt.Print(red + "Apakah Anda yakin ingin menghapus data ini? (y/n): " + reset)
	konfirmasi, _ := reader.ReadString('\n')
	konfirmasi = strings.TrimSpace(konfirmasi)

	if konfirmasi == "y" {
		loadingAnimasi("Menghapus data")
		dataMahasiswa = append(dataMahasiswa[:index], dataMahasiswa[index+1:]...)
		fmt.Println(green + "Mahasiswa berhasil dihapus!" + reset)
	} else {
		fmt.Println(yellow + "Penghapusan dibatalkan." + reset)
	}
}

func main() {
	for {
		fmt.Println(blue + "\n=== Sistem Manajemen Mahasiswa ===" + reset)
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Tampilkan Mahasiswa")
		fmt.Println("3. Cari Mahasiswa")
		fmt.Println("4. Update Mahasiswa")
		fmt.Println("5. Hapus Mahasiswa")
		fmt.Println("6. Keluar")
		fmt.Print(yellow + "Pilih menu: " + reset)

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahMahasiswa()
		case 2:
			tampilkanMahasiswa()
		case 3:
			cariMahasiswa()
		case 4:
			updateMahasiswa()
		case 5:
			hapusMahasiswa()
		case 6:
			fmt.Println(green + "Terima kasih telah menggunakan sistem ini!" + reset)
			os.Exit(0)
		default:
			fmt.Println(red + "Pilihan tidak valid, coba lagi!" + reset)
		}
	}
}
