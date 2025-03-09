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

type Status struct {
	Status bool `json:"td_status"`
}

type PublicUser struct {
	UNama         string    `json:"u_nama"`
	UEmail        string    `json:"u_email"`
	URole         string    `json:"u_role"`
	UTanggalLahir time.Time `json:"u_tanggal_lahir"`
	UNoPonsel     string    `json:"u_no_ponsel"`
	UAlamat       string    `json:"u_alamat"`
	UFoto         string    `json:"u_foto"`
}
