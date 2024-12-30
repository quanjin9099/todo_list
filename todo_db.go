package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

var (
	MainDb *gorm.DB
	SecDb  *gorm.DB

	DbList = make([]*gorm.DB, 0) //模拟分布式数据库
)

func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate")
	}
	err = db.AutoMigrate(&TodoInfo{})
	if err != nil {
		panic("failed to migrate")
	}

	randData(db)

	return db
}

func randData(curDb *gorm.DB) {
	var user User
	curDb.First(&user)
	if user.Id > 0 {
		return
	}

	for i := 0; i < 10; i++ {
		curUser := User{
			Name: "user" + string(rune('a'+i)),
		}
		curDb.Create(&curUser)

		for j := 0; j < 30; j++ {
			curTodo := TodoInfo{
				UserID:      curUser.Id,
				Name:        randStr(16),
				Description: randStr(100),
				Score:       rand.Intn(100),
				Status:      0,
				CreatedAt:   time.Now(),
			}
			curDb.Create(&curTodo)
			fmt.Printf("create todo %+v\n", curTodo)
		}
	}

}
