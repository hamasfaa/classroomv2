package entities

import "time"

type User struct {
	UID           string    `gorm:"primary_key" json:"u_id"`
	UNama         string    `json:"u_nama"`
	UEmail        string    `json:"u_email"`
	UPassword     string    `json:"u_password"`
	URole         string    `json:"u_role"`
	UTanggalLahir time.Time `json:"u_tanggal_lahir"`
	UNoPonsel     string    `json:"u_no_ponsel"`
	UAlamat       string    `json:"u_alamat"`
	UFoto         string    `json:"u_foto"`
}
