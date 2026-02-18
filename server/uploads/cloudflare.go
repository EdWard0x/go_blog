package uploads

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"server/utils"
	"strings"
	"time"

	"io"
	"log"
	"server/global"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Cloudflare struct{}

func (c *Cloudflare) UploadImage(file *multipart.FileHeader) (string, string, error) {
	s3Client := InitS3()

	f, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	urlPre, filename, err := UploadFileS3(s3Client, file, f)
	if err != nil {
		return "", "", err
	}
	return urlPre + filename, filename, nil
}
func (c *Cloudflare) DeleteImage(key string) error {
	s3Client := InitS3()
	err := DeleteFileCf(s3Client, key)
	if err != nil {
		return err
	}
	return nil
}

func InitS3() *s3.Client {

	var accountId = global.Config.CF.AccountId
	var accessKeyId = global.Config.CF.AccessKeyId
	var accessKeySecret = global.Config.CF.AccessKeySecret
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"), // Required by SDK but not used by R2
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId))
	})

	return client
}

func ListObj(client *s3.Client) error {
	var bucketName = global.Config.CF.BucketName

	listObjectsOutput, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &bucketName,
	})
	if err != nil {
		return err
	}

	for _, object := range listObjectsOutput.Contents {
		obj, _ := json.MarshalIndent(object, "", "\t")
		fmt.Println(string(obj))
	}
	return nil
}

func UploadFileS3(client *s3.Client, file *multipart.FileHeader, reader io.Reader) (string, string, error) {
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		return "", "", fmt.Errorf("the image size exceeds the set size, the current size is: %.2f MB, the set size is: %d MB", size, global.Config.Upload.Size)

	}

	ext := filepath.Ext(file.Filename)
	name := strings.TrimSuffix(file.Filename, ext)
	if _, exists := WhiteImageList[ext]; !exists {
		return "", "", errors.New("don't upload files that aren't image types")
	}

	filename := utils.MD5V([]byte(name)) + "-" + time.Now().Format("20060102150405") + ext

	var bucketName = global.Config.CF.BucketName
	_, err := client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		//ChecksumAlgorithm: types.ChecksumAlgorithmCrc32,
		//Body:              strings.NewReader("Hello World"),
		Body: reader,
	})
	if err != nil {
		return "", "", err
	}
	return "https://678645.xyz/", filename, nil
}

func DeleteFileCf(client *s3.Client, name string) error {
	var bucketName = global.Config.CF.BucketName
	_, err := client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(name),
	})
	if err != nil {
		return err
	}
	return nil
}
