package entities

import "time"

type AbsenDosen struct {
	AD_ID            uint64
	AD_TanggalDibuat time.Time
	AD_Deskripsi     string
	AD_Pertemuan     uint8
	AD_Kode          string
	Kelas_K_ID       Kelas
	User_U_ID        User
}
