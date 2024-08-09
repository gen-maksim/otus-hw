package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("testRunCmd", func(t *testing.T) {
		args := []string{"ls"}
		os.Setenv("UNSET", "unset")
		require.Equal(t, "unset", os.Getenv("UNSET"))

		RunCmd(args, Environment{
			"BAR":   EnvValue{Value: "bar", NeedRemove: false},
			"UNSET": EnvValue{Value: "", NeedRemove: true},
		})

		require.Equal(t, "bar", os.Getenv("BAR"))
		require.Equal(t, "", os.Getenv("UNSET"))
	})
}
