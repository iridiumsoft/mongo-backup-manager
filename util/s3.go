package util

import (
	// Core Packages
	"fmt"
	"bytes"

	// Local Packages
	"github.com/sirupsen/logrus"
	"github.com/iridiumsoft/mongo-backup-manager/config"

	// Third Party Packages
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"os"
	"net/http"
	"io"
)

type S3Service struct {
	Config config.Main
}

func (s S3Service) New(c config.Main) *S3Service {
	s.Config = c
	return &s
}

func (s *S3Service) GetObject(Key string, filePath string) (err error) {
	logrus.Println("Key", Key)
	svc := s.GetS3Session();
	object, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.Config.AWS.Bucket),
		Key:    aws.String(Key),
	})
	if err == nil {
		outFile, _ := os.Create(filePath)
		// handle err
		defer outFile.Close()
		_, err = io.Copy(outFile, object.Body)
	}
	return
}

func (s *S3Service) PutObject(FilePath string, Key string, ACL string) {
	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Printf("err opening file: %s", err)
		return
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	// Get S3 Session
	svc := s.GetS3Session();
	_, err = svc.PutObject(&s3.PutObjectInput{
		Body:                 fileBytes,
		Bucket:               aws.String(s.Config.AWS.Bucket),
		Key:                  aws.String(Key),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(fileType),
		ACL:                  aws.String(ACL),
		ServerSideEncryption: aws.String("AES256"),
	})

	if err != nil {
		logrus.Error("Error while putObject: ", Key, " : ", err.Error())
		return
	}
}

func (s *S3Service) GetS3Session() (svc *s3.S3) {
	creds := credentials.NewStaticCredentials(s.Config.AWS.AccessKey, s.Config.AWS.SecreteKey, "")
	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
	}
	cfg := aws.NewConfig().WithRegion(s.Config.AWS.Region).WithCredentials(creds)
	svc = s3.New(session.New(), cfg)
	return svc
}

func (s *S3Service) DeleteObject() {

}
