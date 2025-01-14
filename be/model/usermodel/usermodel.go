package usermodel

import (
	"be/config"
	"be/entities"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Login(email string) (*entities.User, error) {
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

func getLastID(role string) string {
	var maxID sql.NullString

	err := config.DB.QueryRow("SELECT MAX(U_ID) AS max_id FROM User WHERE U_Role = ?", role).Scan(&maxID)
	if err != nil {
		log.Println("Error querying the database:", err)
		return ""
	}

	prefixID := 50
	if role == "dosen" {
		prefixID = 10
	}

	newID := ""
	if maxID.Valid && maxID.String != "" {
		numPart, _ := strconv.Atoi(maxID.String[2:])
		numPart++
		newID = fmt.Sprintf("%d%07d", prefixID, numPart)
	} else {
		newID = fmt.Sprintf("%d0000001", prefixID)
	}
	return newID
}

func Register(name, dob, role, email, password string) error {
	var user entities.User

	user.U_ID = getLastID(role)
	user.U_Nama = name
	user.U_Email = email
	user.U_Role = role
	user.U_Foto = "default.jpg"
	user.U_Password = password

	dobTime, err := time.Parse("2006-01-02", dob)
	if err != nil {
		log.Println("Error parsing the dob:", err)
		return err
	}
	user.U_TanggalLahir = dobTime

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.U_Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing the password:", err)
		return err
	}

	_, err = config.DB.Exec(
		"INSERT INTO User (U_ID, U_Nama, U_Email, U_Password, U_Role, U_TanggalLahir, U_NoPonsel, U_Alamat, U_Foto) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		user.U_ID, user.U_Nama, user.U_Email, hashedPassword, user.U_Role, user.U_TanggalLahir, user.U_NoPonsel, user.U_Alamat, user.U_Foto,
	)

	if err != nil {
		log.Println("Error inserting the user into the database:", err)
		return err
	}
	return nil
}

func GetEmail(email string) bool {
	var userID string

	err := config.DB.QueryRow("SELECT U_ID FROM User WHERE U_Email = ?", email).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Println("Error querying the database:", err)
		return false
	}

	return true
}
