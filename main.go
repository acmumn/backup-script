package main

import "log"

func main() {
	mysql, aws, schedule := LoadConfig("backup.toml")
	_, _ = aws, schedule

	if err := backup(mysql, "backup", "basedir"); err != nil {
		log.Panicln(err)
	}

	//secretKey := os.Getenv("SECRET_KEY")
	//s3Bucket := os.Getenv("S3_BUCKET")

	//sess := session.Must(session.NewSession())
	//uploader := s3manager.NewUploader(sess)

	//// TODO
}
