package gateways

import "github.com/jinzhu/gorm"

type DbHandler interface {
	GetDB() *gorm.DB
}
