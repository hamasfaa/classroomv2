package usercontroller

import (
	"be/model/usermodel"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type loginResponse struct {
	Role  string `json:"role"`
	Error string `json:"error,omitempty"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(loginResponse{
			Error: "Masukkan Email dan Password!",
		})
		return
	}

	user, err := usermodel.GetUserByEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(loginResponse{
			Error: "Email atau Password Salah!",
		})
		return
	}
	fmt.Println(user)
	if !usermodel.VerifyPassword(user.U_Password, password) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(loginResponse{
			Error: "Email atau Password Salah!",
		})
		return
	}

	noPonsel := ""
	if user.U_NoPonsel.Valid {
		noPonsel = user.U_NoPonsel.String
	}

	alamat := ""
	if user.U_Alamat.Valid {
		alamat = user.U_Alamat.String
	}

	foto := ""
	if user.U_Foto.Valid {
		foto = user.U_Foto.String
	}

	user.U_NoPonsel = sql.NullString{String: noPonsel, Valid: true}
	user.U_Alamat = sql.NullString{String: alamat, Valid: true}
	user.U_Foto = sql.NullString{String: foto, Valid: true}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(loginResponse{
		Role: user.U_Role,
	})
}
