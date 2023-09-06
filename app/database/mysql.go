package database

import (
	"fmt"
	"immersive_project/klp3/app/config"
	class "immersive_project/klp3/features/classes/data"
	mentee "immersive_project/klp3/features/mentee/data"
	menteeLog "immersive_project/klp3/features/menteelogs/data"
	user "immersive_project/klp3/features/users/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func InittialMigration(db *gorm.DB) {
	db.AutoMigrate(&user.User{}, &class.Classes{}, &mentee.Mentee{}, &menteeLog.MenteeLog{}, &mentee.Status{})
}
