package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	path := "../."
	c, err := LoadConfig(path)
	require.NoError(t, err)
	require.NotEmpty(t, c)

}
