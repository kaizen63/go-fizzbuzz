package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	if os.Getenv("IN_TESTMAIN") == "1" {
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestMain")
	cmd.Env = append(os.Environ(), "IN_TESTMAIN=1")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("process ran with err %v, want exit status 0", err)
	}
}

func TestMain2(t *testing.T) {
	main()
	return
}
