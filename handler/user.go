package handler

import (
	"ElmoBeacon/db"
	"ElmoBeacon/model"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (a *App) GetUserList() ([]model.User, error) {
	var userList []model.User
	err := db.Engine.Find(&userList)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, errors.New("error occurred when query user list from db")
	}
	return userList, nil
}
