package entities

import "time"

type TugasDosen struct {
	TD_ID            uint64
	TD_Judul         string
	TD_Deskripsi     string
	TD_TanggalDibuat time.Time
	TD_Deadline      time.Time
	TD_Status        bool
	TD_FileSoal      string
	Kelas_K_ID       Kelas
	User_U_ID        User
}
