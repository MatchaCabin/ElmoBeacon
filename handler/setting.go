package handler

import (
	"ElmoBeacon/db"
	"ElmoBeacon/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (a *App) GetSettings() ([]model.Setting, error) {
	var settingList []model.Setting
	err := db.Engine.Find(&settingList)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, errors.New("error occurred when query setting list from db")
	}
	return settingList, nil
}

func (a *App) SetSetting(key, value string) error {
	has, err := db.Engine.Get(&model.Setting{Key: key})
	if err != nil {
		log.Error().Err(err).Msg("")
		return errors.New("error occurred when get setting from db")
	}
	if !has {
		_, err = db.Engine.Insert(&model.Setting{Key: key, Value: value})
		if err != nil {
			log.Error().Err(err).Msg("")
			return errors.New("error occurred when set setting to db")
		}
	} else {
		_, err = db.Engine.Update(&model.Setting{Key: key, Value: value}, &model.Setting{Key: key})
		if err != nil {
			log.Error().Err(err).Msg("")
			return errors.New("error occurred when set setting to db")
		}
	}
	return nil
}
