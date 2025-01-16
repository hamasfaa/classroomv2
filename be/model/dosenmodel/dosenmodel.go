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
