package db

import (
	"ElmoBeacon/model"
	"github.com/rs/zerolog/log"
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var Engine *xorm.Engine

func InitDB() {
	var err error
	Engine, err = xorm.NewEngine("sqlite", "ElmoBeacon.db")
	if err != nil {
		log.Panic().Err(err).Msg("")
	}

	Engine.SetMapper(names.GonicMapper{})

	err = Engine.Sync(new(model.User), new(model.Record), new(model.Setting))
	if err != nil {
		log.Panic().Err(err).Msg("")
	}
}
