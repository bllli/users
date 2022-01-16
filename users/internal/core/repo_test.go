package core

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestUserRepo(t *testing.T) {
	dsn := "root:q123q123@tcp(127.0.0.1:3306)/users?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		t.Error(err)
	}
	user := UserModel{
		UserId:   1,
		Nickname: "test",
	}

	repo := NewUserRepo(context.Background(), db)
	err = repo.SaveUser(&user)
	if err != nil {
		t.Error(err)
	}

	gotUser, err := repo.GetUserByUserId(1)
	if err != nil {
		t.Error(err)
	}

	gotUser.WechatAuth = WechatAuthModel{
		OpenId:     "test0",
		UnionId:    "test1",
		SessionKey: "test2",
	}
	err = repo.SaveUser(gotUser)
	if err != nil {
		t.Error(err)
	}

	got2User, err := repo.GetUserByOpenId("test0")
	if err != nil {
		t.Error(err)
	}
	if got2User.UserId != 1 {
		t.Error("user id not match")
	}

}
