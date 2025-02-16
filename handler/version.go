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

func (a *App) GetLatestVersion() (string, error) {
	return request.GetLatestVersion()
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
