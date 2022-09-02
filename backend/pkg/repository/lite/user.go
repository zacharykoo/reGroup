package lite

import (
	"github.com/zacharykoo/reGroup/backend/pkg/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func GetUserRepository(db *gorm.DB) userRepository {
	return userRepository{
		db: db,
	}
}

func (c *userRepository) Get(ID int) ([]model.User, error) {
	var users []model.User
	option := func(db *gorm.DB, ID int) *gorm.DB {
		if ID != 0 {
			return db.Where("userID = ?", ID)
		}
		return db
	}
	err := option(c.db, ID).Find(&users).Error
	return users, err
}

func (c *userRepository) Create(user model.User) (model.User, error) {
	c.db.Save(&user)
	user.UserID = user.ID
	err := c.db.Updates(&user).Error
	return user, err
}

func (c *userRepository) Edit(user model.User) (model.User, error) {
	err := c.db.Where("userID = ?", user.UserID).Updates(&user).Error
	return user, err
}
