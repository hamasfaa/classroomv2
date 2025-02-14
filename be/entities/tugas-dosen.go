package entities

import "time"

type TugasDosen struct {
	TDID            string    `gorm:"primaryKey;type:varchar(255)" json:"td_id"`
	TDJudul         string    `gorm:"type:varchar(100)" json:"td_judul"`
	TDDeskripsi     string    `gorm:"type:longtext" json:"td_deskripsi"`
	TDTanggalDibuat time.Time `gorm:"type:date" json:"td_tanggal_dibuat"`
	TDDeadline      time.Time `gorm:"type:date" json:"td_deadline"`
	TDStatus        bool      `gorm:"type:boolean"  json:"td_status"`
	KelasKID        string    `gorm:"type:varchar(255)" json:"kelas_k_id"`
	UserUID         string    `gorm:"type:varchar(255)" json:"user_u_id"`

	Kelas Kelas `gorm:"foreignKey:KelasKID;references:KID" json:"kelas,omitempty"`
	User  User  `gorm:"foreignKey:UserUID;references:UID" json:"user,omitempty"`
}
