package domain_test

import (
	"encoder/domain"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIDIsNotUUID(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "test"
	video.FilePath = "/aws/s3/test"
	video.ResourceId = "external_id"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = "/aws/s3/test"
	video.ResourceId = "external_id"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}