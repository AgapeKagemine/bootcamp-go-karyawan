package main

import (
	"bufio"
	"fmt"
	"karyawan/model"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	DaftarKaryawan := []model.Karyawan{}
	for {
		cls()
		fmt.Printf("Daftar Kehadiran Karyawan: \n\n")
		if len(DaftarKaryawan) > 0 {
			model.PenampilanDaftarKaryawan(DaftarKaryawan)
		} else {
			fmt.Printf("Tidak ada data untuk sekarang\n")
		}
		fmt.Printf("\n[1] Tambah karyawan baru ke daftar karyawan\n")
		fmt.Printf("[2] Ubah status karyawan\n")
		fmt.Printf("[3] Hapus karyawan dari daftar karyawan\n")
		fmt.Printf("[4] Keluar program\n")

		menu := 0
		fmt.Printf("\nPilih Menu: ")
		if sc.Scan() {
			menu, _ = strconv.Atoi(sc.Text())
		}

		cls()
		if menu < 1 || menu > 5 {
			fmt.Printf("\n\nPilihan menu diantara 1 hingga 5!")
			fmt.Scanln()
			continue
		}

		switch menu {
		case 1:
			menu_tambah(&DaftarKaryawan, sc)
			fmt.Scanln()
		case 2:
			menu_update_status(&DaftarKaryawan, sc)
			fmt.Scanln()
		case 3:
			menu_hapus(&DaftarKaryawan, sc)
			fmt.Scanln()
		case 4:
			fmt.Printf("\n\nProgram Selesai")
			fmt.Scanln()
			cls()
			return
		}
	}
}

func menu_tambah(list *[]model.Karyawan, sc *bufio.Scanner) {
	nama := ""
	for {
		fmt.Printf("Masukkan Nama: ")
		if sc.Scan() {
			nama = sc.Text()
		}
		if len(nama) < 1 || nama == "" {
			fmt.Printf("\n\nNama tidak boleh kosong\n")
			continue
		}
		break
	}

	cls()
	fmt.Printf("\n\nDaftar karyawan berhasil di tambahkan\n")

	if len(*list) == 0 {
		*list = append(*list, model.Karyawan{
			Id:        1,
			Nama:      nama,
			Kehadiran: false,
		})
		return
	}

	*list = append(*list, model.Karyawan{
		Id:        (*list)[len(*list)-1].Id + 1,
		Nama:      nama,
		Kehadiran: false,
	})
}

func menu_update_status(list *[]model.Karyawan, sc *bufio.Scanner) {
	if len(*list) < 1 {
		fmt.Printf("\n\nTidak ada data untuk diubah\n")
		return
	}
	menu := 0
	for {
		cls()
		var err error
		model.PenampilanDaftarKaryawan(*list)
		fmt.Printf("\nPilih Id Karyawan: ")
		if sc.Scan() {
			menu, err = strconv.Atoi(sc.Text())
			if err != nil {
				cls()
				fmt.Printf("\n\nInput tidak valid!\n")
				fmt.Scanln()
				continue
			}
		}
		if menu > len(*list) {
			cls()
			fmt.Printf("\n\nIndex tidak di temukan %d!\n", menu)
			fmt.Scanln()
			continue
		}
		break
	}
	status := ""
	for {
		cls()
		fmt.Printf("\nPerubahan status karyawan(hadir/tidak): ")
		if sc.Scan() {
			status = sc.Text()
		}
		if status != "hadir" && status != "tidak" {
			cls()
			fmt.Printf("\n\nStatus karyawan tidak valid\n")
			fmt.Scanln()
			continue
		}
		break
	}
	cls()
	for idx, karyawan := range *list {
		if karyawan.Id == int64(menu) {
			if status == "hadir" {
				(*list)[idx].Kehadiran = true
			} else {
				(*list)[idx].Kehadiran = false
			}
			break
		}
	}
	fmt.Println("\n\nDaftar karyawan berhasil diubah")
}

func menu_hapus(list *[]model.Karyawan, sc *bufio.Scanner) {
	if len(*list) < 1 {
		fmt.Printf("\n\nTidak ada data untuk dihapus\n")
		return
	}
	menu := 0
	for {
		cls()
		var err error
		model.PenampilanDaftarKaryawan(*list)
		fmt.Printf("\nPilih Id Karyawan: ")
		if sc.Scan() {
			menu, err = strconv.Atoi(sc.Text())
			if err != nil {
				cls()
				fmt.Printf("\n\nInput tidak valid!\n")
				fmt.Scanln()
				continue
			}
		}
		if menu > len(*list) {
			cls()
			fmt.Printf("\n\nIndex tidak di temukan %d!\n", menu)
			fmt.Scanln()
			continue
		}
		for idx, karyawan := range *list {
			if karyawan.Id == int64(menu) {
				*list = append((*list)[:menu-1], (*list)[menu:]...)
				for i := idx; i < len(*list); i++ {
					(*list)[i].Id = (*list)[i].Id - 1
				}
			}
		}
		break
	}
	cls()
	fmt.Println("Daftar karyawan berhasil dihapus")
}

func cls() {
	fmt.Print("\033[100;1H\033[2J")
}
