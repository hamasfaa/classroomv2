package repositories

import (
	"be/entities"

	"gorm.io/gorm"
)

type DosenRepository interface {
	GetAllClass(userUID string) ([]entities.Kelas, error)
	CreateClass(kelas *entities.Kelas) error
	AddUserToClass(userUID string, classID string) error
	GetAllTask(userUID string) ([]entities.TugasDosen, error)
	GetAllMeeting() ([]entities.AbsenDosen, error)
}

type dosenRepositoryGorm struct {
	db *gorm.DB
}

func NewDosenRepository(db *gorm.DB) *dosenRepositoryGorm {
	return &dosenRepositoryGorm{db: db}
}

func (r *dosenRepositoryGorm) GetAllClass(userUID string) ([]entities.Kelas, error) {
	var kelas []entities.Kelas

	query := `SELECT * FROM kelas K INNER JOIN user_kelas UK ON K.k_id = UK.kelas_k_id WHERE UK.user_u_id = ?`

	if err := r.db.Raw(query, userUID).Scan(&kelas).Error; err != nil {
		return nil, err
	}

	return kelas, nil
}

func (r *dosenRepositoryGorm) CreateClass(kelas *entities.Kelas) error {
	if err := r.db.Create(kelas).Error; err != nil {
		return err
	}

	return nil
}

func (r *dosenRepositoryGorm) AddUserToClass(userUID string, classID string) error {
	query := `INSERT INTO user_kelas (user_uid, kelas_k_id) VALUES (?, ?)`

	if err := r.db.Exec(query, userUID, classID).Error; err != nil {
		return err
	}

	return nil
}

func (r *dosenRepositoryGorm) GetAllTask(userUID string) ([]entities.TugasDosen, error) {
	var tugas []entities.TugasDosen

	if err := r.db.Find(&tugas).Error; err != nil {
		return nil, err
	}

	return tugas, nil
}

func (r *dosenRepositoryGorm) GetAllMeeting() ([]entities.AbsenDosen, error) {
	var absen []entities.AbsenDosen

	if err := r.db.Find(&absen).Error; err != nil {
		return nil, err
	}

	return absen, nil
}
