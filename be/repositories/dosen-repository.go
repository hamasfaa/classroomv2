package repositories

import (
	"be/entities"

	"gorm.io/gorm"
)

type DosenRepository interface {
	GetAllClass(userUID string) ([]entities.Kelas, error)
	CreateClass(kelas *entities.Kelas) error
	AddUserToClass(userUID string, classID string) error
	GetAllUserInClass(classID string) ([]entities.User, error)
	GetDetailClass(classID string) (entities.Kelas, error)
	DeleteClass(classID string) error
	GetAllTask(userUID string, classID string) ([]entities.TugasDosen, error)
	GetTaskByID(taskID string) (entities.TugasDosen, error)
	GetAllMeeting() ([]entities.AbsenDosen, error)
	CreateTaskWithFiles(task *entities.TugasDosen, files []entities.TugasFile) error
	DeleteTaskWithFiles(taskID string) error
	UpdateStatusTask(taskID string, status bool) error
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

func (r *dosenRepositoryGorm) GetAllUserInClass(classID string) ([]entities.User, error) {
	var users []entities.User

	query := `SELECT * FROM users U INNER JOIN user_kelas UK ON U.uid = UK.user_uid WHERE UK.kelas_k_id = ?`
	if err := r.db.Raw(query, classID).Scan(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *dosenRepositoryGorm) GetDetailClass(classID string) (entities.Kelas, error) {
	var kelas entities.Kelas

	if err := r.db.Where("k_id = ?", classID).First(&kelas).Error; err != nil {
		return entities.Kelas{}, err
	}

	return kelas, nil
}

func (r *dosenRepositoryGorm) DeleteClass(classID string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`
            DELETE FROM tugas_files 
            WHERE tugas_td_id IN (
                SELECT td_id 
                FROM tugas_dosens 
                WHERE kelas_k_id = ?
            )`, classID).Error; err != nil {
			return err
		}

		if err := tx.Exec("DELETE FROM tugas_dosens WHERE kelas_k_id = ?", classID).Error; err != nil {
			return err
		}

		if err := tx.Exec("DELETE FROM user_kelas WHERE kelas_k_id = ?", classID).Error; err != nil {
			return err
		}

		if err := tx.Exec("DELETE FROM kelas WHERE k_id = ?", classID).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *dosenRepositoryGorm) GetAllTask(userUID string, classUID string) ([]entities.TugasDosen, error) {
	var tugas []entities.TugasDosen

	query := `SELECT * FROM tugas_dosens WHERE kelas_k_id = ? AND user_uid = ?`

	if err := r.db.Raw(query, classUID, userUID).Scan(&tugas).Error; err != nil {
		return nil, err
	}

	return tugas, nil
}

func (r *dosenRepositoryGorm) GetTaskByID(taskID string) (entities.TugasDosen, error) {
	var tugas entities.TugasDosen

	if err := r.db.Where("td_id = ?", taskID).First(&tugas).Error; err != nil {
		return entities.TugasDosen{}, err
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

//not safety without transaction
// func (r *dosenRepositoryGorm) CreateTaskWithFiles(task *entities.TugasDosen, files []entities.TugasFile) error {
// 	// If this succeeds but files creation fails
// 	if err := r.db.Create(task).Error; err != nil {
// 		return err
// 	}
// 	// If this fails, we'll have a task without its files
// 	if len(files) > 0 {
// 		if err := r.db.Create(&files).Error; err != nil {
// 			return err
// 		}
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

func (r *dosenRepositoryGorm) DeleteTaskWithFiles(taskID string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM tugas_files WHERE tugas_td_id = ?", taskID).Error; err != nil {
			return err
		}

		if err := tx.Exec("DELETE FROM tugas_dosens WHERE td_id = ?", taskID).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *dosenRepositoryGorm) UpdateStatusTask(taskID string, status bool) error {
	if err := r.db.Exec("UPDATE tugas_dosens SET td_status = ? WHERE td_id = ?", status, taskID).Error; err != nil {
		return err
	}

	return nil
}
