package entities

import "time"

type User struct {
	UID           string    `gorm:"primaryKey;type:varchar(255)" json:"u_id"`
	UNama         string    `gorm:"type:varchar(100)" json:"u_nama"`
	UEmail        string    `gorm:"type:varchar(100)" json:"u_email"`
	UPassword     string    `gorm:"type:varchar(255)" json:"u_password"`
	URole         string    `gorm:"type:enum('dosen','mahasiswa')" json:"u_role"`
	UTanggalLahir time.Time `gorm:"type:date" json:"u_tanggal_lahir"`
	UNoPonsel     string    `gorm:"type:varchar(15)" json:"u_no_ponsel"`
	UAlamat       string    `gorm:"type:varchar(255)" json:"u_alamat"`
	UFoto         string    `gorm:"type:varchar(255)" json:"u_foto"`
}
