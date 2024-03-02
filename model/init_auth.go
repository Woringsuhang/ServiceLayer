package model

import "github.com/Woringsuhang/user/model"

func InitAutoMigrate() error {
	err := model.DB.AutoMigrate(&Users{})
	if err != nil {
		return err
	}
	return nil
}
