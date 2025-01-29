package entities

import "time"

type TugasMahasiswa struct {
	TMID               string    `gorm:"primaryKey" json:"tm_id"`
	TMWaktuPengumpulan time.Time `json:"tm_waktu_pengumpulan"`
	TMStatus           bool      `json:"tm_status"`
	TMFileTugas        string    `json:"tm_file_tugas"`
	TMNilaiTugas       int       `json:"tm_nilai_tugas"`
	TugasDosenTDID     string    `json:"tugas_dosen_td_id"`
	KelasKID           string    `json:"kelas_k_id"`
	UserUID            string    `json:"user_u_id"`

	TugasDosen TugasDosen `gorm:"foreignKey:TugasDosenTDID;references:TDID" json:"tugas_dosen"`
	Kelas      Kelas      `gorm:"foreignKey:KelasKID;references:KID" json:"kelas"`
	User       User       `gorm:"foreignKey:UserUID;references:UID" json:"user"`
}
