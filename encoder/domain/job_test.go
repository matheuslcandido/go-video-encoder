package domain_test

import (
	"encoder/domain"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "/aws/s3/converted"
	video.ResourceId = "external_id"

	job, err := domain.NewJob("/aws/s3/converted", "converted", video)

	require.NotNil(t, job)
	require.Nil(t, err)
}
