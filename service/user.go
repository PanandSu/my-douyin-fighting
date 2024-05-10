package service

import (
	"errors"
	"my-douyin-fighting/glob"
	"my-douyin-fighting/model"
)

func Login(username, password string) (*model.User, error) {
	var (
		user *model.User
		err  error
	)
	result := glob.DB.Where("name=?", username).Limit(1).Find(user)
	if result.RowsAffected == 0 {
		err = errors.New("")
		return nil, err
	}
	if user.Password != password {
		err = errors.New("")
		return nil, err
	}
	return user, nil
}

func Register(username, password string) (*model.User, error) {
	var (
		user *model.User
		err  error
	)
	result := glob.DB.Where("name=?", username).Limit(1).Find(user)
	if result.RowsAffected == 0 {
		err = errors.New("")
		return nil, err
	}
	user.Name = username
	user.Password = password //要加密
	uid, _ := glob.IdGenerator()
	err = glob.DB.Create(user).Error
	return user, err
}
