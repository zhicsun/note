package pkg

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGorm(t *testing.T) {
	dns := "root:@tcp(127.0.0.1:3306)/eo_oslms_scorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	fmt.Println(db, err)
}
