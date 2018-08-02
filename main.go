package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPass := os.Getenv("MYSQL_PASSWORD")
	secretKey := os.Getenv("SECRET_KEY")
	s3Bucket := os.Getenv("S3_BUCKET")

	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)
}

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

	err := cmd.Wait()
	if err != nil {
		return "", err
	}
}
