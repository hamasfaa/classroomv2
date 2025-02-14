package entities

import "time"

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Tugas struct {
	TDJudul     string     `json:"td_judul"`
	TDDeskripsi string     `json:"td_deskripsi"`
	TDDeadline  time.Time  `json:"td_deadline"`
	Files       []FileData `json:"files"`
}

type FileData struct {
	TFNama    string `json:"tf_nama"`
	TFContent string `json:"tf_content"`
	TFType    string `json:"tf_type"`
}
