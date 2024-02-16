package domain_test

import (
	"testing"
	"time"

	"github.com/gurtinho/video_encoder/domain"
	"github.com/stretchr/testify/require"
	uuid "github.com/satori/go.uuid"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path", "convertido", video)
	require.NotNil(t, job)
	require.Nil(t, err)
}