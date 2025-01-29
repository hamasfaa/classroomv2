package entities

type UserKelas struct {
	KelasKID string `gorm:"type:varchar(255)" json:"kelas_k_id"`
	UserUID  string `gorm:"type:varchar(255)" json:"user_u_id"`

	Kelas Kelas `gorm:"foreignKey:KelasKID;references:KID" json:"kelas,omitempty"`
	User  User  `gorm:"foreignKey:UserUID;references:UID" json:"user,omitempty"`
}
