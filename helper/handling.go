package helper

import (
	"gorm.io/gorm"
)

func PanicHandling(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicHandlerGORM(query gorm.DB) {
	if query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			panic("data not found")
		} else {
			panic("db server error")
		}
	}
}
