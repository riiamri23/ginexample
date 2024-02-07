package repositories

import (
	"errors"

	"example.com/ginexample/model"
	"github.com/davecgh/go-spew/spew"
	"gorm.io/gorm"
)

type AuthorizationRepository interface {
	FindUser(username string, password string) (users model.Users, err error)
}

type AuthorizationRepositoryImp struct {
	Db *gorm.DB
}

func NewAuthorizationRepositoryImpl(Db *gorm.DB) AuthorizationRepository {
	return &AuthorizationRepositoryImp{Db: Db}
}

// Find User
func (t *AuthorizationRepositoryImp) FindUser(username string, password string) (users model.Users, err error) {
	var user model.Users

	result := t.Db.Where("username = ?", username).Where("password = ?", password).Find(&user)
	// result := t.Db.Where("username = ?", username).Where("password = ?", password).First(&user).Error

	spew.Dump(user)

	if result != nil {
		return user, nil
	} else {
		return user, errors.New("username or password is wrong")
	}

}
