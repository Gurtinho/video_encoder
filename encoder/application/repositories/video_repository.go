package repositories

import (
	"fmt"

	"github.com/gurtinho/video_encoder/domain"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type VideoRepository interface {
	Insert(video *domain.Video) (*domain.Video, error)
	Find(id string) (*domain.Video, error)
}

type VideoRepositoryDB struct {
	DB *gorm.DB
}

// instancia essa função passando a conexão, uma vez passada já é possivel sair usando a struct
func NewVideoRepository(db *gorm.DB) *VideoRepositoryDB {
	return &VideoRepositoryDB{
		DB: db,
	}
}

// métodos implementando o video repository DB
func (repo *VideoRepositoryDB) Insert(video *domain.Video) (*domain.Video, error) {
	if video.ID == "" {
		video.ID = uuid.NewV4().String()
	}

	err := repo.DB.Create(video).Error

	if err != nil {
		return nil, err
	}

	return video, nil
}

func (repo *VideoRepositoryDB) Find(id string) (*domain.Video, error) {
	var video domain.Video
	repo.DB.First(&video, "id = ?", id) // preenche o video

	if video.ID == "" {
		return nil, fmt.Errorf("video does not exist")
	}

	return &video, nil
}