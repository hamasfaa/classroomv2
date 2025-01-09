package entities

import "time"

type absen_dosen struct {
	AD_ID            uint64
	AD_TanggalDibuat time.Time
	AD_Deskripsi     string
	AD_Pertemuan     uint8
	AD_Kode          uint64
}
