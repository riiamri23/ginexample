package repositories

import (
	"errors"

	"example.com/ginexample/data/requests"
	"example.com/ginexample/helpers"
	"example.com/ginexample/model"
	"gorm.io/gorm"
)

type PlatformsRepository interface {
	FindAll() []model.Platforms
	FindById(id int) (platforms model.Platforms, err error)
	Create(platforms model.Platforms)
	Update(platforms model.Platforms)
	Delete(id int)
}

type PlatformsRepositoryImp struct {
	Db *gorm.DB
}

func NewPlatformsRepositoryImpl(Db *gorm.DB) PlatformsRepository {
	return &PlatformsRepositoryImp{Db: Db}
}

// Find All
func (t *PlatformsRepositoryImp) FindAll() []model.Platforms {
	var platforms []model.Platforms

	result := t.Db.Find(&platforms)

	helpers.ErrorPanic(result.Error)
	return platforms
}

// Find By Id
func (t *PlatformsRepositoryImp) FindById(id int) (platforms model.Platforms, err error) {
	var platform model.Platforms

	result := t.Db.Find(&platform, id)
	if result != nil {
		return platform, nil
	} else {
		return platform, errors.New("platform is not found")
	}
}

// Add
func (t *PlatformsRepositoryImp) Create(platforms model.Platforms) {
	result := t.Db.Create(&platforms)
	helpers.ErrorPanic(result.Error)
}

// Update
func (t *PlatformsRepositoryImp) Update(platforms model.Platforms) {
	var updatePlatform = requests.UpdatePlatformsRequest{
		Id:   platforms.Id,
		Name: platforms.Name,
	}
	result := t.Db.Model(&platforms).Updates(updatePlatform)
	helpers.ErrorPanic(result.Error)
}

// Delete
func (t *PlatformsRepositoryImp) Delete(id int) {
	var platform model.Platforms

	result := t.Db.Where("id = ?", id).Delete(&platform)
	helpers.ErrorPanic(result.Error)
}
