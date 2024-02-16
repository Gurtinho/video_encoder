package domain_test

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/gurtinho/video_encoder/domain"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()
	require.Error(t, err)
}

func TestVideoIdIsNotAUuid(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "aaaa"
	video.ResourceID = "bbbb"
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	err := video.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "bbbb"
	video.FilePath = "path"
	video.CreatedAt = time.Now()
	err := video.Validate()
	require.Nil(t, err)
}