package entities

type AbsenMahasiswa struct {
	AM_ID             uint64
	AM_Status         uint8
	AM_Deskripsi      string
	Absen_Dosen_AD_ID AbsenDosen
	Kelas_K_ID        Kelas
	User_U_ID         User
}
