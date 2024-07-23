package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("simple copy", func(t *testing.T) {
		err := Copy("testdata/input.txt", "testdata/output.txt", 0, 0, false)
		if err != nil {
			require.NoError(t, err, err.Error())
			return
		}
		defer os.Remove("testdata/output.txt")
		in, err := readFile("testdata/input.txt")
		if err != nil {
			require.NoError(t, err, err.Error())
			return
		}
		out, err := readFile("testdata/output.txt")
		if err != nil {
			require.NoError(t, err, err.Error())
			return
		}
		require.Equal(t, in, out)
	})

	t.Run("limit and offset", func(t *testing.T) {
		limit := 1000
		err := Copy("testdata/input.txt", "testdata/output.txt", 100, limit, false)
		if err != nil {
			require.NoError(t, err, err.Error())
			return
		}
		defer os.Remove("testdata/output.txt")
		out, err := readFile("testdata/output.txt")
		if err != nil {
			require.NoError(t, err, err.Error())
			return
		}

		require.Equal(t, limit, len(out), "Suceess")
	})

	t.Run("test big offset", func(t *testing.T) {
		offset := 6000
		err := Copy("testdata/input.txt", "testdata/output.txt", offset, 1000, false)
		if err != nil {
			require.NoError(t, err, err.Error())
			return
		}
		defer os.Remove("testdata/output.txt")
		in, err := readFile("testdata/input.txt")
		if err != nil {
			require.NoError(t, err, err.Error())
			return
		}

		out, err := readFile("testdata/output.txt")
		if err != nil {
			require.NoError(t, err, err.Error())
			return
		}

		require.Equal(t, len(in)-offset, len(out))
	})

	t.Run("test too big offset", func(t *testing.T) {
		offset := 8000

		err := Copy("testdata/input.txt", "testdata/output.txt", offset, 1000, false)
		require.Error(t, err, "Error should be returned")
		require.ErrorContains(t, err, ErrOffsetExceedsFileSize.Error())
	})
}

func readFile(name string) ([]byte, error) {
	content, _, err := ReadFile(name, 0, 0, false)
	if err != nil {
		return nil, err
	}
	return content, nil
}
