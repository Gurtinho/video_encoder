package repositories_test

import (
	"testing"
	"time"

	"github.com/gurtinho/video_encoder/application/repositories"
	"github.com/gurtinho/video_encoder/domain"
	"github.com/gurtinho/video_encoder/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryDBInsert(t *testing.T) {
	db := database.NewDBTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	// injetando o db dentro do repo
	repo := repositories.VideoRepositoryDB{DB: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}
