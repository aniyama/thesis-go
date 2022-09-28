package gateways

import "gorm.io/gorm"

type DbHandler interface {
	GetDB() *gorm.DB
}
