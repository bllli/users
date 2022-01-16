package core

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	UserId    int64     `gorm:"column:user_id;primaryKey"`
	Nickname  string    `gorm:"column:nickname"`
	Status    int       `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	WechatAuth WechatAuthModel `gorm:"foreignkey:user_id"`
}

func (u *UserModel) TableName() string {
	return "users"
}

type WechatAuthModel struct {
	UserId     int64     `gorm:"column:user_id;primaryKey"`
	OpenId     string    `gorm:"column:openid"`
	UnionId    string    `gorm:"column:union_id"`
	SessionKey string    `gorm:"column:session_key"`
	CreatedAt  time.Time `gorm:"column:created_at;<-:create"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (w *WechatAuthModel) TableName() string {
	return "wechat_auth"
}

type UserRepo interface {
	GetUserByOpenId(openId string) (*UserModel, error)
	GetUserByUserId(userId int64) (*UserModel, error)
	SaveUser(user *UserModel) error
}

type UserRepoImpl struct {
	logx.Logger
	db *gorm.DB
}

func NewUserRepo(context context.Context, db *gorm.DB) UserRepo {
	return &UserRepoImpl{
		Logger: logx.WithContext(context),
		db:     db,
	}
}

func (u *UserRepoImpl) GetUserByUserId(userId int64) (*UserModel, error) {
	var user UserModel
	if err := u.db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepoImpl) GetUserByOpenId(openId string) (*UserModel, error) {
	var wechatAuth WechatAuthModel
	result := u.db.Model(&wechatAuth).Where("openid = ?", openId).First(&wechatAuth)
	if result.Error != nil {
		return nil, result.Error
	}
	return u.GetUserByUserId(wechatAuth.UserId)
}

func (u *UserRepoImpl) SaveUser(user *UserModel) error {
	if err := u.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(user).Error; err != nil {
		return err
	}
	return nil
}
