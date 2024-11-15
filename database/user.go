package database

import (
	"eshop_user/log"
	"eshop_user/model"
	"gorm.io/gorm"
)

func InsertOneUser(db *gorm.DB, user *model.User) (err error) {
	if db == nil {
		db, err = GetDB()
		if err != nil {
			log.Errorf("error: %v", err)
			return err
		}
	}
	err = db.Table("user").Create(&user).Error
	if err != nil {
		log.Errorf("error: %v", err)
	}
	return err
}

func QueryOneUserById(db *gorm.DB, userId string) (user *model.User, err error) {
	if db == nil {
		db, err = GetDB()
		if err != nil {
			log.Errorf("error: %v", err)
			return nil, err
		}
	}
	var tmp []*model.User
	err = db.Table("user").Where("uid = ?", userId).Find(&tmp).Error
	if err != nil {
		log.Errorf("error: %v", err)
	}
	if len(tmp) == 0 {
		return nil, err
	} else {
		user = tmp[0]
	}
	return user, err
}

func QueryOneUserByName(db *gorm.DB, Name string) (user *model.User, err error) {
	if db == nil {
		db, err = GetDB()
		if err != nil {
			log.Errorf("error: %v", err)
			return nil, err
		}
	}
	var tmp []*model.User
	err = db.Table("user").Where("name = ?", Name).Find(&tmp).Error
	if err != nil {
		log.Errorf("error: %v", err)
	}
	if len(tmp) == 0 {
		return nil, err
	} else {
		user = tmp[0]
	}
	return user, err
}
