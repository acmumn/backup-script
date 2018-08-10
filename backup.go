package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func backup(mysql MysqlConfig, dir, baseDir string) error {
	f, err := ioutil.TempFile("", "foo")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	defer os.Remove(f.Name())

	err = f.Chmod(0400)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(f, "[client]\nuser=%s\npassword=%q\n", mysql.User, mysql.Pass)
	if err != nil {
		return err
	}

	if mysql.Host != "" {
		_, err = fmt.Fprintf(f, "host=%s\n", mysql.Host)
		if err != nil {
			return err
		}
	}
	if mysql.Port != 0 {
		_, err = fmt.Fprintf(f, "port=%d\n", mysql.Port)
		if err != nil {
			return err
		}
	}

	err = f.Close()
	if err != nil {
		return err
	}

	args := []string{"--defaults-file=" + f.Name(), "--backup", "--target-dir", dir}
	if baseDir != "" {
		if IsDir(baseDir) {
			args = append(args, "--incremental-basedir", baseDir)
		} else {
			log.Printf("`%s' does not exist. Doing a full backup instead.", baseDir)
		}
	}

	cmd := exec.Command("mariabackup", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Wait()
}
