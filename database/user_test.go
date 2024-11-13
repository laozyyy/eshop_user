package database

import (
	"eshop_user/model"
	"fmt"
	"testing"
	"time"
)

func TestDB(t *testing.T) {
	InsertOneUser(nil, model.User{
		UID:        "12332132",
		Name:       "tset",
		Phone:      "123123213123",
		Email:      "1233123@13213.com",
		Password:   "akjgjfnasfkasjfna",
		Role:       0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		IsDeleted:  false,
	})
	user, _ := QueryOneUserById(nil, "12332132")
	fmt.Print(user)
}
