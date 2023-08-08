package command

import (
	"testing"
)

func TestIsSymlink(t *testing.T) {
	testCases := []commandCase{
		{"is not symlink", isSymlink{Link: "./is_symlink_test.go", Expect: false, IgnoreError: false}},
		{"is not exist", isSymlink{Link: "./not_exist", Expect: false, IgnoreError: true}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cmd.Exec()
			if err != nil {
				t.Errorf(testExecReturnErrorFailedMessage)
			}
		})
	}
}

func TestIsSymlink_ReturnError(t *testing.T) {
	testCases := []commandCase{
		{"is not symlink", isSymlink{Link: "./is_symlink_test.go", Expect: true, IgnoreError: false}},
		{"is not exist", isSymlink{Link: "./not_exist", Expect: false, IgnoreError: false}},
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
