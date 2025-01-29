package entities

import "time"

type Kelas struct {
	KID            string    `gorm:"primaryKey;type:varchar(255)" json:"k_id"`
	KMataKuliah    string    `gorm:"type:varchar(50)" json:"k_mata_kuliah"`
	KNamaKelas     string    `gorm:"type:varchar(50)" json:"k_nama_kelas"`
	KTanggalDibuat time.Time `gorm:"type:date" json:"k_tanggal_dibuat"`
	KKodeKelas     string    `gorm:"type:char(6)" json:"k_kode_kelas"`
}
