package utils_test

import (
	"encoder/framework/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsJson(t *testing.T) {
	json := `{
		"id": "931b3c0b-d91a-4064-aad6-989fb05dffe5",
		"file_path": "convite.mp4",
		"status": "pending"
	}`

	err := utils.IsJson(json)
	require.Nil(t, err)

	json = "wes"

	err = utils.IsJson(json)
	require.Error(t, err)
}
