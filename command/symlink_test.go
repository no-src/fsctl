package command

import (
	"os"
	"testing"
)

func TestSymlink(t *testing.T) {
	workspace := "./symlink_workspace"
	os.MkdirAll(workspace, os.ModePerm)
	defer os.RemoveAll(workspace)
	testCases := []commandCase{
		{"ignore symlink error", symlink{Link: workspace + "/symlink_test.go.symlink", Dest: "./symlink_test.go", IgnoreError: true}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cmd.Exec()
			if err != nil {
				t.Errorf("test symlink command error => %v", err)
			}
		})
	}
}

func TestSymlink_ReturnError(t *testing.T) {
	testCases := []commandCase{
		{"file exist error", symlink{Link: "./symlink_test.go", Dest: "./symlink_test.go", IgnoreError: false}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cmd.Exec()
			if err == nil {
				t.Errorf(testExecReturnErrorFailedMessage)
			}
		})
	}
}
