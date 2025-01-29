package entities

type AbsenMahasiswa struct {
	AMID           string `gorm:"primaryKey;type:varchar(255)" json:"am_id"`
	AMStatus       uint8  `gorm:"type:int" json:"am_status"`
	AMDeskripsi    string `gorm:"type:longtext" json:"am_deskripsi"`
	AbsenDosenADID string `gorm:"type:varchar(255)" json:"absen_dosen_ad_id"`
	UserUID        string `gorm:"type:varchar(255)" json:"user_u_id"`
	KelasKID       string `gorm:"type:varchar(255)" json:"kelas_k_id"`

	AbsenDosen AbsenDosen `gorm:"foreignKey:AbsenDosenADID;references:ADID" json:"absen_dosen,omitempty"`
	User       User       `gorm:"foreignKey:UserUID;references:UID" json:"user,omitempty"`
	Kelas      Kelas      `gorm:"foreignKey:KelasKID;references:KID" json:"kelas,omitempty"`
}
