package domain

import (
	"github.com/asaskevich/govalidator"
	"time"
)

// função do go que é executada antes de tudo
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// entidade video
type Video struct {
	ID         string    `json:"encoder_video_folder" valid:"uuid" gorm:"type:uuid;primary_key"` // já passa direto o caminho da pasta
	ResourceID string    `json:"resource_id" valid:"notnull" gorm:"type:varchar(255)"` // ID de fora
	FilePath   string    `json:"file_path" valid:"notnull" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `json:"-" valid:"-"`
	Jobs       []*Job    `json:"-" valid:"-" gorm:"ForeignKey:videoId"` // slice de jobs
}

func NewVideo() *Video {
	return &Video{}
}

// função implementada na struct
func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video)
	if err != nil {
		return err
	}
	return nil
}
