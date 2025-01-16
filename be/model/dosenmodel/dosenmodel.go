package dosenmodel

import (
	"be/config"
	"be/entities"
	"database/sql"
	"log"
)

func Data(id string) (*entities.User, error) {
	var user entities.User

	err := config.DB.QueryRow("SELECT U_Nama, U_Role, U_Foto FROM User WHERE U_ID = ?", id).Scan(&user.U_Nama, &user.U_Role, &user.U_Foto)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No user found with the given id")
			return nil, err
		}
		log.Println("Error querying the database:", err)
		return nil, err
	}
	return &user, nil
}

func ListKelas(id string) ([]entities.Kelas, error) {
	var kelasList []entities.Kelas

	rows, err := config.DB.Query("SELECT K.K_NamaKelas, K.K_MataKuliah FROM Kelas K INNER JOIN User_Kelas UK ON K.K_ID = UK.Kelas_K_ID WHERE UK.User_U_ID = ?", id)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No kelas found for the given id")
			return nil, err
		}
		log.Println("Error querying the database:", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var kelas entities.Kelas
		if err := rows.Scan(&kelas.K_NamaKelas, &kelas.K_MataKuliah); err != nil {
			log.Println("Error scanning the database:", err)
			return nil, err
		}

		kelasList = append(kelasList, kelas)
	}
	return kelasList, nil
}

func ListPertemuan(id string, bulan int, tahun int) ([]entities.AbsenDosen, error) {
	var pertemuanList []entities.AbsenDosen

	rows, err := config.DB.Query("SELECT AD.AD_Pertemuan, AD.AD_TanggalDibuat, UK.Kelas_K_ID, K.K_MataKuliah, AD.AD_ID FROM Absen_Dosen AD INNER JOIN User_Kelas UK ON AD.Kelas_K_ID = UK.Kelas_K_ID INNER JOIN Kelas K ON AD.Kelas_K_ID = K.K_ID WHERE UK.User_U_ID = ? AND MONTH(AD.AD_TanggalDibuat) = ? AND YEAR(AD.AD_TanggalDibuat) = ? GROUP BY AD.AD_ID", id, bulan, tahun)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No pertemuan found for the given id")
			return nil, err
		}
		log.Println("Error querying the database:", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var pertemuan entities.AbsenDosen
		if err := rows.Scan(&pertemuan.AD_Pertemuan, &pertemuan.AD_TanggalDibuat, &pertemuan.Kelas_K_ID.K_ID, &pertemuan.Kelas_K_ID.K_MataKuliah, &pertemuan.AD_ID); err != nil {
			log.Println("Error scanning the database:", err)
			return nil, err
		}
		pertemuanList = append(pertemuanList, pertemuan)
	}
	return pertemuanList, nil
}

func ListTugas(id string, bulan int, tahun int) ([]entities.TugasDosen, error) {
	var listTugas []entities.TugasDosen

	rows, err := config.DB.Query("SELECT TD.TD_Judul, TD.TD_Deadline FROM Tugas_Dosen TD INNER JOIN User_Kelas UK ON TD.Kelas_K_ID = UK.Kelas_K_ID INNER JOIN User U ON UK.User_U_ID = U.U_ID WHERE U.U_ID = ? AND MONTH(TD.TD_Deadline) = ? AND YEAR(TD.TD_Deadline) = ?", id, bulan, tahun)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No pertemuan found for the given id")
			return nil, err
		}
		log.Println("Error querying the database:", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var tugas entities.TugasDosen
		if err := rows.Scan(&tugas.TD_Judul, &tugas.TD_Deadline); err != nil {
			log.Println("Error scanning the database:", err)
			return nil, err
		}
		listTugas = append(listTugas, tugas)

	}
	return listTugas, nil
}
