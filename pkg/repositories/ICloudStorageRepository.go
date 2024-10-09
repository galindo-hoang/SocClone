package repositories

import "github.com/PostService/pkg/repositories/models"

type ICloudStorageRepository interface {
	UploadFiles(entities []models.CloudEntity, bucket string) ([]*models.Images, error)
	//DeleteFiles()
	//DeleteFolder()
}
