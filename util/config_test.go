package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	path := ".././config"
	c, err := LoadConfig(path)
	require.NoError(t, err)
	require.NotEmpty(t, c)

}
