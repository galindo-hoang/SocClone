package impl

import (
	"context"
	"github.com/PostService/pkg/database"
	"github.com/PostService/pkg/repositories/models"
	"github.com/PostService/pkg/utils"
	"github.com/minio/minio-go"
	"time"
)

type CloudStorageRepository struct {
}

func NewCloudStorageRepository() *CloudStorageRepository {
	return &CloudStorageRepository{}
}

func getPresignedURL(client *minio.Client, bucketName string, entity models.CloudEntity) (string, error) {
	expiry := time.Second * 24 * 60 * 60 // 1 day.
	presignedURL, err := client.PresignedPutObject(bucketName, entity.ObjectName, expiry)
	if err != nil {
		return "", nil
	}
	return presignedURL.String(), nil
}

func (s *CloudStorageRepository) UploadFiles(entities []models.CloudEntity, bucket string) ([]*models.Images, error) {
	var client, err = database.InitMinio(bucket)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	images := make([]*models.Images, 0)
	for _, entity := range entities {
		_, err := client.FPutObjectWithContext(
			ctx,
			bucket,
			entity.ObjectName,
			entity.FilePath,
			minio.PutObjectOptions{ContentType: entity.ContentType})
		if err != nil {
			isSuccess := false
			for i := 0; i < entity.Retry; i++ {
				_, err := client.FPutObjectWithContext(
					ctx,
					bucket,
					entity.ObjectName,
					entity.FilePath,
					minio.PutObjectOptions{ContentType: entity.ContentType})
				if err != nil {
					panic(err)
				} else {
					url := utils.GetPathFromMinio(bucket, entity.ObjectName)
					if err != nil {
						panic(err)
					} else {
						images = append(images, &models.Images{S3: url})
						isSuccess = true
						break
					}
				}
			}

			if !isSuccess {
				return nil, err
			}
		} else {
			url := utils.GetPathFromMinio(bucket, entity.ObjectName)
			if err != nil {
				panic(err)
			} else {
				images = append(images, &models.Images{S3: url})
			}
		}
	}
	return images, nil
}

//func (s *CloudStorageRepository) GetFilesFrom(entities []*models.CloudEntity, bucket string) error {}
