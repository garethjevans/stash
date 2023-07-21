package cmd_test

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/garethjevans/stash/pkg/cmd"

	"github.com/stretchr/testify/assert"
)

func TestStashCmd_Run(t *testing.T) {
	tests := []struct {
		name          string
		workingDir    string
		includes      string
		expectedFiles []string
	}{
		{
			name:          "stash01",
			workingDir:    filepath.Join("testdata"),
			includes:      "**",
			expectedFiles: []string{"file01.txt"},
		},
	}

	u := cmd.NewStashCmd()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmd.Name = tc.name

			stdout := bytes.NewBufferString("")
			u.SetOut(stdout)

			stderr := bytes.NewBufferString("")
			u.SetErr(stderr)

			err := u.Execute()
			assert.NoError(t, err)

			//actualOut, err := io.ReadAll(stdout)
			//assert.NoError(t, err)
			//assert.Equal(t, tc.expected, string(actualOut))
			//
			//actualErr, err := io.ReadAll(stderr)
			//assert.NoError(t, err)
			//assert.Equal(t, tc.expectedErrMessage, strings.TrimSpace(string(actualErr)))
		})
	}
}
