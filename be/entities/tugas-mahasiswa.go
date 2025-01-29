package entities

import "time"

type TugasMahasiswa struct {
	TMID               string    `gorm:"primaryKey;type:varchar(255)" json:"tm_id"`
	TMWaktuPengumpulan time.Time `gorm:"type:datetime" json:"tm_waktu_pengumpulan"`
	TMStatus           bool      `gorm:"type:boolean" json:"tm_status"`
	TMFileTugas        string    `gorm:"type:varchar(255)" json:"tm_file_tugas"`
	TMNilaiTugas       int       `gorm:"type:int" json:"tm_nilai_tugas"`
	TugasDosenTDID     string    `gorm:"type:varchar(255)" json:"tugas_dosen_td_id"`
	KelasKID           string    `gorm:"type:varchar(255)" json:"kelas_k_id"`
	UserUID            string    `gorm:"type:varchar(255)" json:"user_u_id"`

	TugasDosen TugasDosen `gorm:"foreignKey:TugasDosenTDID;references:TDID" json:"tugas_dosen"`
	Kelas      Kelas      `gorm:"foreignKey:KelasKID;references:KID" json:"kelas"`
	User       User       `gorm:"foreignKey:UserUID;references:UID" json:"user"`
}
