package entities

import "time"

type UserKelas struct {
	Kelas_K_ID   Kelas
	User_U_ID    User
	TanggalAmbil time.Time
}
