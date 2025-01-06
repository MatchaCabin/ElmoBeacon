package handler

import (
	"ElmoBeacon/request"
	"github.com/inconshreveable/go-update"
	"github.com/pkg/errors"
	"net/http"
)

const Version = ""

func (a *App) GetVersion() string {
	return Version
}

func (a *App) CheckUpdate() (bool, error) {
	remoteVersion, err := request.CheckUpdate()
	if err != nil {
		return false, err
	}

	if remoteVersion != Version {
		return true, nil
	}
	return false, nil
}

func (a *App) UpdateSelf() error {
	resp, err := http.Get("https://gfl2bucket.mcc.wiki/ElmoBeacon.exe")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		return err
	}

	return nil
}
