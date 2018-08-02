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
	err = backup(mysqlUser, mysqlPass)
	if err != nil {
		log.Panicln(err)
	}

	secretKey := os.Getenv("SECRET_KEY")
	s3Bucket := os.Getenv("S3_BUCKET")

	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)

	// TODO
}
