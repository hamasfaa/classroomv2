package entities

import (
	"database/sql"
	"time"
)

type User struct {
	U_ID           uint64
	U_Nama         string
	U_Email        string
	U_Password     string
	U_Role         string
	U_TanggalLahir time.Time
	U_NoPonsel     sql.NullString
	U_Alamat       sql.NullString
	U_Foto         sql.NullString
}
