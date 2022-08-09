package db

import (
	"go-standard/entity/base"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestNewSysUserDao(t *testing.T) {
	tests := []struct {
		name string
		want SysUserDao
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSysUserDao(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSysUserDao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sysUserDaoImpl_FindByUserId(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		userId uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SysUser
		wantErr bool
	}{
		{
			"userNotExists", fields{base.DB}, args{uint(0)}, nil, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &sysUserDaoImpl{
				DB: tt.fields.DB,
			}
			got, err := d.FindByUserId(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sysUserDaoImpl_Save(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		user *SysUser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    func(uint) bool
		wantErr bool
	}{
		{
			"save", fields{base.DB}, args{&SysUser{Username: "test11"}}, func(userId uint) bool {
				return userId > 0
			}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sy := sysUserDaoImpl{
				DB: tt.fields.DB,
			}
			got, err := sy.Save(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want(got) {
				t.Errorf("Save() got = %v, want %v", got, tt.want(got))
			}
		})
	}
}
