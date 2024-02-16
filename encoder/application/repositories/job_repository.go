package repositories

import (
	"github.com/gurtinho/video_encoder/domain"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type JobInterface interface {
	Insert(job *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
	Update(job *domain.Job) (*domain.Job, error)
}

type JobRepositoryDB struct {
	DB *gorm.DB
}

func (repo *JobRepositoryDB) Insert(job *domain.Job) (*domain.Job, error) {
	if job.ID == "" {
		job.ID = uuid.NewV4().String()
	}

	err := repo.DB.Create(job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}
