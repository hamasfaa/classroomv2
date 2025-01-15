package entities

import "time"

type TugasMahasiswa struct {
	TM_ID               uint64
	TM_WaktuPengumpulan time.Time
	TM_Status           bool
	TM_FileTugas        string
	TM_NilaiTugas       uint8
	Tugas_Dosen_TD_ID   TugasDosen
	Kelas_K_ID          Kelas
	User_U_ID           User
}
