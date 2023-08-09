package command

import (
	"testing"
)

func TestMkdir(t *testing.T) {
	testCases := []commandCase{
		{"default permission", mkdir{Source: "./default-permission"}},
		{"custom permission", mkdir{Source: "./custom-permission", Perm: 0644}},
		{"custom permission with full perm", mkdir{Source: "./full-permission", Perm: 0777}},
		{"force custom permission with full perm", mkdir{Source: "./force-full-permission", Perm: 0777, MustPerm: 0777}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.cmd.Exec()
			if err != nil {
				t.Errorf("test mkdir command error => %v", err)
			}
		})
	}
}
