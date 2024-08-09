package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCmd(t *testing.T) {
	t.Run("testRunCmd", func(t *testing.T) {
		args := []string{"ls"}
		os.Setenv("UNSET", "unset")
		assert.Equal(t, "unset", os.Getenv("UNSET"))

		RunCmd(args, Environment{
			"BAR":   EnvValue{Value: "bar", NeedRemove: false},
			"UNSET": EnvValue{Value: "", NeedRemove: true},
		})

		assert.Equal(t, "bar", os.Getenv("BAR"))
		assert.Equal(t, "", os.Getenv("UNSET"))
	})
}
