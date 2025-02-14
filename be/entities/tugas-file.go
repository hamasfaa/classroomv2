package entities

type TugasFile struct {
	TFID      string `gorm:"primaryKey;type:varchar(255)" json:"tf_id"`
	TFNama    string `gorm:"type:varchar(255)" json:"tf_nama"`
	TFPath    string `gorm:"type:varchar(255)" json:"tf_path"`
	TugasTDID string `gorm:"type:varchar(255)" json:"tugas_td_id"`

	TugasDosen TugasDosen `gorm:"foreignKey:TugasTDID;references:TDID" json:"tugas,omitempty"`
}
