package model

import "fmt"

type Karyawan struct {
	Id        int64
	Nama      string
	Kehadiran bool
}

func PenampilanDaftarKaryawan(list []Karyawan) {
	for idx, karyawan := range list {
		fmt.Printf("Index: %d untuk karyawan Id: %d - Karyawan Dengan Nama: %s - Status Kehadiran: %s\n", idx+1, karyawan.Id, karyawan.Nama, kehadiran(karyawan.Kehadiran))
	}
}

// GX PUNYA TERNARY !!! - buat print aja :<
func kehadiran(test bool) string {
	if test {
		return "Hadir"
	}
	return "Tidak Hadir"
}
