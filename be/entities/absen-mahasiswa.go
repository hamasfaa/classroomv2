package entities

type AbsenMahasiswa struct {
	AMID           string `gorm:"primaryKey" json:"am_id"`
	AMStatus       uint8  `json:"am_status"`
	AMDeskripsi    string `json:"am_deskripsi"`
	AbsenDosenADID string `json:"absen_dosen_ad_id"`
	UserUID        string `json:"user_u_id"`
	KelasKID       string `json:"kelas_k_id"`

	AbsenDosen AbsenDosen `gorm:"foreignKey:AbsenDosenADID;references:ADID" json:"absen_dosen,omitempty"`
	User       User       `gorm:"foreignKey:UserUID;references:UID" json:"user,omitempty"`
	Kelas      Kelas      `gorm:"foreignKey:KelasKID;references:KID" json:"kelas,omitempty"`
}
