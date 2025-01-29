package entities

import "time"

type AbsenDosen struct {
	ADID            string    `gorm:"primaryKey" json:"ad_id"`
	ADTanggalDibuat time.Time `json:"ad_tanggal_dibuat"`
	ADDeskripsi     string    `json:"ad_deskripsi"`
	ADPertemuan     uint      `json:"ad_pertemuan"`
	ADKode          string    `json:"ad_kode"`
	KelasKID        string    `json:"kelas_k_id"`
	UserUID         string    `json:"user_uid"`

	Kelas Kelas `gorm:"foreignKey:KelasKID;references:KID" json:"kelas,omitempty"`
	User  User  `gorm:"foreignKey:UserUID;references:UID" json:"user,omitempty"`
}
