package entities

import "time"

type Kelas struct {
	KID            string    `gorm:"primary_key" json:"k_id"`
	KMataKuliah    string    `json:"k_mata_kuliah"`
	KNamaKelas     string    `json:"k_nama_kelas"`
	KTanggalDibuat time.Time `json:"k_tanggal_dibuat"`
	KKodeKelas     string    `json:"k_kode_kelas"`
}
