package repositories

import "github.com/jrolstad/ldap-api/internal/pkg/models"

type DirectoryObjectRepository interface {
	Save(item *models.DirectoryObject) error
	Destroy()
}

func NewDirectoryObjectRepository() DirectoryObjectRepository {
	instance := &S3DirectoryObjectRepository{}
	instance.init()

	return instance
}
