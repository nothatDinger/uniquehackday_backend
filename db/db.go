package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"unique_hackday/conf"
)

type Model struct {
	ID int `json:"id" gorm:"primary_key"`
}
var db *gorm.DB

func init() {
	var dialector gorm.Dialector
	sec, err := conf.Cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}
	dbName := sec.Key("NAME").String()
	user := sec.Key("USER").String()
	password := sec.Key("PASSWORD").String()
	host := sec.Key("HOST").String()
	tablePrefix := sec.Key("TABLE_PREFIX").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName)

	switch dbType := sec.Key("TYPE").String(); dbType {
	case "mysql":
		dialector = mysql.Open(dsn)
	}
	
	db, err = gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: true,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
	})

	if err != nil {
		log.Fatalf("Fail to connect to database: %v", err)
	}

	err = db.AutoMigrate(&Tag{}, &Scene{})
	if err != nil {
		log.Fatalf("Fail to synchronize with schema Tag: %v", err)
	}

}
