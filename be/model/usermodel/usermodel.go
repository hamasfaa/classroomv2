package usermodel

import (
	"be/config"
	"be/entities"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := config.DB.QueryRow("SELECT * FROM User WHERE U_Email = ?", email).Scan(&user.U_ID, &user.U_Nama, &user.U_Email, &user.U_Password, &user.U_Role, &user.U_TanggalLahir, &user.U_NoPonsel, &user.U_Alamat, &user.U_Foto)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No user found with the given email")
			return nil, err
		}
		log.Println("Error querying the database:", err)
		return nil, err
	}
	return &user, nil
}

func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
