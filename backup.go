package main

import "os/exec"

func backup(user, pass string) (string, error) {
	cmd := exec.Command("mariabackup", "--defaults-file", "/dev/stdin", "--backup", "--target-dir", dir)
	err := cmd.Start()
	if err != nil {
		return "", err
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	// TODO

	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	return dir, err
}
