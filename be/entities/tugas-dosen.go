package entities

import "time"

type TugasDosen struct {
	TDID            string    `gorm:"primaryKey" json:"td_id"`
	TDJudul         string    `json:"td_judul"`
	TDDeskripsi     string    `json:"td_deskripsi"`
	TDTanggalDibuat time.Time `json:"td_tanggal_dibuat"`
	TDDeadline      time.Time `json:"td_deadline"`
	TDStatus        bool      `json:"td_status"`
	TDFileSoal      string    `json:"td_file_soal"`
	KelasKID        string    `json:"kelas_k_id"`
	UserUID         string    `json:"user_u_id"`

	Kelas Kelas `gorm:"foreignKey:KelasKID;references:KID" json:"kelas,omitempty"`
	User  User  `gorm:"foreignKey:UserUID;references:UID" json:"user,omitempty"`
}
