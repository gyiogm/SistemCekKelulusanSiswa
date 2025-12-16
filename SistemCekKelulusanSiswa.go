package main

import (
	"fmt"
)

func main() {
	var jumlahSiswa int
	var nama, nis string
	var nilaiTugas, nilaiUTS, nilaiUAS float64
	var nilaiAkhir float64
	var status string
	var valid bool

	fmt.Print("Masukkan jumlah siswa: ")
	fmt.Scan(&jumlahSiswa)

	for i := 1; i <= jumlahSiswa; i++ {
		fmt.Println("\n==============================")
		fmt.Println("Data siswa ke-", i)

		// Input nama dan NIS
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&nama)

		fmt.Print("Masukkan NIS: ")
		fmt.Scan(&nis)

		// Validasi nilai Tugas
		valid = false
		for valid == false {
			fmt.Print("Masukkan nilai Tugas (0-100): ")
			fmt.Scan(&nilaiTugas)

			if nilaiTugas >= 0 && nilaiTugas <= 100 {
				valid = true
			} else {
				fmt.Println("Error: nilai harus 0–100.")
			}
		}

		// Validasi nilai UTS
		valid = false
		for valid == false {
			fmt.Print("Masukkan nilai UTS (0-100): ")
			fmt.Scan(&nilaiUTS)

			if nilaiUTS >= 0 && nilaiUTS <= 100 {
				valid = true
			} else {
				fmt.Println("Error: nilai harus 0–100.")
			}
		}

		// Validasi nilai UAS
		valid = false
		for valid == false {
			fmt.Print("Masukkan nilai UAS (0-100): ")
			fmt.Scan(&nilaiUAS)

			if nilaiUAS >= 0 && nilaiUAS <= 100 {
				valid = true
			} else {
				fmt.Println("Error: nilai harus 0–100.")
			}
		}

		// Hitung nilai akhir
		nilaiAkhir = (0.3 * nilaiTugas) + (0.3 * nilaiUTS) + (0.4 * nilaiUAS)

		// Tentukan status kelulusan
		if nilaiAkhir >= 75 {
			status = "Lulus"
		} else if nilaiAkhir >= 60 {
			status = "Remedial"
		} else {
			status = "Tidak Lulus"
		}

		// Output
		fmt.Println("\n===== HASIL =====")
		fmt.Println("Nama        :", nama)
		fmt.Println("NIS         :", nis)
		fmt.Println("Nilai Akhir :", nilaiAkhir)
		fmt.Println("Status      :", status)
		fmt.Println("==============================")
	}

	fmt.Println("\nProses pengecekan selesai.")
}
