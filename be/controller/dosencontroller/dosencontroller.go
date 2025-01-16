package dosencontroller

import (
	"be/model/dosenmodel"
	"be/session"
	"encoding/json"
	"net/http"
)

type indexResponse struct {
	Error string `json:"error,omitempty"`
	Foto  string `json:"foto,omitempty"`
	Nama  string `json:"nama,omitempty"`
	Role  string `json:"role,omitempty"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	sess, err := session.GetSession(r, "user")
	if err != nil || sess.Values["user_id"] == nil {
		json.NewEncoder(w).Encode(indexResponse{
			Error: "Tidak ada pengguna yang login",
		})
		return
	}

	userID := sess.Values["user_id"].(string)

	user, err := dosenmodel.Data(userID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(indexResponse{
			Error: "Terjadi kesalahan pada server",
		})
		return
	}

	json.NewEncoder(w).Encode(user)
}
