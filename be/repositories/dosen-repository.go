package repositories

import (
	"be/entities"

	"gorm.io/gorm"
)

type DosenRepository interface {
	GetAllClass(userUID string) ([]entities.Kelas, error)
	CreateClass(kelas *entities.Kelas) error
	AddUserToClass(userUID string, classID string) error
	DeleteClass(classID string) error
	GetAllTask(userUID string) ([]entities.TugasDosen, error)
	GetAllMeeting() ([]entities.AbsenDosen, error)
	// CreateTask(task *entities.TugasDosen) error
	CreateTaskWithFiles(task *entities.TugasDosen, files []entities.TugasFile) error
}

type dosenRepositoryGorm struct {
	db *gorm.DB
}

func NewDosenRepository(db *gorm.DB) *dosenRepositoryGorm {
	return &dosenRepositoryGorm{db: db}
}

func (r *dosenRepositoryGorm) GetAllClass(userUID string) ([]entities.Kelas, error) {
	var kelas []entities.Kelas

	query := `SELECT * FROM kelas K INNER JOIN user_kelas UK ON K.k_id = UK.kelas_k_id WHERE UK.user_uid = ?`

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

func (r *dosenRepositoryGorm) DeleteClass(classID string) error {
	if err := r.db.Exec("DELETE FROM user_kelas WHERE kelas_k_id = ?", classID).Error; err != nil {
		return err
	}

	if err := r.db.Exec("DELETE FROM kelas WHERE k_id = ?", classID).Error; err != nil {
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

// func (r *dosenRepositoryGorm) CreateTask(task *entities.TugasDosen) error {
// 	if err := r.db.Create(task).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

func (r *dosenRepositoryGorm) CreateTaskWithFiles(task *entities.TugasDosen, files []entities.TugasFile) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(task).Error; err != nil {
			return err
		}

		if len(files) > 0 {
			if err := tx.Create(&files).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
