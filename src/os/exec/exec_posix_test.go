// +build !plan9,!windows,!nacl

package exec_test

import (
	"os/user"
	"strconv"
	"syscall"
	"testing"
)

func TestCredential(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		t.Fatalf("error getting current user: %v", err)
	}

	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		t.Fatalf("error converting Uid to integer: %v", err)
	}

	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		t.Fatalf("error converting Gid to integer: %v", err)
	}

	cmd := helperCommand(t, "echo", "foo bar", "baz")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Credential: &syscall.Credential{
			Uid: uint32(uid),
			Gid: uint32(gid),
		},
	}

	if err = cmd.Run(); err != nil {
		t.Errorf("echo: %v", err)
	}
}
