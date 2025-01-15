package usercontroller

import (
	"be/model/usermodel"
	"be/session"
	"database/sql"
	"encoding/json"
	"net/http"
)

type loginResponse struct {
	Role  string `json:"role"`
	Error string `json:"error,omitempty"`
}

type registerResponse struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	sess, err := session.GetSession(r, "user")
	if err == nil && sess.Values["user_id"] != nil {
		role := sess.Values["user_role"].(string)
		http.Redirect(w, r, "/"+role, http.StatusSeeOther)
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

	user, err := usermodel.Login(email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(loginResponse{
			Error: "Email atau Password Salah!",
		})
		return
	}
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

	user.U_NoPonsel = sql.NullString{String: noPonsel, Valid: true}
	user.U_Alamat = sql.NullString{String: alamat, Valid: true}

	sess, err = session.GetSession(r, "user")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sess.Values["user_id"] = user.U_ID
	sess.Values["user_role"] = user.U_Role
	err = session.SaveSession(w, r, sess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(loginResponse{
		Role: user.U_Role,
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	sess, err := session.GetSession(r, "user")
	if err == nil && sess.Values["user_id"] != nil {
		role := sess.Values["user_role"].(string)
		http.Redirect(w, r, "/"+role, http.StatusSeeOther)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	dob := r.FormValue("dob")
	role := r.FormValue("role")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	if name == "" || dob == "" || role == "" || email == "" || password == "" || confirmPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(registerResponse{
			Error: "Semua kolom harus diisi!",
		})
		return
	} else {
		if password != confirmPassword {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(registerResponse{
				Error: "Password tidak sama!",
			})
			return
		} else {
			if usermodel.GetEmail(email) {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(registerResponse{
					Error: "Email sudah terdaftar!",
				})
				return
			} else {
				if err := usermodel.Register(name, dob, role, email, password); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(registerResponse{
						Error: "Gagal mendaftarkan pengguna!",
					})
					return
				}
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(registerResponse{
					Message: "Pendaftaran berhasil!",
				})
			}
		}
	}
}
