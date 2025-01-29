package entities

import "time"

type AbsenDosen struct {
	ADID            string    `gorm:"primaryKey;type:varchar(255)" json:"ad_id"`
	ADTanggalDibuat time.Time `gorm:"type:date" json:"ad_tanggal_dibuat"`
	ADDeskripsi     string    `gorm:"type:longtext" json:"ad_deskripsi"`
	ADPertemuan     uint      `gorm:"type:smallint" json:"ad_pertemuan"`
	ADKode          string    `gorm:"type:char(6)" json:"ad_kode"`
	KelasKID        string    `gorm:"type:varchar(255)" json:"kelas_k_id"`
	UserUID         string    `gorm:"type:varchar(255)" json:"user_uid"`

	Kelas Kelas `gorm:"foreignKey:KelasKID;references:KID" json:"kelas,omitempty"`
	User  User  `gorm:"foreignKey:UserUID;references:UID" json:"user,omitempty"`
}
