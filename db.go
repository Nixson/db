package db

import (
	"github.com/Nixson/db/postgres"
	"github.com/Nixson/environment"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var fncs = make([]func(), 0)

func InitDb() {
	env := environment.GetEnv()
	switch env.GetString("db.driver") {
	case "postgres":
		postgres.InitDb()
		Instance = postgres.Get()
	}
}

func Get() *gorm.DB {
	if Instance == nil {
		InitDb()
		for _, fnc := range fncs {
			fnc()
		}
	}
	return Instance
}
func AfterInit(function func()) {
	if Instance == nil {
		fncs = append(fncs, function)
	} else {
		function()
	}
}
