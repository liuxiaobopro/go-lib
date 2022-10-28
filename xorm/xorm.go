package xorm

import (
	"github.com/liuxiaobopro/go-lib/xorm/models"
)

type Model struct {
	CreateTime   models.LocalTime `json:"create_time" xorm:"default '0000-00-00 00:00:00' comment('创建时间') datetime 'create_time' created"`
	UpdateTime   models.LocalTime `json:"-" xorm:"default '0000-00-00 00:00:00' comment('修改时间') datetime 'update_time' updated"`
	DeleteTime   models.LocalTime `json:"-" xorm:"default '0000-00-00 00:00:00' comment('删除时间') datetime 'delete_time' deleted"`
	LastUpdateId string           `json:"-" xorm:"default '' comment('最后修改用户标识') varchar(50) 'last_update_id'"`
}
