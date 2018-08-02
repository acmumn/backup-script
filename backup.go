package main

import (
	"fmt"
	"os/exec"
)

func backup(user, pass, dir, baseDir string) (string, error) {
	args := []string{
		"--defaults-file", "/dev/stdin", "--backup", "--target-dir", dir,
	}
	if baseDir != "" {
		args = append(args, "--incremental-basedir", baseDir)
	}
	cmd := exec.Command("mariabackup", *args)

	err := cmd.Start()
	if err != nil {
		return err
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(stdin, "[client]\nusername = %s\npassword = %q\n", user, pass)
	if err != nil {
		return err
	}

	err = stdin.Close()
	if err != nil {
		return err
	}

	return cmd.Wait()
}
