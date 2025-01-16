package dosencontroller

import (
	"be/entities"
	"be/model/dosenmodel"
	"be/session"
	"encoding/json"
	"net/http"
	"time"
)

type indexResponse struct {
	Error         string                `json:"error,omitempty"`
	User          *entities.User        `json:"user,omitempty"`
	ListKelas     []entities.Kelas      `json:"list_kelas,omitempty"`
	ListPertemuan []entities.AbsenDosen `json:"list_pertemuan,omitempty"`
	ListTugas     []entities.TugasDosen `json:"list_tugas,omitempty"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	t := time.Now()
	bulanString := t.Month()
	bulan := int(bulanString)
	tahun := t.Year()

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

	listKelas, err := dosenmodel.ListKelas(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(indexResponse{
			Error: "Terjadi kesalahan pada server",
		})
		return
	}

	listPertemuan, err := dosenmodel.ListPertemuan(userID, bulan, tahun)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(indexResponse{
			Error: "Terjadi kesalahan pada server",
		})
		return
	}

	listTugas, err := dosenmodel.ListTugas(userID, bulan, tahun)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(indexResponse{
			Error: "Terjadi kesalahan pada server",
		})
		return
	}

	json.NewEncoder(w).Encode(indexResponse{
		User:          user,
		ListKelas:     listKelas,
		ListPertemuan: listPertemuan,
		ListTugas:     listTugas,
	})
}
