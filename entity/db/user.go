package db

import (
	"errors"
	"fmt"
	"go-standard/entity/base"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type SysUser struct {
	UserId     uint `gorm:"primaryKey"`
	Username   string
	Password   string
	Name       string
	Mobile     string
	Photo      string
	Email      string
	Status     uint8
	CreateBy   uint
	CreateTime *time.Time `gorm:"autoCreateTime"`
	UpdateBy   uint
	UpdateTime *time.Time `gorm:"autoUpdateTime"`
	Remarks    string
}

type SysUserDao interface {
	FindByUserId(uint) (*SysUser, error)
	// Save 保存用户
	// @return uint UserId
	// @return error
	Save(*SysUser) (uint, error)
}

func NewSysUserDao() SysUserDao {
	return &sysUserDaoImpl{base.DB}
}

type sysUserDaoImpl struct {
	DB *gorm.DB
}

func (d *sysUserDaoImpl) FindByUserId(userId uint) (*SysUser, error) {
	if userId == 0 {
		return nil, errors.New(fmt.Sprintf("userId must not equal %d", userId))
	}
	var sysUser SysUser
	result := d.DB.Limit(1).Find(&sysUser, &SysUser{UserId: userId})
	if result.Error != nil || result.RowsAffected <= 0 {
		return nil, result.Error
	}
	return &sysUser, nil
}

func (d *sysUserDaoImpl) Save(user *SysUser) (uint, error) {
	err := d.DB.Transaction(func(tx *gorm.DB) error {
		if result := tx.Create(user); result.Error != nil {
			return result.Error
		} else {
			if result.RowsAffected != 1 {
				return errors.New("unexpect rowsAffected " + strconv.Itoa(int(result.RowsAffected)))
			}
		}
		// 提交事务
		return nil
	})
	if err != nil {
		return 0, err
	}
	return user.UserId, nil
}
