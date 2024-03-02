package model

import (
	"github.com/Woringsuhang/user/model"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `gorm:"index"`
	Password string `gorm:"type:varchar(100)"`
	Mobile   string `gorm:"type:char(11)"`
	Sex      int    `gorm:"type:tinyint(1)"`
	Age      int    `gorm:"type:tinyint(3)"`
	Address  string `gorm:"type:varchar(1024)"`
}

func Get(id int64) (*Users, error) {
	userInfo := new(Users)
	err := model.DB.Where("id", id).First(userInfo).Error
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func GetMobileUser(mobile string) (*Users, error) {
	userInfo := new(Users)
	err := model.DB.Where("mobile", mobile).First(userInfo).Error
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func Create(user *Users) (*Users, error) {
	err := model.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int64) error {
	userInfo := new(Users)
	return model.DB.Where("id = ?", id).Delete(userInfo).Error
}

func Update(id int64, user *Users) (*Users, error) {
	err := model.DB.Where("id = ?", id).Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
