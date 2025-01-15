package entities

import "time"

type Kelas struct {
	K_ID            uint64
	K_MataKuliah    string
	K_NamaKelas     string
	K_TanggalDibuat time.Time
	K_KodeKelas     string
}
