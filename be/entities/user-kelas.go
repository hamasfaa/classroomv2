package entities

type UserKelas struct {
	KelasKID string `json:"kelas_k_id"`
	UserUID  string `json:"user_u_id"`

	Kelas Kelas `gorm:"foreignKey:KelasKID;references:KID" json:"kelas,omitempty"`
	User  User  `gorm:"foreignKey:UserUID;references:UID" json:"user,omitempty"`
}
